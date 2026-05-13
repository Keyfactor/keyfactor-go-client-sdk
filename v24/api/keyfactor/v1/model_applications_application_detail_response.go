/*
Copyright 2025 Keyfactor
Licensed under the Apache License, Version 2.0 (the "License"); you may
not use this file except in compliance with the License.  You may obtain a
copy of the License at http://www.apache.org/licenses/LICENSE-2.0.  Unless
required by applicable law or agreed to in writing, software distributed
under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
OR CONDITIONS OF ANY KIND, either express or implied. See the License for
thespecific language governing permissions and limitations under the
License.

Keyfactor Command Version: 25+

API version: 1
*/

package v1

import (
	"encoding/json"
)

// ApplicationsCertStoreRef is a certificate store reference within an application.
type ApplicationsCertStoreRef struct {
	Id *string `json:"Id,omitempty"` // Store GUID (UUID)
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplicationsCertStoreRef) GetId() string {
	if o == nil || o.Id == nil {
		return ""
	}
	return *o.Id
}

// ApplicationsApplicationDetailResponse is the full application detail from GET /Applications/{id}.
type ApplicationsApplicationDetailResponse struct {
	Id                 *int32                                      `json:"Id,omitempty"`
	Name               *string                                     `json:"Name,omitempty"`
	OverwriteSchedules *bool                                       `json:"OverwriteSchedules,omitempty"`
	Schedule           *KeyfactorCommonSchedulingKeyfactorSchedule `json:"Schedule,omitempty"`
	CertificateStores  []ApplicationsCertStoreRef                  `json:"CertificateStores,omitempty"`
}

// NewApplicationsApplicationDetailResponse instantiates a new ApplicationsApplicationDetailResponse.
func NewApplicationsApplicationDetailResponse() *ApplicationsApplicationDetailResponse {
	return &ApplicationsApplicationDetailResponse{}
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplicationsApplicationDetailResponse) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApplicationsApplicationDetailResponse) GetName() string {
	if o == nil || o.Name == nil {
		return ""
	}
	return *o.Name
}

// GetOverwriteSchedules returns the OverwriteSchedules field value if set, zero value otherwise.
func (o *ApplicationsApplicationDetailResponse) GetOverwriteSchedules() bool {
	if o == nil || o.OverwriteSchedules == nil {
		return false
	}
	return *o.OverwriteSchedules
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *ApplicationsApplicationDetailResponse) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule {
	if o == nil || o.Schedule == nil {
		var ret KeyfactorCommonSchedulingKeyfactorSchedule
		return ret
	}
	return *o.Schedule
}

// GetCertificateStores returns the CertificateStores slice.
func (o *ApplicationsApplicationDetailResponse) GetCertificateStores() []ApplicationsCertStoreRef {
	if o == nil {
		return nil
	}
	return o.CertificateStores
}

func (o ApplicationsApplicationDetailResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationsApplicationDetailResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["Id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OverwriteSchedules != nil {
		toSerialize["OverwriteSchedules"] = o.OverwriteSchedules
	}
	if o.Schedule != nil {
		toSerialize["Schedule"] = o.Schedule
	}
	if o.CertificateStores != nil {
		toSerialize["CertificateStores"] = o.CertificateStores
	}
	return toSerialize, nil
}
