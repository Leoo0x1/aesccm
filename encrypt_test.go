/*
 * MIT License
 *
 * Copyright (C) 2021 Crimson Technologies LLC. All rights reserved.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package aesccm

import (
	"bytes"
	"crypto/aes"
	"testing"
)

func TestExamples(t *testing.T) {
	// fill byte slice with its own indexes.
	// used for example C4
	c4a := make([]byte, 524288/8)
	for i := range c4a {
		c4a[i] = byte(i)
	}

	examples := []struct {
		Key        []byte
		Nonce      []byte
		Data       []byte
		PlainText  []byte
		CipherText []byte
		TagLen     int
	}{
		{ // C.1
			[]byte{0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4e, 0x4f},
			[]byte{0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16},
			[]byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07},
			[]byte{0x20, 0x21, 0x22, 0x23},
			[]byte{0x71, 0x62, 0x01, 0x5b, 0x4d, 0xac, 0x25, 0x5d},
			4,
		},
		{ // C.2
			[]byte{0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4e, 0x4f},
			[]byte{0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17},
			[]byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f},
			[]byte{0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f},
			[]byte{0xd2, 0xa1, 0xf0, 0xe0, 0x51, 0xea, 0x5f, 0x62, 0x08, 0x1a, 0x77, 0x92, 0x07, 0x3d, 0x59, 0x3d, 0x1f, 0xc6, 0x4f, 0xbf, 0xac, 0xcd},
			6,
		},
		{ // C.3
			[]byte{0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4e, 0x4f},
			[]byte{0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b},
			[]byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10, 0x11, 0x12, 0x13},
			[]byte{0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37},

			[]byte{0xe3, 0xb2, 0x01, 0xa9, 0xf5, 0xb7, 0x1a, 0x7a, 0x9b, 0x1c, 0xea, 0xec, 0xcd, 0x97, 0xe7, 0x0b, 0x61, 0x76, 0xaa, 0xd9, 0xa4, 0x42, 0x8a, 0xa5, 0x48, 0x43, 0x92, 0xfb, 0xc1, 0xb0, 0x99, 0x51},
			8,
		},
		{ // C.4
			[]byte{0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4e, 0x4f},
			[]byte{0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c},
			c4a,
			[]byte{0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f},

			[]byte{0x69, 0x91, 0x5d, 0xad, 0x1e, 0x84, 0xc6, 0x37, 0x6a, 0x68, 0xc2, 0x96, 0x7e, 0x4d, 0xab, 0x61, 0x5a, 0xe0, 0xfd, 0x1f, 0xae, 0xc4, 0x4c, 0xc4, 0x84, 0x82, 0x85, 0x29, 0x46, 0x3c, 0xcf, 0x72, 0xb4, 0xac, 0x6b, 0xec, 0x93, 0xe8, 0x59, 0x8e, 0x7f, 0x0d, 0xad, 0xbc, 0xea, 0x5b},
			14,
		},
	}

	// test each example
	for i, example := range examples {
		t.Logf("Testing example #%d...", i)

		c, err := aes.NewCipher(example.Key)
		if err != nil {
			panic(err)
		}

		aesCCM, err := NewCCM(c, len(example.Nonce), example.TagLen)
		if err != nil {
			panic(err)
		}

		cipherText := aesCCM.Seal(nil, example.Nonce, example.PlainText, example.Data)
		if !bytes.Equal(example.CipherText, cipherText) {
			t.Logf("cipher text doesn't match\n")
			t.Fail()
		}

		plainText, err := aesCCM.Open(nil, example.Nonce, example.CipherText, example.Data)
		if err != nil {
			panic(err)
		}
		if !bytes.Equal(example.PlainText, plainText) {
			t.Logf("plain text doesn't match\n")
			t.Fail()
		}

		t.Logf("OK\n")
	}
}