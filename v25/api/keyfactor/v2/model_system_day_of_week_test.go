/*
Copyright 2025 Keyfactor
Licensed under the Apache License, Version 2.0 (the "License"); you may
not use this file except in compliance with the License.  You may obtain a
copy of the License at http://www.apache.org/licenses/LICENSE-2.0.  Unless
required by applicable law or agreed to in writing, software distributed
under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
OR CONDITIONS OF ANY KIND, either express or implied. See the License for
the specific language governing permissions and limitations under the
License.
*/

package v2

import (
	"encoding/json"
	"testing"
	"time"
)

// TestSystemDayOfWeek_UnmarshalJSON_IntForm verifies that the original
// generated wire form (a JSON integer) still deserializes correctly.
func TestSystemDayOfWeek_UnmarshalJSON_IntForm(t *testing.T) {
	var got SystemDayOfWeek
	if err := json.Unmarshal([]byte(`1`), &got); err != nil {
		t.Fatalf("unexpected error unmarshaling int form: %v", err)
	}
	if got != SYSTEMDAYOFWEEK_Monday {
		t.Errorf("expected %v, got %v", SYSTEMDAYOFWEEK_Monday, got)
	}
}

// TestSystemDayOfWeek_UnmarshalJSON_StringForm covers the regression from
// GitHub issue #185: Keyfactor Command v25.5 serializes WeeklyModel.Days as
// day-name strings (e.g. "Monday") rather than integers.
func TestSystemDayOfWeek_UnmarshalJSON_StringForm(t *testing.T) {
	var got SystemDayOfWeek
	if err := json.Unmarshal([]byte(`"Monday"`), &got); err != nil {
		t.Fatalf("unexpected error unmarshaling string form: %v", err)
	}
	if got != SYSTEMDAYOFWEEK_Monday {
		t.Errorf("expected %v, got %v", SYSTEMDAYOFWEEK_Monday, got)
	}
}

// TestSystemDayOfWeek_UnmarshalJSON_InvalidString verifies a malformed day
// name still produces a clear error rather than silently defaulting.
func TestSystemDayOfWeek_UnmarshalJSON_InvalidString(t *testing.T) {
	var got SystemDayOfWeek
	err := json.Unmarshal([]byte(`"Funday"`), &got)
	if err == nil {
		t.Fatalf("expected error for invalid day-name string, got nil (value=%v)", got)
	}
}

// TestSystemDayOfWeek_UnmarshalJSON_OutOfRangeInt verifies an out-of-range
// integer still produces a clear error rather than silently accepting it.
func TestSystemDayOfWeek_UnmarshalJSON_OutOfRangeInt(t *testing.T) {
	var got SystemDayOfWeek
	err := json.Unmarshal([]byte(`42`), &got)
	if err == nil {
		t.Fatalf("expected error for out-of-range int, got nil (value=%v)", got)
	}
}

// TestWeeklyModel_UnmarshalJSON_DayNameStrings is the full round-trip
// regression test for issue #185: a Weekly-shaped schedule payload as
// returned by GET /CertificateAuthority on a v25.5 Command instance must
// deserialize into KeyfactorCommonSchedulingModelsWeeklyModel without error.
func TestWeeklyModel_UnmarshalJSON_DayNameStrings(t *testing.T) {
	payload := `{"Days":["Monday","Friday"],"Time":"2000-01-01T07:00:00Z"}`

	var model KeyfactorCommonSchedulingModelsWeeklyModel
	if err := json.Unmarshal([]byte(payload), &model); err != nil {
		t.Fatalf("unexpected error unmarshaling WeeklyModel with day-name strings: %v", err)
	}

	wantDays := []SystemDayOfWeek{SYSTEMDAYOFWEEK_Monday, SYSTEMDAYOFWEEK_Friday}
	if len(model.Days) != len(wantDays) {
		t.Fatalf("expected %d days, got %d (%v)", len(wantDays), len(model.Days), model.Days)
	}
	for i, want := range wantDays {
		if model.Days[i] != want {
			t.Errorf("Days[%d]: expected %v, got %v", i, want, model.Days[i])
		}
	}

	wantTime, err := time.Parse(time.RFC3339, "2000-01-01T07:00:00Z")
	if err != nil {
		t.Fatalf("failed to parse expected time: %v", err)
	}
	if model.Time == nil || !model.Time.Equal(wantTime) {
		t.Errorf("Time: expected %v, got %v", wantTime, model.Time)
	}
}

// TestWeeklyModel_UnmarshalJSON_DayIndexInts verifies the pre-v25.5 integer
// wire form of WeeklyModel.Days still round-trips correctly, guarding
// against a regression in the other direction.
func TestWeeklyModel_UnmarshalJSON_DayIndexInts(t *testing.T) {
	payload := `{"Days":[1,5],"Time":"2000-01-01T07:00:00Z"}`

	var model KeyfactorCommonSchedulingModelsWeeklyModel
	if err := json.Unmarshal([]byte(payload), &model); err != nil {
		t.Fatalf("unexpected error unmarshaling WeeklyModel with int days: %v", err)
	}

	wantDays := []SystemDayOfWeek{SYSTEMDAYOFWEEK_Monday, SYSTEMDAYOFWEEK_Friday}
	if len(model.Days) != len(wantDays) {
		t.Fatalf("expected %d days, got %d (%v)", len(wantDays), len(model.Days), model.Days)
	}
	for i, want := range wantDays {
		if model.Days[i] != want {
			t.Errorf("Days[%d]: expected %v, got %v", i, want, model.Days[i])
		}
	}
}
