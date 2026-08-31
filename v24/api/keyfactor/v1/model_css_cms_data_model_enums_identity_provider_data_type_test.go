/*
Copyright 2026 Keyfactor
Licensed under the Apache License, Version 2.0 (the "License"); you may
not use this file except in compliance with the License.  You may obtain a
copy of the License at http://www.apache.org/licenses/LICENSE-2.0.  Unless
required by applicable law or agreed to in writing, software distributed
under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
OR CONDITIONS OF ANY KIND, either express or implied. See the License for
the specific language governing permissions and limitations under the
License.
*/

package v1

import (
	"encoding/json"
	"strconv"
	"testing"
)

// TestCSSCMSDataModelEnumsIdentityProviderDataType_UnmarshalJSON_KnownValues
// verifies that all wire values Command is known to emit for this enum
// deserialize without error. Value 4 regresses a bug where GET
// /IdentityProviders on an OAuth-enabled Command instance (observed against
// an Authentik-backed identity provider) returned a Parameters[].DataType of
// 4, but this generated enum only allowed [1,2,3], causing the entire
// []IdentityProviderIdentityProviderGetResponse response to fail to decode
// with "4 is not a valid CSSCMSDataModelEnumsIdentityProviderDataType". The
// v25 module's canonical swagger-generated enum already includes 4; this
// hand-edit brings v24 in line with it.
func TestCSSCMSDataModelEnumsIdentityProviderDataType_UnmarshalJSON_KnownValues(t *testing.T) {
	for _, want := range AllowedCSSCMSDataModelEnumsIdentityProviderDataTypeEnumValues {
		var got CSSCMSDataModelEnumsIdentityProviderDataType
		payload := []byte(strconv.Itoa(int(want)))
		if err := json.Unmarshal(payload, &got); err != nil {
			t.Errorf("unexpected error unmarshaling value %d: %v", want, err)
		}
		if got != want {
			t.Errorf("expected %v, got %v", want, got)
		}
	}
}

// TestCSSCMSDataModelEnumsIdentityProviderDataType_UnmarshalJSON_OutOfRange
// verifies that a value outside the known set still produces a clear error
// instead of silently accepting it.
func TestCSSCMSDataModelEnumsIdentityProviderDataType_UnmarshalJSON_OutOfRange(t *testing.T) {
	var got CSSCMSDataModelEnumsIdentityProviderDataType
	err := json.Unmarshal([]byte("42"), &got)
	if err == nil {
		t.Fatalf("expected error for out-of-range value, got nil (value=%v)", got)
	}
}
