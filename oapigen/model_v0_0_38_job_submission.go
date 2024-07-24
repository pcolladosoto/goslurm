/*
Slurm Rest API

API to access and control Slurm.

API version: 0.0.38
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package oapigen

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the V0038JobSubmission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0038JobSubmission{}

// V0038JobSubmission struct for V0038JobSubmission
type V0038JobSubmission struct {
	// Executable script (full contents) to run in batch step
	Script string `json:"script"`
	Job *V0038JobProperties `json:"job,omitempty"`
	// Properties of an HetJob
	Jobs []V0038JobProperties `json:"jobs,omitempty"`
}

type _V0038JobSubmission V0038JobSubmission

// NewV0038JobSubmission instantiates a new V0038JobSubmission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0038JobSubmission(script string) *V0038JobSubmission {
	this := V0038JobSubmission{}
	this.Script = script
	return &this
}

// NewV0038JobSubmissionWithDefaults instantiates a new V0038JobSubmission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0038JobSubmissionWithDefaults() *V0038JobSubmission {
	this := V0038JobSubmission{}
	return &this
}

// GetScript returns the Script field value
func (o *V0038JobSubmission) GetScript() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Script
}

// GetScriptOk returns a tuple with the Script field value
// and a boolean to check if the value has been set.
func (o *V0038JobSubmission) GetScriptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Script, true
}

// SetScript sets field value
func (o *V0038JobSubmission) SetScript(v string) {
	o.Script = v
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *V0038JobSubmission) GetJob() V0038JobProperties {
	if o == nil || IsNil(o.Job) {
		var ret V0038JobProperties
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038JobSubmission) GetJobOk() (*V0038JobProperties, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *V0038JobSubmission) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given V0038JobProperties and assigns it to the Job field.
func (o *V0038JobSubmission) SetJob(v V0038JobProperties) {
	o.Job = &v
}

// GetJobs returns the Jobs field value if set, zero value otherwise.
func (o *V0038JobSubmission) GetJobs() []V0038JobProperties {
	if o == nil || IsNil(o.Jobs) {
		var ret []V0038JobProperties
		return ret
	}
	return o.Jobs
}

// GetJobsOk returns a tuple with the Jobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038JobSubmission) GetJobsOk() ([]V0038JobProperties, bool) {
	if o == nil || IsNil(o.Jobs) {
		return nil, false
	}
	return o.Jobs, true
}

// HasJobs returns a boolean if a field has been set.
func (o *V0038JobSubmission) HasJobs() bool {
	if o != nil && !IsNil(o.Jobs) {
		return true
	}

	return false
}

// SetJobs gets a reference to the given []V0038JobProperties and assigns it to the Jobs field.
func (o *V0038JobSubmission) SetJobs(v []V0038JobProperties) {
	o.Jobs = v
}

func (o V0038JobSubmission) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0038JobSubmission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["script"] = o.Script
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	if !IsNil(o.Jobs) {
		toSerialize["jobs"] = o.Jobs
	}
	return toSerialize, nil
}

func (o *V0038JobSubmission) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"script",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varV0038JobSubmission := _V0038JobSubmission{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0038JobSubmission)

	if err != nil {
		return err
	}

	*o = V0038JobSubmission(varV0038JobSubmission)

	return err
}

type NullableV0038JobSubmission struct {
	value *V0038JobSubmission
	isSet bool
}

func (v NullableV0038JobSubmission) Get() *V0038JobSubmission {
	return v.value
}

func (v *NullableV0038JobSubmission) Set(val *V0038JobSubmission) {
	v.value = val
	v.isSet = true
}

func (v NullableV0038JobSubmission) IsSet() bool {
	return v.isSet
}

func (v *NullableV0038JobSubmission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0038JobSubmission(val *V0038JobSubmission) *NullableV0038JobSubmission {
	return &NullableV0038JobSubmission{value: val, isSet: true}
}

func (v NullableV0038JobSubmission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0038JobSubmission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


