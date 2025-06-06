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
	"time"
	"bytes"
	"fmt"
)

// checks if the Message type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Message{}

// Message Represents a message in a conversation.
type Message struct {
	// Unique identifier for the message.
	Id int64 `json:"id"`
	// Conversation ID this message belongs to.
	ConversationId int64 `json:"conversation_id"`
	// User ID who sent the message.
	UserId int64 `json:"user_id"`
	// Content of the message.
	Content string `json:"content"`
	MessageType MessageType `json:"message_type"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at"`
}

type _Message Message

// NewMessage instantiates a new Message object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessage(id int64, conversationId int64, userId int64, content string, messageType MessageType, createdAt time.Time) *Message {
	this := Message{}
	this.Id = id
	this.ConversationId = conversationId
	this.UserId = userId
	this.Content = content
	this.MessageType = messageType
	this.CreatedAt = createdAt
	return &this
}

// NewMessageWithDefaults instantiates a new Message object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageWithDefaults() *Message {
	this := Message{}
	return &this
}

// GetId returns the Id field value
func (o *Message) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Message) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Message) SetId(v int64) {
	o.Id = v
}

// GetConversationId returns the ConversationId field value
func (o *Message) GetConversationId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ConversationId
}

// GetConversationIdOk returns a tuple with the ConversationId field value
// and a boolean to check if the value has been set.
func (o *Message) GetConversationIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConversationId, true
}

// SetConversationId sets field value
func (o *Message) SetConversationId(v int64) {
	o.ConversationId = v
}

// GetUserId returns the UserId field value
func (o *Message) GetUserId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *Message) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *Message) SetUserId(v int64) {
	o.UserId = v
}

// GetContent returns the Content field value
func (o *Message) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *Message) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *Message) SetContent(v string) {
	o.Content = v
}

// GetMessageType returns the MessageType field value
func (o *Message) GetMessageType() MessageType {
	if o == nil {
		var ret MessageType
		return ret
	}

	return o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value
// and a boolean to check if the value has been set.
func (o *Message) GetMessageTypeOk() (*MessageType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageType, true
}

// SetMessageType sets field value
func (o *Message) SetMessageType(v MessageType) {
	o.MessageType = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Message) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Message) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Message) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o Message) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Message) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["conversation_id"] = o.ConversationId
	toSerialize["user_id"] = o.UserId
	toSerialize["content"] = o.Content
	toSerialize["message_type"] = o.MessageType
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

func (o *Message) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"conversation_id",
		"user_id",
		"content",
		"message_type",
		"created_at",
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

	varMessage := _Message{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessage)

	if err != nil {
		return err
	}

	*o = Message(varMessage)

	return err
}

type NullableMessage struct {
	value *Message
	isSet bool
}

func (v NullableMessage) Get() *Message {
	return v.value
}

func (v *NullableMessage) Set(val *Message) {
	v.value = val
	v.isSet = true
}

func (v NullableMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessage(val *Message) *NullableMessage {
	return &NullableMessage{value: val, isSet: true}
}

func (v NullableMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


