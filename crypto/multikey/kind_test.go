// Copyright 2019 Finobo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package multikey

import (
	"errors"
	"testing"

	"github.com/mailchain/mailchain/crypto"
	"github.com/mailchain/mailchain/crypto/ed25519"
	"github.com/mailchain/mailchain/crypto/secp256k1"
	"github.com/stretchr/testify/assert"
)

func TestKeyKindFromSignature(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		pubKey  []byte
		message []byte
		sig     []byte
		kinds   []string
	}
	tests := []struct {
		name    string
		args    args
		wantKey crypto.PublicKey
		wantErr error
	}{
		{
			"success-ed25519-duplicate-key-kinds",
			args{
				[]byte{0x72, 0x3c, 0xaa, 0x23, 0xa5, 0xb5, 0x11, 0xaf, 0x5a, 0xd7, 0xb7, 0xef, 0x60, 0x76, 0xe4, 0x14, 0xab, 0x7e, 0x75, 0xa9, 0xdc, 0x91, 0xe, 0xa6, 0xe, 0x41, 0x7a, 0x2b, 0x77, 0xa, 0x56, 0x71},
				[]byte("egassem"),
				[]byte{0xde, 0x6c, 0x88, 0xe6, 0x9c, 0x9f, 0x93, 0xb, 0x59, 0xdd, 0xf4, 0x80, 0xc2, 0x9a, 0x55, 0x79, 0xec, 0x89, 0x5c, 0xa9, 0x7a, 0x36, 0xf6, 0x69, 0x74, 0xc1, 0xf0, 0x15, 0x5c, 0xc0, 0x66, 0x75, 0x2e, 0xcd, 0x9a, 0x9b, 0x41, 0x35, 0xd2, 0x72, 0x32, 0xe0, 0x54, 0x80, 0xbc, 0x98, 0x58, 0x1, 0xa9, 0xfd, 0xe4, 0x27, 0xc7, 0xef, 0xa5, 0x42, 0x5f, 0xf, 0x46, 0x49, 0xb8, 0xad, 0xbd, 0x5},
				[]string{crypto.ED25519, crypto.ED25519, crypto.ED25519},
			},
			func() crypto.PublicKey {
				k, _ := ed25519.PublicKeyFromBytes([]byte{0x72, 0x3c, 0xaa, 0x23, 0xa5, 0xb5, 0x11, 0xaf, 0x5a, 0xd7, 0xb7, 0xef, 0x60, 0x76, 0xe4, 0x14, 0xab, 0x7e, 0x75, 0xa9, 0xdc, 0x91, 0xe, 0xa6, 0xe, 0x41, 0x7a, 0x2b, 0x77, 0xa, 0x56, 0x71})
				return k
			}(),
			nil,
		},
		{
			"success-secp256k1-duplicate-key-kinds",
			args{
				func() []byte {
					k, _ := secp256k1.PrivateKeyFromHex("01901E63389EF02EAA7C5782E08B40D98FAEF835F28BD144EECF5614A415943F")
					return k.PublicKey().Bytes()
				}(),
				[]byte("egassem"),
				[]byte{0xe9, 0x33, 0xe, 0x4a, 0xe3, 0x5, 0x19, 0xea, 0x36, 0x37, 0x19, 0xdd, 0xbc, 0x91, 0xfd, 0x4f, 0xd3, 0x64, 0x9b, 0xdc, 0xf0, 0x74, 0x36, 0x16, 0xc9, 0x81, 0xfc, 0x6d, 0x3c, 0x7e, 0xb0, 0xd0, 0x6e, 0xdd, 0x4, 0x13, 0xfd, 0x15, 0xe5, 0xec, 0x64, 0x6e, 0x63, 0xe0, 0x84, 0xdb, 0xb2, 0xd7, 0xcf, 0x18, 0x3d, 0x81, 0x1e, 0x31, 0x36, 0x77, 0x39, 0x86, 0x4b, 0x58, 0xb8, 0x23, 0xed, 0xc, 0x1},
				[]string{crypto.SECP256K1, crypto.SECP256K1},
			},
			func() crypto.PublicKey {
				k, _ := secp256k1.PublicKeyFromBytes([]byte{0x2, 0x69, 0xd9, 0x8, 0x51, 0xe, 0x35, 0x5b, 0xeb, 0x1d, 0x5b, 0xf2, 0xdf, 0x81, 0x29, 0xe5, 0xb6, 0x40, 0x1e, 0x19, 0x69, 0x89, 0x1e, 0x80, 0x16, 0xa0, 0xb2, 0x30, 0x7, 0x39, 0xbb, 0xb0, 0x6})
				return k
			}(),
			nil,
		},
		{
			"success-ed25519-sofia-key",
			args{
				[]byte{0x72, 0x3c, 0xaa, 0x23, 0xa5, 0xb5, 0x11, 0xaf, 0x5a, 0xd7, 0xb7, 0xef, 0x60, 0x76, 0xe4, 0x14, 0xab, 0x7e, 0x75, 0xa9, 0xdc, 0x91, 0xe, 0xa6, 0xe, 0x41, 0x7a, 0x2b, 0x77, 0xa, 0x56, 0x71},
				[]byte("egassem"),
				[]byte{0xde, 0x6c, 0x88, 0xe6, 0x9c, 0x9f, 0x93, 0xb, 0x59, 0xdd, 0xf4, 0x80, 0xc2, 0x9a, 0x55, 0x79, 0xec, 0x89, 0x5c, 0xa9, 0x7a, 0x36, 0xf6, 0x69, 0x74, 0xc1, 0xf0, 0x15, 0x5c, 0xc0, 0x66, 0x75, 0x2e, 0xcd, 0x9a, 0x9b, 0x41, 0x35, 0xd2, 0x72, 0x32, 0xe0, 0x54, 0x80, 0xbc, 0x98, 0x58, 0x1, 0xa9, 0xfd, 0xe4, 0x27, 0xc7, 0xef, 0xa5, 0x42, 0x5f, 0xf, 0x46, 0x49, 0xb8, 0xad, 0xbd, 0x5},
				[]string{crypto.ED25519},
			},
			func() crypto.PublicKey {
				k, _ := ed25519.PublicKeyFromBytes([]byte{0x72, 0x3c, 0xaa, 0x23, 0xa5, 0xb5, 0x11, 0xaf, 0x5a, 0xd7, 0xb7, 0xef, 0x60, 0x76, 0xe4, 0x14, 0xab, 0x7e, 0x75, 0xa9, 0xdc, 0x91, 0xe, 0xa6, 0xe, 0x41, 0x7a, 0x2b, 0x77, 0xa, 0x56, 0x71})
				return k
			}(),
			nil,
		},
		{
			"success-secp256k1-sofia-key",
			args{
				func() []byte {
					k, _ := secp256k1.PrivateKeyFromHex("01901E63389EF02EAA7C5782E08B40D98FAEF835F28BD144EECF5614A415943F")
					return k.PublicKey().Bytes()
				}(),
				[]byte("egassem"),
				[]byte{0xe9, 0x33, 0xe, 0x4a, 0xe3, 0x5, 0x19, 0xea, 0x36, 0x37, 0x19, 0xdd, 0xbc, 0x91, 0xfd, 0x4f, 0xd3, 0x64, 0x9b, 0xdc, 0xf0, 0x74, 0x36, 0x16, 0xc9, 0x81, 0xfc, 0x6d, 0x3c, 0x7e, 0xb0, 0xd0, 0x6e, 0xdd, 0x4, 0x13, 0xfd, 0x15, 0xe5, 0xec, 0x64, 0x6e, 0x63, 0xe0, 0x84, 0xdb, 0xb2, 0xd7, 0xcf, 0x18, 0x3d, 0x81, 0x1e, 0x31, 0x36, 0x77, 0x39, 0x86, 0x4b, 0x58, 0xb8, 0x23, 0xed, 0xc, 0x1},
				[]string{crypto.SECP256K1},
			},
			func() crypto.PublicKey {
				k, _ := secp256k1.PublicKeyFromBytes([]byte{0x2, 0x69, 0xd9, 0x8, 0x51, 0xe, 0x35, 0x5b, 0xeb, 0x1d, 0x5b, 0xf2, 0xdf, 0x81, 0x29, 0xe5, 0xb6, 0x40, 0x1e, 0x19, 0x69, 0x89, 0x1e, 0x80, 0x16, 0xa0, 0xb2, 0x30, 0x7, 0x39, 0xbb, 0xb0, 0x6})
				return k
			}(),
			nil,
		},
		{
			"success-ed25519-charlotte-key",
			args{
				[]byte{0x2e, 0x32, 0x2f, 0x87, 0x40, 0xc6, 0x1, 0x72, 0x11, 0x1a, 0xc8, 0xea, 0xdc, 0xdd, 0xa2, 0x51, 0x2f, 0x90, 0xd0, 0x6d, 0xe, 0x50, 0x3e, 0xf1, 0x89, 0x97, 0x9a, 0x15, 0x9b, 0xec, 0xe1, 0xe8},
				[]byte("message"),
				[]byte{0x7d, 0x51, 0xea, 0xfa, 0x52, 0x78, 0x31, 0x69, 0xd0, 0xa9, 0x4a, 0xc, 0x9f, 0x2b, 0xca, 0xd5, 0xe0, 0x3d, 0x29, 0x17, 0x33, 0x0, 0x93, 0xf, 0xf3, 0xc7, 0xd6, 0x3b, 0xfd, 0x64, 0x17, 0xae, 0x1b, 0xc8, 0x1f, 0xef, 0x51, 0xba, 0x14, 0x9a, 0xe8, 0xa1, 0xe1, 0xda, 0xe0, 0x5f, 0xdc, 0xa5, 0x7, 0x8b, 0x14, 0xba, 0xc4, 0xcf, 0x26, 0xcc, 0xc6, 0x1, 0x1e, 0x5e, 0xab, 0x77, 0x3, 0xc},
				[]string{crypto.ED25519},
			},
			func() crypto.PublicKey {
				k, _ := ed25519.PublicKeyFromBytes([]byte{0x2e, 0x32, 0x2f, 0x87, 0x40, 0xc6, 0x1, 0x72, 0x11, 0x1a, 0xc8, 0xea, 0xdc, 0xdd, 0xa2, 0x51, 0x2f, 0x90, 0xd0, 0x6d, 0xe, 0x50, 0x3e, 0xf1, 0x89, 0x97, 0x9a, 0x15, 0x9b, 0xec, 0xe1, 0xe8})
				return k
			}(),
			nil,
		},
		{
			"success-secp256k1-charlotte-key",
			args{
				func() []byte {
					k, _ := secp256k1.PrivateKeyFromHex("DF4BA9F6106AD2846472F759476535E55C5805D8337DF5A11C3B139F438B98B3")
					return k.PublicKey().Bytes()
				}(),
				[]byte("message"),
				[]byte{0x9d, 0xf7, 0x76, 0xab, 0xde, 0x8c, 0x20, 0x55, 0xc3, 0x4, 0x68, 0x37, 0xa8, 0x66, 0xf8, 0x89, 0x95, 0xf9, 0x82, 0xf0, 0x4b, 0xb8, 0x23, 0x40, 0xf0, 0x3, 0x8, 0x6a, 0x32, 0xa7, 0xac, 0xef, 0x5f, 0xa, 0xea, 0xda, 0x60, 0xbf, 0x9, 0xd5, 0xc3, 0x27, 0x61, 0xa, 0xc5, 0xc8, 0x33, 0xe3, 0xa0, 0x79, 0xdf, 0x6d, 0xe1, 0x9c, 0xa8, 0xcc, 0x33, 0xea, 0x1d, 0xe6, 0x3, 0x34, 0xb1, 0xa1, 0x0},
				[]string{crypto.SECP256K1},
			},
			func() crypto.PublicKey {
				k, _ := secp256k1.PublicKeyFromBytes([]byte{0x3, 0xbd, 0xf6, 0xfb, 0x97, 0xc9, 0x7c, 0x12, 0x6b, 0x49, 0x21, 0x86, 0xa4, 0xd5, 0xb2, 0x8f, 0x34, 0xf0, 0x67, 0x1a, 0x5a, 0xac, 0xc9, 0x74, 0xda, 0x3b, 0xde, 0xb, 0xe9, 0x3e, 0x45, 0xa1, 0xc5})
				return k
			}(),
			nil,
		},
		{
			"success-ed25519-mix-key-kinds",
			args{
				[]byte{0x72, 0x3c, 0xaa, 0x23, 0xa5, 0xb5, 0x11, 0xaf, 0x5a, 0xd7, 0xb7, 0xef, 0x60, 0x76, 0xe4, 0x14, 0xab, 0x7e, 0x75, 0xa9, 0xdc, 0x91, 0xe, 0xa6, 0xe, 0x41, 0x7a, 0x2b, 0x77, 0xa, 0x56, 0x71},
				[]byte("egassem"),
				[]byte{0xde, 0x6c, 0x88, 0xe6, 0x9c, 0x9f, 0x93, 0xb, 0x59, 0xdd, 0xf4, 0x80, 0xc2, 0x9a, 0x55, 0x79, 0xec, 0x89, 0x5c, 0xa9, 0x7a, 0x36, 0xf6, 0x69, 0x74, 0xc1, 0xf0, 0x15, 0x5c, 0xc0, 0x66, 0x75, 0x2e, 0xcd, 0x9a, 0x9b, 0x41, 0x35, 0xd2, 0x72, 0x32, 0xe0, 0x54, 0x80, 0xbc, 0x98, 0x58, 0x1, 0xa9, 0xfd, 0xe4, 0x27, 0xc7, 0xef, 0xa5, 0x42, 0x5f, 0xf, 0x46, 0x49, 0xb8, 0xad, 0xbd, 0x5},
				[]string{crypto.ED25519, crypto.SECP256K1},
			},
			func() crypto.PublicKey {
				k, _ := ed25519.PublicKeyFromBytes([]byte{0x72, 0x3c, 0xaa, 0x23, 0xa5, 0xb5, 0x11, 0xaf, 0x5a, 0xd7, 0xb7, 0xef, 0x60, 0x76, 0xe4, 0x14, 0xab, 0x7e, 0x75, 0xa9, 0xdc, 0x91, 0xe, 0xa6, 0xe, 0x41, 0x7a, 0x2b, 0x77, 0xa, 0x56, 0x71})
				return k
			}(),
			nil,
		},
		{
			"success-secp256k1-mix-key-kinds",
			args{
				func() []byte {
					k, _ := secp256k1.PrivateKeyFromHex("DF4BA9F6106AD2846472F759476535E55C5805D8337DF5A11C3B139F438B98B3")
					return k.PublicKey().Bytes()
				}(),
				[]byte("message"),
				[]byte{0x9d, 0xf7, 0x76, 0xab, 0xde, 0x8c, 0x20, 0x55, 0xc3, 0x4, 0x68, 0x37, 0xa8, 0x66, 0xf8, 0x89, 0x95, 0xf9, 0x82, 0xf0, 0x4b, 0xb8, 0x23, 0x40, 0xf0, 0x3, 0x8, 0x6a, 0x32, 0xa7, 0xac, 0xef, 0x5f, 0xa, 0xea, 0xda, 0x60, 0xbf, 0x9, 0xd5, 0xc3, 0x27, 0x61, 0xa, 0xc5, 0xc8, 0x33, 0xe3, 0xa0, 0x79, 0xdf, 0x6d, 0xe1, 0x9c, 0xa8, 0xcc, 0x33, 0xea, 0x1d, 0xe6, 0x3, 0x34, 0xb1, 0xa1, 0x0},
				[]string{crypto.SECP256K1, crypto.ED25519},
			},
			func() crypto.PublicKey {
				k, _ := secp256k1.PublicKeyFromBytes([]byte{0x3, 0xbd, 0xf6, 0xfb, 0x97, 0xc9, 0x7c, 0x12, 0x6b, 0x49, 0x21, 0x86, 0xa4, 0xd5, 0xb2, 0x8f, 0x34, 0xf0, 0x67, 0x1a, 0x5a, 0xac, 0xc9, 0x74, 0xda, 0x3b, 0xde, 0xb, 0xe9, 0x3e, 0x45, 0xa1, 0xc5})
				return k
			}(),
			nil,
		},
		{
			"err-invalid-public-key-bytes",
			args{
				[]byte{0x1},
				[]byte("unknown"),
				nil,
				[]string{crypto.ED25519, crypto.SECP256K1},
			},
			nil,
			ErrNoMatch,
		},
		{
			"err-no-key-kinds",
			args{
				func() []byte {
					k, _ := secp256k1.PrivateKeyFromHex("DF4BA9F6106AD2846472F759476535E55C5805D8337DF5A11C3B139F438B98B3")
					return k.PublicKey().Bytes()
				}(),
				[]byte("message"),
				[]byte{0x9d, 0xf7, 0x76, 0xab, 0xde, 0x8c, 0x20, 0x55, 0xc3, 0x4, 0x68, 0x37, 0xa8, 0x66, 0xf8, 0x89, 0x95, 0xf9, 0x82, 0xf0, 0x4b, 0xb8, 0x23, 0x40, 0xf0, 0x3, 0x8, 0x6a, 0x32, 0xa7, 0xac, 0xef, 0x5f, 0xa, 0xea, 0xda, 0x60, 0xbf, 0x9, 0xd5, 0xc3, 0x27, 0x61, 0xa, 0xc5, 0xc8, 0x33, 0xe3, 0xa0, 0x79, 0xdf, 0x6d, 0xe1, 0x9c, 0xa8, 0xcc, 0x33, 0xea, 0x1d, 0xe6, 0x3, 0x34, 0xb1, 0xa1, 0x0},
				nil,
			},
			nil,
			ErrNoMatch,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := KeyKindFromSignature(tt.args.pubKey, tt.args.message, tt.args.sig, tt.args.kinds)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("KeyKindFromSignature() err = %v, want %v", err, tt.wantErr)
			}

			if !assert.Equal(key, tt.wantKey) {
				t.Errorf("KeyKindFromSignature() key = %v, want %v", key, tt.wantKey)
			}
		})
	}
}

func TestRemoveDuplicates(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		in   []string
		want []string
	}{
		{
			[]string{"x", "x", "y"},
			[]string{"x", "y"},
		},
		{
			nil,
			nil,
		},
		{
			[]string{},
			[]string{},
		},
		{
			[]string{"c", "a", "c"},
			[]string{"c", "a"},
		},
		{
			[]string{"d", "a", "f", "a", "e", "e", "z", "f", "t"},
			[]string{"d", "a", "f", "e", "z", "t"},
		},
	}
	for _, tt := range tests {
		t.Run("remove-duplicates", func(t *testing.T) {
			got := removeDuplicates(tt.in)
			if !assert.Equal(tt.want, got) {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}