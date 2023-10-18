/*
Metal API

# Introduction Equinix Metal provides a RESTful HTTP API which can be reached at <https://api.equinix.com/metal/v1>. This document describes the API and how to use it.  The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account. Every feature of the Equinix Metal web interface is accessible through the API.  The API docs are generated from the Equinix Metal OpenAPI specification and are officially hosted at <https://metal.equinix.com/developers/api>.  # Common Parameters  The Equinix Metal API uses a few methods to minimize network traffic and improve throughput. These parameters are not used in all API calls, but are used often enough to warrant their own section. Look for these parameters in the documentation for the API calls that support them.  ## Pagination  Pagination is used to limit the number of results returned in a single request. The API will return a maximum of 100 results per page. To retrieve additional results, you can use the `page` and `per_page` query parameters.  The `page` parameter is used to specify the page number. The first page is `1`. The `per_page` parameter is used to specify the number of results per page. The maximum number of results differs by resource type.  ## Sorting  Where offered, the API allows you to sort results by a specific field. To sort results use the `sort_by` query parameter with the root level field name as the value. The `sort_direction` parameter is used to specify the sort direction, either either `asc` (ascending) or `desc` (descending).  ## Filtering  Filtering is used to limit the results returned in a single request. The API supports filtering by certain fields in the response. To filter results, you can use the field as a query parameter.  For example, to filter the IP list to only return public IPv4 addresses, you can filter by the `type` field, as in the following request:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/projects/id/ips?type=public_ipv4 ```  Only IP addresses with the `type` field set to `public_ipv4` will be returned.  ## Searching  Searching is used to find matching resources using multiple field comparissons. The API supports searching in resources that define this behavior. Currently the search parameter is only available on devices, ssh_keys, api_keys and memberships endpoints.  To search resources you can use the `search` query parameter.  ## Include and Exclude  For resources that contain references to other resources, sucha as a Device that refers to the Project it resides in, the Equinix Metal API will returns `href` values (API links) to the associated resource.  ```json {   ...   \"project\": {     \"href\": \"/metal/v1/projects/f3f131c8-f302-49ef-8c44-9405022dc6dd\"   } } ```  If you're going need the project details, you can avoid a second API request.  Specify the contained `href` resources and collections that you'd like to have included in the response using the `include` query parameter.  For example:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=projects ```  The `include` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests where `href` resources are presented.  To have multiple resources include, use a comma-separated list (e.g. `?include=emails,projects,memberships`).  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=emails,projects,memberships ```  You may also include nested associations up to three levels deep using dot notation (`?include=memberships.projects`):  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=memberships.projects ```  To exclude resources, and optimize response delivery, use the `exclude` query parameter. The `exclude` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests for fields with nested object responses. When excluded, these fields will be replaced with an object that contains only an `href` field.

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// checks if the SupportRequestInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SupportRequestInput{}

// SupportRequestInput struct for SupportRequestInput
type SupportRequestInput struct {
	DeviceId             *string                      `json:"device_id,omitempty"`
	Message              string                       `json:"message"`
	Priority             *SupportRequestInputPriority `json:"priority,omitempty"`
	ProjectId            *string                      `json:"project_id,omitempty"`
	Subject              string                       `json:"subject"`
	AdditionalProperties map[string]interface{}
}

type _SupportRequestInput SupportRequestInput

// NewSupportRequestInput instantiates a new SupportRequestInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportRequestInput(message string, subject string) *SupportRequestInput {
	this := SupportRequestInput{}
	this.Message = message
	this.Subject = subject
	return &this
}

// NewSupportRequestInputWithDefaults instantiates a new SupportRequestInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportRequestInputWithDefaults() *SupportRequestInput {
	this := SupportRequestInput{}
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *SupportRequestInput) GetDeviceId() string {
	if o == nil || IsNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportRequestInput) GetDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *SupportRequestInput) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *SupportRequestInput) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetMessage returns the Message field value
func (o *SupportRequestInput) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SupportRequestInput) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SupportRequestInput) SetMessage(v string) {
	o.Message = v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *SupportRequestInput) GetPriority() SupportRequestInputPriority {
	if o == nil || IsNil(o.Priority) {
		var ret SupportRequestInputPriority
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportRequestInput) GetPriorityOk() (*SupportRequestInputPriority, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *SupportRequestInput) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given SupportRequestInputPriority and assigns it to the Priority field.
func (o *SupportRequestInput) SetPriority(v SupportRequestInputPriority) {
	o.Priority = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *SupportRequestInput) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportRequestInput) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *SupportRequestInput) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *SupportRequestInput) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetSubject returns the Subject field value
func (o *SupportRequestInput) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *SupportRequestInput) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *SupportRequestInput) SetSubject(v string) {
	o.Subject = v
}

func (o SupportRequestInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupportRequestInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceId) {
		toSerialize["device_id"] = o.DeviceId
	}
	toSerialize["message"] = o.Message
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.ProjectId) {
		toSerialize["project_id"] = o.ProjectId
	}
	toSerialize["subject"] = o.Subject

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SupportRequestInput) UnmarshalJSON(bytes []byte) (err error) {
	requiredProperties := []string{
		"message",
		"subject",
	}

	allProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &allProperties); err == nil {
		for _, requiredProperty := range requiredProperties {
			if _, exists := allProperties[requiredProperty]; !exists {
				return MissingRequiredFieldError(requiredProperty)
			}
		}
	}

	varSupportRequestInput := _SupportRequestInput{}

	err = json.Unmarshal(bytes, &varSupportRequestInput)

	if err != nil {
		return err
	}

	*o = SupportRequestInput(varSupportRequestInput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "device_id")
		delete(additionalProperties, "message")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "project_id")
		delete(additionalProperties, "subject")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSupportRequestInput struct {
	value *SupportRequestInput
	isSet bool
}

func (v NullableSupportRequestInput) Get() *SupportRequestInput {
	return v.value
}

func (v *NullableSupportRequestInput) Set(val *SupportRequestInput) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportRequestInput) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportRequestInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportRequestInput(val *SupportRequestInput) *NullableSupportRequestInput {
	return &NullableSupportRequestInput{value: val, isSet: true}
}

func (v NullableSupportRequestInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportRequestInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
