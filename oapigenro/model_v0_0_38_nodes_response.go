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

// checks if the V0038NodesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0038NodesResponse{}

// V0038NodesResponse struct for V0038NodesResponse
type V0038NodesResponse struct {
	Meta *V0038Meta `json:"meta,omitempty"`
	// slurm errors
	Errors []V0038Error `json:"errors,omitempty"`
	// nodes info
	Nodes []V0038Node `json:"nodes,omitempty"`
}

// NewV0038NodesResponse instantiates a new V0038NodesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0038NodesResponse() *V0038NodesResponse {
	this := V0038NodesResponse{}
	return &this
}

// NewV0038NodesResponseWithDefaults instantiates a new V0038NodesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0038NodesResponseWithDefaults() *V0038NodesResponse {
	this := V0038NodesResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0038NodesResponse) GetMeta() V0038Meta {
	if o == nil || IsNil(o.Meta) {
		var ret V0038Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038NodesResponse) GetMetaOk() (*V0038Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0038NodesResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V0038Meta and assigns it to the Meta field.
func (o *V0038NodesResponse) SetMeta(v V0038Meta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0038NodesResponse) GetErrors() []V0038Error {
	if o == nil || IsNil(o.Errors) {
		var ret []V0038Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038NodesResponse) GetErrorsOk() ([]V0038Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0038NodesResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0038Error and assigns it to the Errors field.
func (o *V0038NodesResponse) SetErrors(v []V0038Error) {
	o.Errors = v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *V0038NodesResponse) GetNodes() []V0038Node {
	if o == nil || IsNil(o.Nodes) {
		var ret []V0038Node
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038NodesResponse) GetNodesOk() ([]V0038Node, bool) {
	if o == nil || IsNil(o.Nodes) {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *V0038NodesResponse) HasNodes() bool {
	if o != nil && !IsNil(o.Nodes) {
		return true
	}

	return false
}

// SetNodes gets a reference to the given []V0038Node and assigns it to the Nodes field.
func (o *V0038NodesResponse) SetNodes(v []V0038Node) {
	o.Nodes = v
}

func (o V0038NodesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0038NodesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Nodes) {
		toSerialize["nodes"] = o.Nodes
	}
	return toSerialize, nil
}

type NullableV0038NodesResponse struct {
	value *V0038NodesResponse
	isSet bool
}

func (v NullableV0038NodesResponse) Get() *V0038NodesResponse {
	return v.value
}

func (v *NullableV0038NodesResponse) Set(val *V0038NodesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV0038NodesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV0038NodesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0038NodesResponse(val *V0038NodesResponse) *NullableV0038NodesResponse {
	return &NullableV0038NodesResponse{value: val, isSet: true}
}

func (v NullableV0038NodesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0038NodesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


