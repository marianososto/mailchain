// Copyright 2022 Mailchain Ltd.
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

package nacl

import (
	"time"

	"github.com/mailchain/mailchain/encoding/encodingtest"
	"github.com/mailchain/mailchain/internal/keystore"
	"github.com/mailchain/mailchain/internal/keystore/kdf/scrypt"
)

var (
	encryptedKeyAliceSECP256k1 = keystore.EncryptedKey{
		PublicKeyBytes: []uint8{0x2, 0x69, 0xd9, 0x8, 0x51, 0xe, 0x35, 0x5b, 0xeb, 0x1d, 0x5b, 0xf2, 0xdf, 0x81, 0x29, 0xe5, 0xb6, 0x40, 0x1e, 0x19, 0x69, 0x89, 0x1e, 0x80, 0x16, 0xa0, 0xb2, 0x30, 0x7, 0x39, 0xbb, 0xb0, 0x6},
		CipherText:     []uint8{0xd1, 0x4b, 0xf2, 0xc3, 0x25, 0xd4, 0xd0, 0xc0, 0x47, 0x6f, 0xa8, 0x6f, 0xb6, 0x4, 0x3b, 0x3e, 0x18, 0xd8, 0x6d, 0x33, 0x96, 0xb9, 0xc7, 0x80, 0x6c, 0xe, 0xe, 0x51, 0xa2, 0x11, 0xeb, 0xc7, 0x14, 0x94, 0xb, 0x95, 0x1c, 0x6e, 0xea, 0x9b, 0xfa, 0xd0, 0x79, 0x3, 0x92, 0x22, 0x34, 0x73, 0xb3, 0xa1, 0x7b, 0x57, 0x46, 0xd7, 0x5c, 0x46, 0x37, 0xd3, 0x46, 0x84, 0xe8, 0x20, 0xfc, 0x34, 0x3b, 0x67, 0x95, 0x65, 0x7a, 0x34, 0x1e, 0x77},
		Timestamp:      time.Date(2019, 10, 27, 21, 48, 13, 554552, time.UTC),
		Version:        "dev",
		StorageCipher:  "nacl",
		CurveType:      "secp256k1",
		ID:             "4b1ff06f-6b7a-483d-ba84-1d2feaffd749",
		KDF:            "scrypt",
		ScryptParams: &scrypt.DeriveOpts{
			Len:        32,
			N:          262144,
			P:          1,
			R:          8,
			Salt:       []uint8{0xbc, 0xf5, 0xb2, 0x8d, 0x7a, 0x54, 0x58, 0xd9, 0x9d, 0xa8, 0xe1, 0x18, 0x30, 0x15, 0xb2, 0x3d, 0x7d, 0x24, 0x14, 0xa7, 0x51, 0x1d, 0x9d, 0xb4, 0xe5, 0xcf, 0xd6, 0x4a, 0xb9, 0x79, 0x81, 0xa7},
			Passphrase: "",
		},
	}
	encryptedKeyAliceED25519 = keystore.EncryptedKey{
		PublicKeyBytes: encodingtest.MustDecodeHex("723caa23a5b511af5ad7b7ef6076e414ab7e75a9dc910ea60e417a2b770a5671"),
		CipherText:     []uint8{0x6f, 0x6, 0xad, 0xc4, 0x10, 0xf3, 0xad, 0xf4, 0x90, 0xe3, 0xa9, 0x37, 0x8, 0x15, 0x6b, 0xe4, 0x71, 0xc7, 0x73, 0xc0, 0xf9, 0x11, 0xcc, 0x11, 0x8f, 0xe5, 0xf6, 0x29, 0x44, 0xa1, 0x98, 0x58, 0xf7, 0xc1, 0x45, 0xed, 0x3d, 0x9d, 0x36, 0x55, 0x9f, 0xfd, 0xea, 0xdd, 0xdd, 0xe4, 0x3c, 0x27, 0xd2, 0x9d, 0x82, 0x6d, 0xde, 0xa7, 0xf7, 0x27, 0x43, 0x1, 0x4f, 0x2b, 0xd, 0x81, 0x24, 0xdf, 0x7a, 0xa, 0xb2, 0x9, 0x1c, 0x64, 0x8, 0x89, 0x8b, 0x33, 0x5f, 0x30, 0xed, 0xd6, 0x81, 0x8, 0xba, 0xd6, 0xd2, 0x2e, 0x9a, 0x79, 0x50, 0xd, 0x98, 0x62, 0xf3, 0xb7, 0x54, 0x9b, 0x52, 0xa3, 0x10, 0xe7, 0xaf, 0xf9, 0xca, 0x1, 0xc1, 0xbe},
		CurveType:      "ed25519",
		ID:             "87ec3fcc-91b3-4742-a1ad-92eea526244d",
		Timestamp:      time.Date(2019, 10, 27, 21, 48, 14, 449515, time.UTC),
		Version:        "dev",
		StorageCipher:  "nacl",
		KDF:            "scrypt",
		ScryptParams: &scrypt.DeriveOpts{
			Len:        32,
			N:          262144,
			P:          1,
			R:          8,
			Salt:       []uint8{0x18, 0xb1, 0xf7, 0x56, 0x95, 0x10, 0x76, 0x6a, 0xe9, 0x65, 0x80, 0x44, 0x50, 0x5e, 0x8a, 0x24, 0x6e, 0x24, 0x68, 0x7f, 0x10, 0x3f, 0x8d, 0xb8, 0xd0, 0xc7, 0xa4, 0x96, 0xc5, 0x6, 0xe6, 0x7e},
			Passphrase: "",
		},
	}
	encryptedKeyBobED25519 = keystore.EncryptedKey{
		PublicKeyBytes: encodingtest.MustDecodeHex("2e322f8740c60172111ac8eadcdda2512f90d06d0e503ef189979a159bece1e8"),
		CipherText:     []uint8{0xf4, 0xc2, 0x3a, 0xc8, 0x9e, 0x75, 0xa3, 0x82, 0xf3, 0x90, 0x6d, 0x54, 0x19, 0x28, 0xa, 0xa1, 0x41, 0x8b, 0x4a, 0x3e, 0xcc, 0x98, 0x3f, 0xe2, 0x1c, 0xa2, 0x87, 0xe0, 0x5b, 0x8c, 0x2e, 0x7e, 0xdf, 0x78, 0x2c, 0xbe, 0x71, 0x72, 0x3d, 0xd, 0x95, 0x42, 0xd6, 0x8b, 0xfb, 0x45, 0xff, 0xbd, 0xcd, 0xcb, 0x4c, 0x87, 0x6, 0xf7, 0xd6, 0x79, 0x4a, 0xd6, 0x1f, 0xeb, 0xc8, 0xdd, 0xc4, 0x3a, 0x8d, 0xf, 0x96, 0x91, 0x44, 0x4d, 0xea, 0x14, 0xfe, 0x21, 0x1e, 0x70, 0x9e, 0x8a, 0xd0, 0xdd, 0x8e, 0x36, 0xaa, 0x10, 0x26, 0xf8, 0xf3, 0xf0, 0xea, 0x62, 0x44, 0xfe, 0x89, 0x9, 0x87, 0x5a, 0xc3, 0xa8, 0x14, 0x69, 0x9, 0x1a, 0xd9, 0x8b},
		CurveType:      "ed25519",
		ID:             "5c35894a-eb1a-4737-9b84-c1618e3cbc93",
		Timestamp:      time.Date(2019, 10, 29, 22, 30, 22, 374433, time.UTC),
		Version:        "dev",
		StorageCipher:  "nacl",
		KDF:            "scrypt",
		ScryptParams: &scrypt.DeriveOpts{
			Len:        32,
			N:          262144,
			P:          1,
			R:          8,
			Salt:       []uint8{0x87, 0x4, 0xe8, 0x31, 0x4b, 0x76, 0xbc, 0x4e, 0xc9, 0xfd, 0x6, 0xe4, 0xde, 0x7d, 0x6c, 0x99, 0x5e, 0x87, 0x15, 0xa6, 0xbe, 0xba, 0x3b, 0x81, 0x24, 0x64, 0x72, 0x8b, 0xcc, 0x97, 0x6e, 0x23},
			Passphrase: "",
		},
	}
)

