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

package ethereum

import (
	"reflect"
	"testing"
)

func TestNetworks(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{
			"success",
			[]string{Goerli, Kovan, Mainnet, Rinkeby, Ropsten},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Networks(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Networks() = %v, want %v", got, tt.want)
			}
		})
	}
}
