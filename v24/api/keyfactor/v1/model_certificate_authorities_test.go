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

package v1

import (
	"testing"
)

// TestCARequestFields_CleanupAndEnrollment verifies that the 5 new fields
// (UseForEnrollment, CertificateCleanupEnabled, DeleteWithArchivedKey,
// TimeAfterExpiration, TimeAfterExpirationUnits) on the request struct can
// be set via setters and read back via getters. This is a compile-time +
// nil-safety regression test: if any field or method is removed, the test
// fails to compile.
func TestCARequestFields_CleanupAndEnrollment(t *testing.T) {
	req := NewCertificateAuthoritiesCertificateAuthorityRequest()

	// UseForEnrollment — *bool (plain pointer)
	req.SetUseForEnrollment(true)
	if got := req.GetUseForEnrollment(); got != true {
		t.Errorf("UseForEnrollment: expected true, got %v", got)
	}
	if !req.HasUseForEnrollment() {
		t.Error("HasUseForEnrollment: expected true after Set")
	}

	// CertificateCleanupEnabled — NullableBool
	req.SetCertificateCleanupEnabled(true)
	if got := req.GetCertificateCleanupEnabled(); got != true {
		t.Errorf("CertificateCleanupEnabled: expected true, got %v", got)
	}
	// Verify IsSet via the underlying NullableBool (no Has* on request)
	if !req.CertificateCleanupEnabled.IsSet() {
		t.Error("CertificateCleanupEnabled.IsSet: expected true after Set")
	}

	// DeleteWithArchivedKey — NullableBool
	req.SetDeleteWithArchivedKey(false)
	if got := req.GetDeleteWithArchivedKey(); got != false {
		t.Errorf("DeleteWithArchivedKey: expected false, got %v", got)
	}
	if !req.DeleteWithArchivedKey.IsSet() {
		t.Error("DeleteWithArchivedKey.IsSet: expected true after Set")
	}

	// TimeAfterExpiration — NullableInt32
	req.SetTimeAfterExpiration(90)
	if got := req.GetTimeAfterExpiration(); got != 90 {
		t.Errorf("TimeAfterExpiration: expected 90, got %v", got)
	}
	if !req.TimeAfterExpiration.IsSet() {
		t.Error("TimeAfterExpiration.IsSet: expected true after Set")
	}

	// TimeAfterExpirationUnits — *CSSCMSDataModelEnumsCertificateCleanupTimeUnits
	units := CSSCMSDATAMODELENUMSCERTIFICATECLEANUPTIMEUNITS__1
	req.SetTimeAfterExpirationUnits(units)
	if got := req.GetTimeAfterExpirationUnits(); got != units {
		t.Errorf("TimeAfterExpirationUnits: expected %v, got %v", units, got)
	}
}

