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

package commands

import (
	"testing"
)

func Test_versionCmd(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			"success",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := versionCmd()
			err := cmd.Execute()
			if (err != nil) != tt.wantErr {
				t.Errorf("versionCmd().Execute error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}