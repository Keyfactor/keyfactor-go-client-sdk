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

// ApplicationsApplicationListResponse represents one application entry in the GET /Applications list response.
// The Schedule field is returned as a cron expression string by the list endpoint.
type ApplicationsApplicationListResponse struct {
	Id       *int32  `json:"Id,omitempty"`
	Name     *string `json:"Name,omitempty"`
	Schedule *string `json:"Schedule,omitempty"` // cron expression string
}

// NewApplicationsApplicationListResponse instantiates a new ApplicationsApplicationListResponse.
func NewApplicationsApplicationListResponse() *ApplicationsApplicationListResponse {
	return &ApplicationsApplicationListResponse{}
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplicationsApplicationListResponse) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApplicationsApplicationListResponse) GetName() string {
	if o == nil || o.Name == nil {
		return ""
	}
	return *o.Name
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *ApplicationsApplicationListResponse) GetSchedule() string {
	if o == nil || o.Schedule == nil {
		return ""
	}
	return *o.Schedule
}

func (o ApplicationsApplicationListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationsApplicationListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["Id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Schedule != nil {
		toSerialize["Schedule"] = o.Schedule
	}
	return toSerialize, nil
}
