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

func decodeHeader(val byte) int {
	i := -1

	for val&0x80 != 0 {
		val <<= 1
		i++
	}

	return i
}

func checkIndices(indices *[maxSuccessorN + 1]int16, pack *pack) bool {
	for i := 0; i < pack.bytesUnpacked; i++ {
		if indices[i] > pack.masks[i] {
			return false
		}
	}

	return true
}

func findBestEncoding(indices *[maxSuccessorN + 1]int16, nConsecutive int) int {
	for p := packCount - 1; p >= 0; p-- {
		if nConsecutive >= packs[p].bytesUnpacked && checkIndices(indices, &packs[p]) {
			return p
		}
	}

	return -1
}

// Compress ...
//
// in must not contain any zero-bytes otherwise Decompress will
// fail.
func Compress(in []byte) (out []byte) {
	return compress(in, false)
}

func ProposedCompress(in []byte) (out []byte) {
	return compress(in, true)
}

func compress(in []byte, proposed bool) (out []byte) {
	var buf bytes.Buffer
	buf.Grow(len(in))

	var indices [maxSuccessorN + 1]int16

	for len(in) != 0 {
		// find the longest string of known successors
		indices[0] = int16(chrIdsByChr[in[0]])

		if lastChrIndex := indices[0]; lastChrIndex >= 0 {
			nConsecutive := 1
			for ; nConsecutive <= maxSuccessorN && nConsecutive < len(in); nConsecutive++ {
				currentIndex := chrIdsByChr[in[nConsecutive]]
				if currentIndex < 0 { // '\0' is always -1
					break
				}

				sucessorIndex := successorIdsByChrIdAndChrId[lastChrIndex][currentIndex]
				if sucessorIndex < 0 {
					break
				}

				indices[nConsecutive] = int16(sucessorIndex)
				lastChrIndex = int16(currentIndex)
			}

			if nConsecutive >= 2 {
				if packN := findBestEncoding(&indices, nConsecutive); packN >= 0 {
					code := packs[packN].word
					for i := 0; i < packs[packN].bytesUnpacked; i++ {
						code |= uint32(indices[i]) << packs[packN].offsets[i]
					}

					var codeBuf [4]byte
					binary.BigEndian.PutUint32(codeBuf[:], code)
					buf.Write(codeBuf[:packs[packN].bytesPacked])

					in = in[packs[packN].bytesUnpacked:]
					continue
				}
			}
		}

		if proposed {
			// See https://github.com/Ed-von-Schleck/shoco/issues/11
			if in[0]&0x80 != 0 || in[0] < 0x09 {
				j := byte(1)
				for ; int(j) < len(in) && j <= 0x09; j++ {
					if in[j]&0x80 == 0 && in[j] >= 0x09 {
						break
					}
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

		if in[0]&0x80 != 0 { // non-ascii case
			buf.WriteByte(0x00) // put in a sentinel byte
		}

		buf.WriteByte(in[0])
		in = in[1:]
	}

	return buf.Bytes()
}

func Decompress(in []byte) (out []byte, err error) {
	return decompress(in, false)
}

func ProposedDecompress(in []byte) (out []byte, err error) {
	return decompress(in, true)
}

func decompress(in []byte, proposed bool) (out []byte, err error) {
	var buf bytes.Buffer
	buf.Grow(len(in) * 2)

	for len(in) != 0 {
		mark := decodeHeader(in[0])
		if mark < 0 {
			if proposed {
				// See https://github.com/Ed-von-Schleck/shoco/issues/11
				if in[0] < 0x09 {
					j := in[0] + 1
					if len(in) < int(j) {
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

		if mark >= len(packs) || packs[mark].bytesPacked > len(in) {
			return nil, ErrInvalid
		}

		var codeBuf [4]byte
		copy(codeBuf[:], in[:packs[mark].bytesPacked])
		code := binary.BigEndian.Uint32(codeBuf[:])

		offset, mask := packs[mark].offsets[0], packs[mark].masks[0]

		lastChr := chrsByChrId[(code>>offset)&uint32(mask)]
		buf.WriteByte(lastChr)

		for i := 1; i < packs[mark].bytesUnpacked; i++ {
			offset, mask := packs[mark].offsets[i], packs[mark].masks[i]

			lastChr = chrsByChrAndSuccessorId[lastChr-minChr][(code>>offset)&uint32(mask)]
			buf.WriteByte(lastChr)
		}

		in = in[packs[mark].bytesPacked:]
	}

	return buf.Bytes(), nil
}
