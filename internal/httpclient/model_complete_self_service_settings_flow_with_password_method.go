/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests.
 *
 * API version: 1.0.0
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kratos

import (
	"encoding/json"
)

// CompleteSelfServiceSettingsFlowWithPasswordMethod struct for CompleteSelfServiceSettingsFlowWithPasswordMethod
type CompleteSelfServiceSettingsFlowWithPasswordMethod struct {
	// CSRFToken is the anti-CSRF token  type: string
	CsrfToken *string `json:"csrf_token,omitempty"`
	// Method  Should be set to password when trying to update a password.  type: string
	Method *string `json:"method,omitempty"`
	// Password is the updated password  type: string
	Password string `json:"password"`
}

// NewCompleteSelfServiceSettingsFlowWithPasswordMethod instantiates a new CompleteSelfServiceSettingsFlowWithPasswordMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompleteSelfServiceSettingsFlowWithPasswordMethod(password string) *CompleteSelfServiceSettingsFlowWithPasswordMethod {
	this := CompleteSelfServiceSettingsFlowWithPasswordMethod{}
	this.Password = password
	return &this
}

// NewCompleteSelfServiceSettingsFlowWithPasswordMethodWithDefaults instantiates a new CompleteSelfServiceSettingsFlowWithPasswordMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompleteSelfServiceSettingsFlowWithPasswordMethodWithDefaults() *CompleteSelfServiceSettingsFlowWithPasswordMethod {
	this := CompleteSelfServiceSettingsFlowWithPasswordMethod{}
	return &this
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *CompleteSelfServiceSettingsFlowWithPasswordMethod) GetCsrfToken() string {
	if o == nil || o.CsrfToken == nil {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompleteSelfServiceSettingsFlowWithPasswordMethod) GetCsrfTokenOk() (*string, bool) {
	if o == nil || o.CsrfToken == nil {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *CompleteSelfServiceSettingsFlowWithPasswordMethod) HasCsrfToken() bool {
	if o != nil && o.CsrfToken != nil {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *CompleteSelfServiceSettingsFlowWithPasswordMethod) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *CompleteSelfServiceSettingsFlowWithPasswordMethod) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompleteSelfServiceSettingsFlowWithPasswordMethod) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *CompleteSelfServiceSettingsFlowWithPasswordMethod) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *CompleteSelfServiceSettingsFlowWithPasswordMethod) SetMethod(v string) {
	o.Method = &v
}

// GetPassword returns the Password field value
func (o *CompleteSelfServiceSettingsFlowWithPasswordMethod) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *CompleteSelfServiceSettingsFlowWithPasswordMethod) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *CompleteSelfServiceSettingsFlowWithPasswordMethod) SetPassword(v string) {
	o.Password = v
}

func (o CompleteSelfServiceSettingsFlowWithPasswordMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CsrfToken != nil {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if true {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableCompleteSelfServiceSettingsFlowWithPasswordMethod struct {
	value *CompleteSelfServiceSettingsFlowWithPasswordMethod
	isSet bool
}

func (v NullableCompleteSelfServiceSettingsFlowWithPasswordMethod) Get() *CompleteSelfServiceSettingsFlowWithPasswordMethod {
	return v.value
}

func (v *NullableCompleteSelfServiceSettingsFlowWithPasswordMethod) Set(val *CompleteSelfServiceSettingsFlowWithPasswordMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableCompleteSelfServiceSettingsFlowWithPasswordMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableCompleteSelfServiceSettingsFlowWithPasswordMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompleteSelfServiceSettingsFlowWithPasswordMethod(val *CompleteSelfServiceSettingsFlowWithPasswordMethod) *NullableCompleteSelfServiceSettingsFlowWithPasswordMethod {
	return &NullableCompleteSelfServiceSettingsFlowWithPasswordMethod{value: val, isSet: true}
}

func (v NullableCompleteSelfServiceSettingsFlowWithPasswordMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompleteSelfServiceSettingsFlowWithPasswordMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
