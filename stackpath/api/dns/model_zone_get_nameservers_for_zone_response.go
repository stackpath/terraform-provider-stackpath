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

// checks if the ZoneGetNameserversForZoneResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZoneGetNameserversForZoneResponse{}

// ZoneGetNameserversForZoneResponse A response from a request to retrieve information about a DNS zone's authoritative nameservers
type ZoneGetNameserversForZoneResponse struct {
	// Whether or not all required name servers are configured in the zone
	Configured *bool `json:"configured,omitempty"`
	// The zone's currently configured nameservers
	CurrentNameservers []string `json:"currentNameservers,omitempty"`
	// The nameservers required in the zone's configuration
	RequiredNameservers []string `json:"requiredNameservers,omitempty"`
}

// NewZoneGetNameserversForZoneResponse instantiates a new ZoneGetNameserversForZoneResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneGetNameserversForZoneResponse() *ZoneGetNameserversForZoneResponse {
	this := ZoneGetNameserversForZoneResponse{}
	return &this
}

// NewZoneGetNameserversForZoneResponseWithDefaults instantiates a new ZoneGetNameserversForZoneResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneGetNameserversForZoneResponseWithDefaults() *ZoneGetNameserversForZoneResponse {
	this := ZoneGetNameserversForZoneResponse{}
	return &this
}

// GetConfigured returns the Configured field value if set, zero value otherwise.
func (o *ZoneGetNameserversForZoneResponse) GetConfigured() bool {
	if o == nil || IsNil(o.Configured) {
		var ret bool
		return ret
	}
	return *o.Configured
}

// GetConfiguredOk returns a tuple with the Configured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneGetNameserversForZoneResponse) GetConfiguredOk() (*bool, bool) {
	if o == nil || IsNil(o.Configured) {
		return nil, false
	}
	return o.Configured, true
}

// HasConfigured returns a boolean if a field has been set.
func (o *ZoneGetNameserversForZoneResponse) HasConfigured() bool {
	if o != nil && !IsNil(o.Configured) {
		return true
	}

	return false
}

// SetConfigured gets a reference to the given bool and assigns it to the Configured field.
func (o *ZoneGetNameserversForZoneResponse) SetConfigured(v bool) {
	o.Configured = &v
}

// GetCurrentNameservers returns the CurrentNameservers field value if set, zero value otherwise.
func (o *ZoneGetNameserversForZoneResponse) GetCurrentNameservers() []string {
	if o == nil || IsNil(o.CurrentNameservers) {
		var ret []string
		return ret
	}
	return o.CurrentNameservers
}

// GetCurrentNameserversOk returns a tuple with the CurrentNameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneGetNameserversForZoneResponse) GetCurrentNameserversOk() ([]string, bool) {
	if o == nil || IsNil(o.CurrentNameservers) {
		return nil, false
	}
	return o.CurrentNameservers, true
}

// HasCurrentNameservers returns a boolean if a field has been set.
func (o *ZoneGetNameserversForZoneResponse) HasCurrentNameservers() bool {
	if o != nil && !IsNil(o.CurrentNameservers) {
		return true
	}

	return false
}

// SetCurrentNameservers gets a reference to the given []string and assigns it to the CurrentNameservers field.
func (o *ZoneGetNameserversForZoneResponse) SetCurrentNameservers(v []string) {
	o.CurrentNameservers = v
}

// GetRequiredNameservers returns the RequiredNameservers field value if set, zero value otherwise.
func (o *ZoneGetNameserversForZoneResponse) GetRequiredNameservers() []string {
	if o == nil || IsNil(o.RequiredNameservers) {
		var ret []string
		return ret
	}
	return o.RequiredNameservers
}

// GetRequiredNameserversOk returns a tuple with the RequiredNameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneGetNameserversForZoneResponse) GetRequiredNameserversOk() ([]string, bool) {
	if o == nil || IsNil(o.RequiredNameservers) {
		return nil, false
	}
	return o.RequiredNameservers, true
}

// HasRequiredNameservers returns a boolean if a field has been set.
func (o *ZoneGetNameserversForZoneResponse) HasRequiredNameservers() bool {
	if o != nil && !IsNil(o.RequiredNameservers) {
		return true
	}

	return false
}

// SetRequiredNameservers gets a reference to the given []string and assigns it to the RequiredNameservers field.
func (o *ZoneGetNameserversForZoneResponse) SetRequiredNameservers(v []string) {
	o.RequiredNameservers = v
}

func (o ZoneGetNameserversForZoneResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZoneGetNameserversForZoneResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Configured) {
		toSerialize["configured"] = o.Configured
	}
	if !IsNil(o.CurrentNameservers) {
		toSerialize["currentNameservers"] = o.CurrentNameservers
	}
	if !IsNil(o.RequiredNameservers) {
		toSerialize["requiredNameservers"] = o.RequiredNameservers
	}
	return toSerialize, nil
}

type NullableZoneGetNameserversForZoneResponse struct {
	value *ZoneGetNameserversForZoneResponse
	isSet bool
}

func (v NullableZoneGetNameserversForZoneResponse) Get() *ZoneGetNameserversForZoneResponse {
	return v.value
}

func (v *NullableZoneGetNameserversForZoneResponse) Set(val *ZoneGetNameserversForZoneResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneGetNameserversForZoneResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneGetNameserversForZoneResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneGetNameserversForZoneResponse(val *ZoneGetNameserversForZoneResponse) *NullableZoneGetNameserversForZoneResponse {
	return &NullableZoneGetNameserversForZoneResponse{value: val, isSet: true}
}

func (v NullableZoneGetNameserversForZoneResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneGetNameserversForZoneResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

