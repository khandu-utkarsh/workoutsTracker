/*
TrackBot App API

Comprehensive API for the TrackBot Application helping in logging workouts, exercises and nutrition using AI.  This API provides endpoints for: - User management - Workout tracking and management - Exercise logging (cardio and weight training) - AI conversation and messaging  All timestamps are in ISO 8601 format (UTC). 

API version: 1.0.0
Contact: utkarshkhandelwal2011@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_models

import (
	"encoding/json"
)

// checks if the CreateWorkoutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateWorkoutRequest{}

// CreateWorkoutRequest Request to create a new workout.
type CreateWorkoutRequest struct {
	// ID of the user creating the workout.
	UserId *int64 `json:"user_id,omitempty"`
}

// NewCreateWorkoutRequest instantiates a new CreateWorkoutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWorkoutRequest() *CreateWorkoutRequest {
	this := CreateWorkoutRequest{}
	return &this
}

// NewCreateWorkoutRequestWithDefaults instantiates a new CreateWorkoutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWorkoutRequestWithDefaults() *CreateWorkoutRequest {
	this := CreateWorkoutRequest{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *CreateWorkoutRequest) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWorkoutRequest) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *CreateWorkoutRequest) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *CreateWorkoutRequest) SetUserId(v int64) {
	o.UserId = &v
}

func (o CreateWorkoutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateWorkoutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullableCreateWorkoutRequest struct {
	value *CreateWorkoutRequest
	isSet bool
}

func (v NullableCreateWorkoutRequest) Get() *CreateWorkoutRequest {
	return v.value
}

func (v *NullableCreateWorkoutRequest) Set(val *CreateWorkoutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWorkoutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWorkoutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWorkoutRequest(val *CreateWorkoutRequest) *NullableCreateWorkoutRequest {
	return &NullableCreateWorkoutRequest{value: val, isSet: true}
}

func (v NullableCreateWorkoutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWorkoutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