// TestCAResponseFields_CleanupAndEnrollment verifies the same 5 fields on
// the response struct, including the Has* methods that exist on response
// but not on request.
func TestCAResponseFields_CleanupAndEnrollment(t *testing.T) {
	resp := NewCertificateAuthoritiesCertificateAuthorityResponse()

	// UseForEnrollment — *bool
	resp.SetUseForEnrollment(true)
	if got := resp.GetUseForEnrollment(); got != true {
		t.Errorf("UseForEnrollment: expected true, got %v", got)
	}
	if !resp.HasUseForEnrollment() {
		t.Error("HasUseForEnrollment: expected true after Set")
	}
	if val, ok := resp.GetUseForEnrollmentOk(); !ok || val == nil || *val != true {
		t.Error("GetUseForEnrollmentOk: expected (true, true)")
	}

	// CertificateCleanupEnabled — NullableBool
	resp.SetCertificateCleanupEnabled(true)
	if got := resp.GetCertificateCleanupEnabled(); got != true {
		t.Errorf("CertificateCleanupEnabled: expected true, got %v", got)
	}
	if !resp.HasCertificateCleanupEnabled() {
		t.Error("HasCertificateCleanupEnabled: expected true after Set")
	}
	if val, ok := resp.GetCertificateCleanupEnabledOk(); !ok || val == nil || *val != true {
		t.Error("GetCertificateCleanupEnabledOk: expected (true, true)")
	}

	// DeleteWithArchivedKey — NullableBool
	resp.SetDeleteWithArchivedKey(false)
	if got := resp.GetDeleteWithArchivedKey(); got != false {
		t.Errorf("DeleteWithArchivedKey: expected false, got %v", got)
	}
	if !resp.HasDeleteWithArchivedKey() {
		t.Error("HasDeleteWithArchivedKey: expected true after Set")
	}
	if val, ok := resp.GetDeleteWithArchivedKeyOk(); !ok || val == nil || *val != false {
		t.Error("GetDeleteWithArchivedKeyOk: expected (false, true)")
	}

	// TimeAfterExpiration — NullableInt32
	resp.SetTimeAfterExpiration(30)
	if got := resp.GetTimeAfterExpiration(); got != 30 {
		t.Errorf("TimeAfterExpiration: expected 30, got %v", got)
	}
	if !resp.HasTimeAfterExpiration() {
		t.Error("HasTimeAfterExpiration: expected true after Set")
	}
	if val, ok := resp.GetTimeAfterExpirationOk(); !ok || val == nil || *val != 30 {
		t.Error("GetTimeAfterExpirationOk: expected (30, true)")
	}

	// TimeAfterExpirationUnits — *CSSCMSDataModelEnumsCertificateCleanupTimeUnits
	units := CSSCMSDATAMODELENUMSCERTIFICATECLEANUPTIMEUNITS__2
	resp.SetTimeAfterExpirationUnits(units)
	if got := resp.GetTimeAfterExpirationUnits(); got != units {
		t.Errorf("TimeAfterExpirationUnits: expected %v, got %v", units, got)
	}
	if !resp.HasTimeAfterExpirationUnits() {
		t.Error("HasTimeAfterExpirationUnits: expected true after Set")
	}
	if val, ok := resp.GetTimeAfterExpirationUnitsOk(); !ok || val == nil || *val != units {
		t.Error("GetTimeAfterExpirationUnitsOk: expected (units, true)")
	}
}

// TestCARequestFields_NilSafety verifies that calling getters on a
// zero-value request struct does not panic and returns Go zero values.
func TestCARequestFields_NilSafety(t *testing.T) {
	req := CertificateAuthoritiesCertificateAuthorityRequest{}

	// UseForEnrollment — should return false (zero bool)
	if got := req.GetUseForEnrollment(); got != false {
		t.Errorf("UseForEnrollment zero: expected false, got %v", got)
	}
	if req.HasUseForEnrollment() {
		t.Error("HasUseForEnrollment zero: expected false")
	}

	// CertificateCleanupEnabled — NullableBool unset, should return false
	if got := req.GetCertificateCleanupEnabled(); got != false {
		t.Errorf("CertificateCleanupEnabled zero: expected false, got %v", got)
	}
	if req.CertificateCleanupEnabled.IsSet() {
		t.Error("CertificateCleanupEnabled.IsSet zero: expected false")
	}

	// DeleteWithArchivedKey — NullableBool unset, should return false
	if got := req.GetDeleteWithArchivedKey(); got != false {
		t.Errorf("DeleteWithArchivedKey zero: expected false, got %v", got)
	}
	if req.DeleteWithArchivedKey.IsSet() {
		t.Error("DeleteWithArchivedKey.IsSet zero: expected false")
	}

	// TimeAfterExpiration — NullableInt32 unset, should return 0
	if got := req.GetTimeAfterExpiration(); got != 0 {
		t.Errorf("TimeAfterExpiration zero: expected 0, got %v", got)
	}
	if req.TimeAfterExpiration.IsSet() {
		t.Error("TimeAfterExpiration.IsSet zero: expected false")
	}

	// TimeAfterExpirationUnits — nil pointer, should return zero enum
	if got := req.GetTimeAfterExpirationUnits(); got != 0 {
		t.Errorf("TimeAfterExpirationUnits zero: expected 0, got %v", got)
	}

	// Also test nil receiver safety
	var nilReq *CertificateAuthoritiesCertificateAuthorityRequest
	if got := nilReq.GetUseForEnrollment(); got != false {
		t.Errorf("nil receiver UseForEnrollment: expected false, got %v", got)
	}
	if got := nilReq.GetCertificateCleanupEnabled(); got != false {
		t.Errorf("nil receiver CertificateCleanupEnabled: expected false, got %v", got)
	}
	if got := nilReq.GetDeleteWithArchivedKey(); got != false {
		t.Errorf("nil receiver DeleteWithArchivedKey: expected false, got %v", got)
	}
	if got := nilReq.GetTimeAfterExpiration(); got != 0 {
		t.Errorf("nil receiver TimeAfterExpiration: expected 0, got %v", got)
	}
	if got := nilReq.GetTimeAfterExpirationUnits(); got != 0 {
		t.Errorf("nil receiver TimeAfterExpirationUnits: expected 0, got %v", got)
	}
}

