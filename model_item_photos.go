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

// ItemPhotos struct for ItemPhotos
type ItemPhotos struct {
	// Полноразмерное фото
	F string `json:"f"`
	// Уменьшенное фото
	T string `json:"t"`
}

// NewItemPhotos instantiates a new ItemPhotos object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemPhotos(f string, t string) *ItemPhotos {
	this := ItemPhotos{}
	this.F = f
	this.T = t
	return &this
}

// NewItemPhotosWithDefaults instantiates a new ItemPhotos object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemPhotosWithDefaults() *ItemPhotos {
	this := ItemPhotos{}
	return &this
}

// GetF returns the F field value
func (o *ItemPhotos) GetF() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.F
}

// GetFOk returns a tuple with the F field value
// and a boolean to check if the value has been set.
func (o *ItemPhotos) GetFOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.F, true
}

// SetF sets field value
func (o *ItemPhotos) SetF(v string) {
	o.F = v
}

// GetT returns the T field value
func (o *ItemPhotos) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *ItemPhotos) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *ItemPhotos) SetT(v string) {
	o.T = v
}

func (o ItemPhotos) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["f"] = o.F
	}
	if true {
		toSerialize["t"] = o.T
	}
	return json.Marshal(toSerialize)
}

type NullableItemPhotos struct {
	value *ItemPhotos
	isSet bool
}

func (v NullableItemPhotos) Get() *ItemPhotos {
	return v.value
}

func (v *NullableItemPhotos) Set(val *ItemPhotos) {
	v.value = val
	v.isSet = true
}

func (v NullableItemPhotos) IsSet() bool {
	return v.isSet
}

func (v *NullableItemPhotos) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemPhotos(val *ItemPhotos) *NullableItemPhotos {
	return &NullableItemPhotos{value: val, isSet: true}
}

func (v NullableItemPhotos) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemPhotos) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


