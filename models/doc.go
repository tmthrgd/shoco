// Copyright 2016 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

// Package models contains various compression models for shoco.
package models

import "github.com/tmthrgd/shoco"

func check(model *shoco.Model) {
	// The call to Compress is used for it's side-effect of
	// calling (*Model).checkValid.
	model.Compress(nil)
}

// WordsEn is a model optimised for words of the English langauge.
func WordsEn() *shoco.Model {
	check(shoco.WordsEnModel)
	return shoco.WordsEnModel
}

// TextEn is a model optimised for English langauge text.
func TextEn() *shoco.Model {
	check(shoco.TextEnModel)
	return shoco.TextEnModel
}

// FilePath is a model optimised for filepaths.
func FilePath() *shoco.Model {
	check(shoco.FilePathModel)
	return shoco.FilePathModel
}
