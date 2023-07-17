/*
DNS

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
)

// checks if the ZoneUpdateZoneRecordMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZoneUpdateZoneRecordMessage{}

// ZoneUpdateZoneRecordMessage struct for ZoneUpdateZoneRecordMessage
type ZoneUpdateZoneRecordMessage struct {
	// A zone record's name
	Name *string `json:"name,omitempty"`
	Type *ZoneRecordType `json:"type,omitempty"`
	// A zone record's time to live  A record's TTL is the number of seconds that the record should be cached by DNS resolvers. Use lower TTL values if you expect zone records to change often. Use higher TTL values for records that won't change to prevent extra DNS lookups by clients.
	Ttl *int32 `json:"ttl,omitempty"`
	// A zone record's value  MX record data follows the format `<priority> <value>`, without `<>`s.
	Data *string `json:"data,omitempty"`
	// A zone record's priority  A resource record is replicated in StackPath's DNS infrastructure the number of times of the record's weight, giving it a more likely response to queries if a zone has records with the same name and type.
	Weight *int32 `json:"weight,omitempty"`
	// A key/value pair of user-defined labels for a DNS zone record  Zone record labels are not processed by StackPath and are solely used for users to organize their DNS zones.
	Labels *map[string]string `json:"labels,omitempty"`
}

// NewZoneUpdateZoneRecordMessage instantiates a new ZoneUpdateZoneRecordMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneUpdateZoneRecordMessage() *ZoneUpdateZoneRecordMessage {
	this := ZoneUpdateZoneRecordMessage{}
	var type_ ZoneRecordType = EMPTY
	this.Type = &type_
	return &this
}

// NewZoneUpdateZoneRecordMessageWithDefaults instantiates a new ZoneUpdateZoneRecordMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneUpdateZoneRecordMessageWithDefaults() *ZoneUpdateZoneRecordMessage {
	this := ZoneUpdateZoneRecordMessage{}
	var type_ ZoneRecordType = EMPTY
	this.Type = &type_
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ZoneUpdateZoneRecordMessage) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneUpdateZoneRecordMessage) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ZoneUpdateZoneRecordMessage) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ZoneUpdateZoneRecordMessage) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ZoneUpdateZoneRecordMessage) GetType() ZoneRecordType {
	if o == nil || IsNil(o.Type) {
		var ret ZoneRecordType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneUpdateZoneRecordMessage) GetTypeOk() (*ZoneRecordType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ZoneUpdateZoneRecordMessage) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ZoneRecordType and assigns it to the Type field.
func (o *ZoneUpdateZoneRecordMessage) SetType(v ZoneRecordType) {
	o.Type = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *ZoneUpdateZoneRecordMessage) GetTtl() int32 {
	if o == nil || IsNil(o.Ttl) {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneUpdateZoneRecordMessage) GetTtlOk() (*int32, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *ZoneUpdateZoneRecordMessage) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
func (o *ZoneUpdateZoneRecordMessage) SetTtl(v int32) {
	o.Ttl = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ZoneUpdateZoneRecordMessage) GetData() string {
	if o == nil || IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneUpdateZoneRecordMessage) GetDataOk() (*string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ZoneUpdateZoneRecordMessage) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *ZoneUpdateZoneRecordMessage) SetData(v string) {
	o.Data = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *ZoneUpdateZoneRecordMessage) GetWeight() int32 {
	if o == nil || IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneUpdateZoneRecordMessage) GetWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *ZoneUpdateZoneRecordMessage) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *ZoneUpdateZoneRecordMessage) SetWeight(v int32) {
	o.Weight = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *ZoneUpdateZoneRecordMessage) GetLabels() map[string]string {
	if o == nil || IsNil(o.Labels) {
		var ret map[string]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneUpdateZoneRecordMessage) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *ZoneUpdateZoneRecordMessage) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *ZoneUpdateZoneRecordMessage) SetLabels(v map[string]string) {
	o.Labels = &v
}

func (o ZoneUpdateZoneRecordMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZoneUpdateZoneRecordMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	return toSerialize, nil
}

type NullableZoneUpdateZoneRecordMessage struct {
	value *ZoneUpdateZoneRecordMessage
	isSet bool
}

func (v NullableZoneUpdateZoneRecordMessage) Get() *ZoneUpdateZoneRecordMessage {
	return v.value
}

func (v *NullableZoneUpdateZoneRecordMessage) Set(val *ZoneUpdateZoneRecordMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneUpdateZoneRecordMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneUpdateZoneRecordMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneUpdateZoneRecordMessage(val *ZoneUpdateZoneRecordMessage) *NullableZoneUpdateZoneRecordMessage {
	return &NullableZoneUpdateZoneRecordMessage{value: val, isSet: true}
}

func (v NullableZoneUpdateZoneRecordMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneUpdateZoneRecordMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