var (
	fileAliceSECP256k1 = []byte(`{"public-key":"AmnZCFEONVvrHVvy34Ep5bZAHhlpiR6AFqCyMAc5u7AG","cipher-text":"0UvywyXU0MBHb6hvtgQ7PhjYbTOWuceAbA4OUaIR68cUlAuVHG7qm/rQeQOSIjRzs6F7V0bXXEY300aE6CD8NDtnlWV6NB53","curve-type":"secp256k1","id":"4b1ff06f-6b7a-483d-ba84-1d2feaffd749","timestamp":"2019-10-27T21:48:13.000554552Z","version":"dev","storage-cipher":"nacl","kdf":"scrypt","scrypt-params":{"len":32,"n":262144,"p":1,"r":8,"salt":"vPWyjXpUWNmdqOEYMBWyPX0kFKdRHZ205c/WSrl5gac="}}`)
	fileAliceED25519   = []byte(`{"public-key":"cjyqI6W1Ea9a17fvYHbkFKt+danckQ6mDkF6K3cKVnE=","cipher-text":"bwatxBDzrfSQ46k3CBVr5HHHc8D5EcwRj+X2KUShmFj3wUXtPZ02VZ/96t3d5Dwn0p2Cbd6n9ydDAU8rDYEk33oKsgkcZAiJizNfMO3WgQi61tIumnlQDZhi87dUm1KjEOev+coBwb4=","curve-type":"ed25519","id":"87ec3fcc-91b3-4742-a1ad-92eea526244d","timestamp":"2019-10-27T21:48:14.000449515Z","version":"dev","storage-cipher":"nacl","kdf":"scrypt","scrypt-params":{"len":32,"n":262144,"p":1,"r":8,"salt":"GLH3VpUQdmrpZYBEUF6KJG4kaH8QP4240MeklsUG5n4="}}`)
	fileBobED25519     = []byte(`{"public-key":"LjIvh0DGAXIRGsjq3N2iUS+Q0G0OUD7xiZeaFZvs4eg=","cipher-text":"9MI6yJ51o4LzkG1UGSgKoUGLSj7MmD/iHKKH4FuMLn7feCy+cXI9DZVC1ov7Rf+9zctMhwb31nlK1h/ryN3EOo0PlpFETeoU/iEecJ6K0N2ONqoQJvjz8OpiRP6JCYdaw6gUaQka2Ys=","curve-type":"ed25519","id":"5c35894a-eb1a-4737-9b84-c1618e3cbc93","timestamp":"2019-10-29T22:30:22.000374433Z","version":"dev","storage-cipher":"nacl","kdf":"scrypt","scrypt-params":{"len":32,"n":262144,"p":1,"r":8,"salt":"hwToMUt2vE7J/Qbk3n1smV6HFaa+ujuBJGRyi8yXbiM="}}`)
	fileInvalidCurve   = []byte(`{"public-key":"AmnZCFEONVvrHVvy34Ep5bZAHhlpiR6AFqCyMAc5u7AG","cipher-text":"0UvywyXU0MBHb6hvtgQ7PhjYbTOWuceAbA4OUaIR68cUlAuVHG7qm/rQeQOSIjRzs6F7V0bXXEY300aE6CD8NDtnlWV6NB53","curve-type":"invalid","id":"4b1ff06f-6b7a-483d-ba84-1d2feaffd749","timestamp":"2019-10-27T21:48:13.000554552Z","version":"dev","storage-cipher":"nacl","kdf":"scrypt","scrypt-params":{"len":32,"n":262144,"p":1,"r":8,"salt":"vPWyjXpUWNmdqOEYMBWyPX0kFKdRHZ205c/WSrl5gac="}}`)
)
