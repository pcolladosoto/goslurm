/*
Slurm Rest API RO

API to access Slurm. Only GET requests are implemented.

API version: 0.0.38
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package oapigen

import (
	"encoding/json"
)

// checks if the V0038Diag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0038Diag{}

// V0038Diag struct for V0038Diag
type V0038Diag struct {
	Meta *V0038Meta `json:"meta,omitempty"`
	// slurm errors
	Errors []V0038Error `json:"errors,omitempty"`
	Statistics *V0038DiagStatistics `json:"statistics,omitempty"`
}

// NewV0038Diag instantiates a new V0038Diag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0038Diag() *V0038Diag {
	this := V0038Diag{}
	return &this
}

// NewV0038DiagWithDefaults instantiates a new V0038Diag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0038DiagWithDefaults() *V0038Diag {
	this := V0038Diag{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0038Diag) GetMeta() V0038Meta {
	if o == nil || IsNil(o.Meta) {
		var ret V0038Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Diag) GetMetaOk() (*V0038Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0038Diag) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V0038Meta and assigns it to the Meta field.
func (o *V0038Diag) SetMeta(v V0038Meta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0038Diag) GetErrors() []V0038Error {
	if o == nil || IsNil(o.Errors) {
		var ret []V0038Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Diag) GetErrorsOk() ([]V0038Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0038Diag) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0038Error and assigns it to the Errors field.
func (o *V0038Diag) SetErrors(v []V0038Error) {
	o.Errors = v
}

// GetStatistics returns the Statistics field value if set, zero value otherwise.
func (o *V0038Diag) GetStatistics() V0038DiagStatistics {
	if o == nil || IsNil(o.Statistics) {
		var ret V0038DiagStatistics
		return ret
	}
	return *o.Statistics
}

// GetStatisticsOk returns a tuple with the Statistics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038Diag) GetStatisticsOk() (*V0038DiagStatistics, bool) {
	if o == nil || IsNil(o.Statistics) {
		return nil, false
	}
	return o.Statistics, true
}

// HasStatistics returns a boolean if a field has been set.
func (o *V0038Diag) HasStatistics() bool {
	if o != nil && !IsNil(o.Statistics) {
		return true
	}

	return false
}

// SetStatistics gets a reference to the given V0038DiagStatistics and assigns it to the Statistics field.
func (o *V0038Diag) SetStatistics(v V0038DiagStatistics) {
	o.Statistics = &v
}

func (o V0038Diag) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0038Diag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Statistics) {
		toSerialize["statistics"] = o.Statistics
	}
	return toSerialize, nil
}

type NullableV0038Diag struct {
	value *V0038Diag
	isSet bool
}

func (v NullableV0038Diag) Get() *V0038Diag {
	return v.value
}

func (v *NullableV0038Diag) Set(val *V0038Diag) {
	v.value = val
	v.isSet = true
}

func (v NullableV0038Diag) IsSet() bool {
	return v.isSet
}

func (v *NullableV0038Diag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0038Diag(val *V0038Diag) *NullableV0038Diag {
	return &NullableV0038Diag{value: val, isSet: true}
}

func (v NullableV0038Diag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0038Diag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


