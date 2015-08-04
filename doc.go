// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run gen.go -backup
//go:generate gofmt -w colornames.go

// Package colornames exports a map of named color.RGBA colors based on the
// SVG 1.1 spec. For more details see http://www.w3.org/TR/SVG/types.html
package colornames
