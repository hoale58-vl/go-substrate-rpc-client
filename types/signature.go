// Go Substrate RPC Client (GSRPC) provides APIs and types around Polkadot and any Substrate-based chain RPC calls
//
// Copyright 2019 Centrifuge GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import "fmt"

// Signature is a H512
type Signature [65]byte

// NewSignature creates a new Signature type
func NewSignature(b []byte) Signature {
	h := Signature{}
	copy(h[:], b)
	return h
}

// Hex returns a hex string representation of the value (not of the encoded value)
func (h Signature) Hex() string {
	return fmt.Sprintf("%#x", h[:])
}
