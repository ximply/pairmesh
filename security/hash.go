// Copyright 2021 PairMesh, Inc.
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

package security

import (
	"encoding/base64"

	"golang.org/x/crypto/blake2s"
)

// Hash returns the hash value of the password with salt
func Hash(password, salt string) string {
	sum := blake2s.Sum256([]byte(salt + password))
	return base64.RawStdEncoding.EncodeToString(sum[:])
}