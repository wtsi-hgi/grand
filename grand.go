/*******************************************************************************
 * Copyright (c) 2015 András Belicza
 * Copyright (c) 2017 Jonathan ES Lin
 * Copyright (c) 2023 Genome Research Ltd.
 *
 * Authors:
 *  - András Belicza
 *  - Jonathan ES Lin
 *	- Sendu Bala <sb10@sanger.ac.uk>
 *
 * Permission is hereby granted, free of charge, to any person obtaining
 * a copy of this software and associated documentation files (the
 * "Software"), to deal in the Software without restriction, including
 * without limitation the rights to use, copy, modify, merge, publish,
 * distribute, sublicense, and/or sell copies of the Software, and to
 * permit persons to whom the Software is furnished to do so, subject to
 * the following conditions:
 *
 * The above copyright notice and this permission notice shall be included
 * in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
 * EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
 * MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
 * IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
 * CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
 * TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
 * SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 ******************************************************************************/

package grand

import (
	"math/rand"
	"time"
	"unsafe"
)

// Credits to icza of https://stackoverflow.com/a/31832326/1161743 for the
// original code. And to https://github.com/wtsi-hgi/grand for first making it a
// package.

// Character sets for random string generation.
const (
	CharSetEnglishAlphabet          = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CharSetEnglishAlphabetLowercase = "abcdefghijklmnopqrstuvwxyz"
	CharSetEnglishAlphabetUppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CharSetBase62                   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var defaultLetterBytes = CharSetEnglishAlphabet

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// Generator is the instance that generates random strings from the given
// charset.
type Generator struct {
	charset string
	source  rand.Source
}

// New returns a new Generator instance. NB: the returned Generator is not
// suitable for concurrent use. But re-using it in the same go-routine may be
// beneficial.
func New(charset string) *Generator {
	return &Generator{
		charset: charset,
		source:  rand.NewSource(time.Now().UnixNano()),
	}
}

// String generates a length-n random string from our character set.
func (g *Generator) String(n int) string {
	b := make([]byte, n)

	for i, cache, remain := n-1, g.source.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = g.source.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(g.charset) {
			b[i] = g.charset[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

// String generates a length-n random string from the the English alphabet,
// upper and lower case.
func String(n int) string {
	return New(CharSetEnglishAlphabet).String(n)
}

// LcString generates a length-n random string from the the English alphabet,
// lower case only.
func LcString(n int) string {
	return New(CharSetEnglishAlphabetLowercase).String(n)
}