// TestCAResponseFields_NilSafety verifies that calling getters on a
// zero-value response struct does not panic and returns Go zero values.
func TestCAResponseFields_NilSafety(t *testing.T) {
	resp := CertificateAuthoritiesCertificateAuthorityResponse{}

	// UseForEnrollment — should return false
	if got := resp.GetUseForEnrollment(); got != false {
		t.Errorf("UseForEnrollment zero: expected false, got %v", got)
	}
	if resp.HasUseForEnrollment() {
		t.Error("HasUseForEnrollment zero: expected false")
	}

	// CertificateCleanupEnabled — NullableBool unset
	if got := resp.GetCertificateCleanupEnabled(); got != false {
		t.Errorf("CertificateCleanupEnabled zero: expected false, got %v", got)
	}
	if resp.HasCertificateCleanupEnabled() {
		t.Error("HasCertificateCleanupEnabled zero: expected false")
	}

	// DeleteWithArchivedKey — NullableBool unset
	if got := resp.GetDeleteWithArchivedKey(); got != false {
		t.Errorf("DeleteWithArchivedKey zero: expected false, got %v", got)
	}
	if resp.HasDeleteWithArchivedKey() {
		t.Error("HasDeleteWithArchivedKey zero: expected false")
	}

	// TimeAfterExpiration — NullableInt32 unset
	if got := resp.GetTimeAfterExpiration(); got != 0 {
		t.Errorf("TimeAfterExpiration zero: expected 0, got %v", got)
	}
	if resp.HasTimeAfterExpiration() {
		t.Error("HasTimeAfterExpiration zero: expected false")
	}

	// TimeAfterExpirationUnits — nil pointer
	if got := resp.GetTimeAfterExpirationUnits(); got != 0 {
		t.Errorf("TimeAfterExpirationUnits zero: expected 0, got %v", got)
	}
	if resp.HasTimeAfterExpirationUnits() {
		t.Error("HasTimeAfterExpirationUnits zero: expected false")
	}

	// Also test nil receiver safety
	var nilResp *CertificateAuthoritiesCertificateAuthorityResponse
	if got := nilResp.GetUseForEnrollment(); got != false {
		t.Errorf("nil receiver UseForEnrollment: expected false, got %v", got)
	}
	if got := nilResp.GetCertificateCleanupEnabled(); got != false {
		t.Errorf("nil receiver CertificateCleanupEnabled: expected false, got %v", got)
	}
	if got := nilResp.GetDeleteWithArchivedKey(); got != false {
		t.Errorf("nil receiver DeleteWithArchivedKey: expected false, got %v", got)
	}
	if got := nilResp.GetTimeAfterExpiration(); got != 0 {
		t.Errorf("nil receiver TimeAfterExpiration: expected 0, got %v", got)
	}
	if got := nilResp.GetTimeAfterExpirationUnits(); got != 0 {
		t.Errorf("nil receiver TimeAfterExpirationUnits: expected 0, got %v", got)
	}
}
