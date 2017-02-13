// Copyright 2016 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.
//
// Copyright 2014 Christian Schramm. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package shoco

import (
	"bytes"
	"encoding/binary"
	"errors"
)

var ErrInvalid = errors.New("shoco: invalid input")

type Pack struct {
	Word          uint32
	BytesPacked   int
	BytesUnpacked int
	Offsets       [8]uint
	Masks         [8]int16
	HeaderMask    byte
	Header        byte
}

type Model struct {
	ChrsByChrId                 []byte
	ChrIdsByChr                 [256]int8
	SuccessorIdsByChrIdAndChrId [][]int8
	ChrsByChrAndSuccessorId     [][]byte
	Packs                       []Pack

	MinChr byte
	MaxChr byte

	MaxSuccessorN int
}

func decodeHeader(val byte) int {
	i := -1

	for val&0x80 != 0 {
		val <<= 1
		i++
	}

	return i
}

func checkIndices(indices []int16, pack *Pack) bool {
	for i := 0; i < pack.BytesUnpacked; i++ {
		if indices[i] > pack.Masks[i] {
			return false
		}
	}

	return true
}

func (m *Model) findBestEncoding(indices []int16, nConsecutive int) int {
	for p := len(m.Packs) - 1; p >= 0; p-- {
		if nConsecutive >= m.Packs[p].BytesUnpacked && checkIndices(indices, &m.Packs[p]) {
			return p
		}
	}

	return -1
}

func (m *Model) Compress(in []byte) (out []byte) {
	return m.compress(in, false)
}

func (m *Model) ProposedCompress(in []byte) (out []byte) {
	return m.compress(in, true)
}

func (m *Model) compress(in []byte, proposed bool) (out []byte) {
	var buf bytes.Buffer
	buf.Grow(len(in))

	indices := make([]int16, m.MaxSuccessorN+1)

	for len(in) != 0 {
		// find the longest string of known successors
		indices[0] = int16(m.ChrIdsByChr[in[0]])

		if lastChrIndex := indices[0]; lastChrIndex >= 0 {
			nConsecutive := 1
			for ; nConsecutive <= m.MaxSuccessorN && nConsecutive < len(in); nConsecutive++ {
				currentIndex := m.ChrIdsByChr[in[nConsecutive]]
				if currentIndex < 0 { // '\0' is always -1
					break
				}

				sucessorIndex := m.SuccessorIdsByChrIdAndChrId[lastChrIndex][currentIndex]
				if sucessorIndex < 0 {
					break
				}

				indices[nConsecutive] = int16(sucessorIndex)
				lastChrIndex = int16(currentIndex)
			}

			if nConsecutive >= 2 {
				if packN := m.findBestEncoding(indices, nConsecutive); packN >= 0 {
					code := m.Packs[packN].Word
					for i := 0; i < m.Packs[packN].BytesUnpacked; i++ {
						code |= uint32(indices[i]) << m.Packs[packN].Offsets[i]
					}

					var codeBuf [4]byte
					binary.BigEndian.PutUint32(codeBuf[:], code)
					buf.Write(codeBuf[:m.Packs[packN].BytesPacked])

					in = in[m.Packs[packN].BytesUnpacked:]
					continue
				}
			}
		}

		if proposed {
			// See https://github.com/Ed-von-Schleck/shoco/issues/11
			if in[0]&0x80 != 0 || in[0] < 0x09 {
				j := byte(1)
				for ; int(j) < len(in) && j < 0x09 && (in[j]&0x80 != 0 || in[j] < 0x09); j++ {
				}

				buf.WriteByte(j - 1)
				buf.Write(in[:j])
				in = in[j:]
			} else {
				buf.WriteByte(in[0])
				in = in[1:]
			}

			continue
		}

		if in[0]&0x80 != 0 || in[0] == 0x00 { // non-ascii case or NUL char
			// Encoding NUL chars in this way is not compatible with
			// shoco_compress. shoco_compress terminates the compression
			// upon encountering a NUL char. shoco_decompress will
			// nonetheless correctly decode compressed strings that
			// contained NUL chars.

			buf.WriteByte(0x00) // put in a sentinel byte
		}

		buf.WriteByte(in[0])
		in = in[1:]
	}

	return buf.Bytes()
}

func (m *Model) Decompress(in []byte) (out []byte, err error) {
	return m.decompress(in, false)
}

func (m *Model) ProposedDecompress(in []byte) (out []byte, err error) {
	return m.decompress(in, true)
}

func (m *Model) decompress(in []byte, proposed bool) (out []byte, err error) {
	var buf bytes.Buffer
	buf.Grow(len(in) * 2)

	for len(in) != 0 {
		mark := decodeHeader(in[0])
		if mark < 0 {
			if proposed {
				// See https://github.com/Ed-von-Schleck/shoco/issues/11
				if in[0] < 0x09 {
					j := in[0] + 1
					if len(in) < 1+int(j) {
						return nil, ErrInvalid
					}

					buf.Write(in[1 : 1+j])
					in = in[1+j:]
				} else {
					buf.WriteByte(in[0])
					in = in[1:]
				}

				continue
			}

			if in[0] == 0x00 { // ignore the sentinel value for non-ascii chars
				if len(in) < 2 {
					return nil, ErrInvalid
				}

				buf.WriteByte(in[1])
				in = in[2:]
			} else {
				buf.WriteByte(in[0])
				in = in[1:]
			}

			continue
		}

		if mark >= len(m.Packs) || m.Packs[mark].BytesPacked > len(in) {
			return nil, ErrInvalid
		}

		var codeBuf [4]byte
		copy(codeBuf[:], in[:m.Packs[mark].BytesPacked])
		code := binary.BigEndian.Uint32(codeBuf[:])

		offset, mask := m.Packs[mark].Offsets[0], m.Packs[mark].Masks[0]

		lastChr := m.ChrsByChrId[(code>>offset)&uint32(mask)]
		buf.WriteByte(lastChr)

		for i := 1; i < m.Packs[mark].BytesUnpacked; i++ {
			offset, mask := m.Packs[mark].Offsets[i], m.Packs[mark].Masks[i]

			lastChr = m.ChrsByChrAndSuccessorId[lastChr-m.MinChr][(code>>offset)&uint32(mask)]
			buf.WriteByte(lastChr)
		}

		in = in[m.Packs[mark].BytesPacked:]
	}

	return buf.Bytes(), nil
}

func Compress(in []byte) (out []byte) {
	return DefaultModel.Compress(in)
}

func ProposedCompress(in []byte) (out []byte) {
	return DefaultModel.ProposedCompress(in)
}

func Decompress(in []byte) (out []byte, err error) {
	return DefaultModel.Decompress(in)
}

func ProposedDecompress(in []byte) (out []byte, err error) {
	return DefaultModel.ProposedDecompress(in)
}
