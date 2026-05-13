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

// ApplicationsApplicationRequest is the request body for POST /Applications and PUT /Applications/{id}.
// For PUT, set the Id field. For POST, Id is optional (ignored by server).
type ApplicationsApplicationRequest struct {
	Id                 *int32                                      `json:"Id,omitempty"`
	Name               *string                                     `json:"Name,omitempty"`
	OverwriteSchedules *bool                                       `json:"OverwriteSchedules,omitempty"`
	Schedule           *KeyfactorCommonSchedulingKeyfactorSchedule `json:"Schedule,omitempty"`
}

// NewApplicationsApplicationRequest instantiates a new ApplicationsApplicationRequest.
func NewApplicationsApplicationRequest() *ApplicationsApplicationRequest {
	return &ApplicationsApplicationRequest{}
}

// SetId sets the Id field.
func (o *ApplicationsApplicationRequest) SetId(v int32) {
	o.Id = &v
}

// SetName sets the Name field.
func (o *ApplicationsApplicationRequest) SetName(v string) {
	o.Name = &v
}

// SetOverwriteSchedules sets the OverwriteSchedules field.
func (o *ApplicationsApplicationRequest) SetOverwriteSchedules(v bool) {
	o.OverwriteSchedules = &v
}

// SetSchedule sets the Schedule field.
func (o *ApplicationsApplicationRequest) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule) {
	o.Schedule = &v
}

func (o ApplicationsApplicationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationsApplicationRequest) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}
