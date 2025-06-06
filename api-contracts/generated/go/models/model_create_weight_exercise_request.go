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
	"bytes"
	"fmt"
)

// checks if the CreateWeightExerciseRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateWeightExerciseRequest{}

// CreateWeightExerciseRequest struct for CreateWeightExerciseRequest
type CreateWeightExerciseRequest struct {
	// Name of the exercise
	Name string `json:"name"`
	Type string `json:"type"`
	// Additional notes
	Notes *string `json:"notes,omitempty"`
	// Number of sets
	Sets int32 `json:"sets"`
	// Reps per set
	Reps int32 `json:"reps"`
	// Weight in kg
	Weight float32 `json:"weight"`
}

type _CreateWeightExerciseRequest CreateWeightExerciseRequest

// NewCreateWeightExerciseRequest instantiates a new CreateWeightExerciseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWeightExerciseRequest(name string, type_ string, sets int32, reps int32, weight float32) *CreateWeightExerciseRequest {
	this := CreateWeightExerciseRequest{}
	this.Name = name
	this.Type = type_
	this.Sets = sets
	this.Reps = reps
	this.Weight = weight
	return &this
}

// NewCreateWeightExerciseRequestWithDefaults instantiates a new CreateWeightExerciseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWeightExerciseRequestWithDefaults() *CreateWeightExerciseRequest {
	this := CreateWeightExerciseRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateWeightExerciseRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateWeightExerciseRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateWeightExerciseRequest) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *CreateWeightExerciseRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateWeightExerciseRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateWeightExerciseRequest) SetType(v string) {
	o.Type = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *CreateWeightExerciseRequest) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWeightExerciseRequest) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *CreateWeightExerciseRequest) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *CreateWeightExerciseRequest) SetNotes(v string) {
	o.Notes = &v
}

// GetSets returns the Sets field value
func (o *CreateWeightExerciseRequest) GetSets() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Sets
}

// GetSetsOk returns a tuple with the Sets field value
// and a boolean to check if the value has been set.
func (o *CreateWeightExerciseRequest) GetSetsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sets, true
}

// SetSets sets field value
func (o *CreateWeightExerciseRequest) SetSets(v int32) {
	o.Sets = v
}

// GetReps returns the Reps field value
func (o *CreateWeightExerciseRequest) GetReps() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Reps
}

// GetRepsOk returns a tuple with the Reps field value
// and a boolean to check if the value has been set.
func (o *CreateWeightExerciseRequest) GetRepsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reps, true
}

// SetReps sets field value
func (o *CreateWeightExerciseRequest) SetReps(v int32) {
	o.Reps = v
}

// GetWeight returns the Weight field value
func (o *CreateWeightExerciseRequest) GetWeight() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *CreateWeightExerciseRequest) GetWeightOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *CreateWeightExerciseRequest) SetWeight(v float32) {
	o.Weight = v
}

func (o CreateWeightExerciseRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateWeightExerciseRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	toSerialize["sets"] = o.Sets
	toSerialize["reps"] = o.Reps
	toSerialize["weight"] = o.Weight
	return toSerialize, nil
}

func (o *CreateWeightExerciseRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
		"sets",
		"reps",
		"weight",
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

	varCreateWeightExerciseRequest := _CreateWeightExerciseRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateWeightExerciseRequest)

	if err != nil {
		return err
	}

	*o = CreateWeightExerciseRequest(varCreateWeightExerciseRequest)

	return err
}

type NullableCreateWeightExerciseRequest struct {
	value *CreateWeightExerciseRequest
	isSet bool
}

func (v NullableCreateWeightExerciseRequest) Get() *CreateWeightExerciseRequest {
	return v.value
}

func (v *NullableCreateWeightExerciseRequest) Set(val *CreateWeightExerciseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWeightExerciseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWeightExerciseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWeightExerciseRequest(val *CreateWeightExerciseRequest) *NullableCreateWeightExerciseRequest {
	return &NullableCreateWeightExerciseRequest{value: val, isSet: true}
}

func (v NullableCreateWeightExerciseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWeightExerciseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


