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

// checks if the CreateMessageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateMessageRequest{}

// CreateMessageRequest Payload to create a new message.
type CreateMessageRequest struct {
	// Content of the message.
	Content string `json:"content"`
	MessageType MessageType `json:"message_type"`
}

type _CreateMessageRequest CreateMessageRequest

// NewCreateMessageRequest instantiates a new CreateMessageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMessageRequest(content string, messageType MessageType) *CreateMessageRequest {
	this := CreateMessageRequest{}
	this.Content = content
	this.MessageType = messageType
	return &this
}

// NewCreateMessageRequestWithDefaults instantiates a new CreateMessageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMessageRequestWithDefaults() *CreateMessageRequest {
	this := CreateMessageRequest{}
	return &this
}

// GetContent returns the Content field value
func (o *CreateMessageRequest) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *CreateMessageRequest) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *CreateMessageRequest) SetContent(v string) {
	o.Content = v
}

// GetMessageType returns the MessageType field value
func (o *CreateMessageRequest) GetMessageType() MessageType {
	if o == nil {
		var ret MessageType
		return ret
	}

	return o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value
// and a boolean to check if the value has been set.
func (o *CreateMessageRequest) GetMessageTypeOk() (*MessageType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageType, true
}

// SetMessageType sets field value
func (o *CreateMessageRequest) SetMessageType(v MessageType) {
	o.MessageType = v
}

func (o CreateMessageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMessageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["content"] = o.Content
	toSerialize["message_type"] = o.MessageType
	return toSerialize, nil
}

func (o *CreateMessageRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"content",
		"message_type",
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

	varCreateMessageRequest := _CreateMessageRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateMessageRequest)

	if err != nil {
		return err
	}

	*o = CreateMessageRequest(varCreateMessageRequest)

	return err
}

type NullableCreateMessageRequest struct {
	value *CreateMessageRequest
	isSet bool
}

func (v NullableCreateMessageRequest) Get() *CreateMessageRequest {
	return v.value
}

func (v *NullableCreateMessageRequest) Set(val *CreateMessageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMessageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMessageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMessageRequest(val *CreateMessageRequest) *NullableCreateMessageRequest {
	return &NullableCreateMessageRequest{value: val, isSet: true}
}

func (v NullableCreateMessageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMessageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


