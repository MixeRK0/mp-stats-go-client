/*
mpstats

MPStats API

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// TplsRequestBody struct for TplsRequestBody
type TplsRequestBody struct {
	Tpls *string `json:"tpls,omitempty"`
}

// NewTplsRequestBody instantiates a new TplsRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTplsRequestBody() *TplsRequestBody {
	this := TplsRequestBody{}
	return &this
}

// NewTplsRequestBodyWithDefaults instantiates a new TplsRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTplsRequestBodyWithDefaults() *TplsRequestBody {
	this := TplsRequestBody{}
	return &this
}

// GetTpls returns the Tpls field value if set, zero value otherwise.
func (o *TplsRequestBody) GetTpls() string {
	if o == nil || o.Tpls == nil {
		var ret string
		return ret
	}
	return *o.Tpls
}

// GetTplsOk returns a tuple with the Tpls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TplsRequestBody) GetTplsOk() (*string, bool) {
	if o == nil || o.Tpls == nil {
		return nil, false
	}
	return o.Tpls, true
}

// HasTpls returns a boolean if a field has been set.
func (o *TplsRequestBody) HasTpls() bool {
	if o != nil && o.Tpls != nil {
		return true
	}

	return false
}

// SetTpls gets a reference to the given string and assigns it to the Tpls field.
func (o *TplsRequestBody) SetTpls(v string) {
	o.Tpls = &v
}

func (o TplsRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tpls != nil {
		toSerialize["tpls"] = o.Tpls
	}
	return json.Marshal(toSerialize)
}

type NullableTplsRequestBody struct {
	value *TplsRequestBody
	isSet bool
}

func (v NullableTplsRequestBody) Get() *TplsRequestBody {
	return v.value
}

func (v *NullableTplsRequestBody) Set(val *TplsRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableTplsRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableTplsRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTplsRequestBody(val *TplsRequestBody) *NullableTplsRequestBody {
	return &NullableTplsRequestBody{value: val, isSet: true}
}

func (v NullableTplsRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTplsRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

