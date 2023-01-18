// Additional necessary methods that codegen did not include in utils.go

package v1

import "encoding/json"

type AnyOfarraystring struct {
	value []string
}

func (v AnyOfarraystring) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

type NullableAnyOfarraystring struct {
	value *AnyOfarraystring
	isSet bool
}

func (v NullableAnyOfarraystring) Get() *AnyOfarraystring {
	return v.value
}

func (v *NullableAnyOfarraystring) Set(val *AnyOfarraystring) {
	v.value = val
	v.isSet = true
}

func (v NullableAnyOfarraystring) IsSet() bool {
	return v.isSet
}

func (v *NullableAnyOfarraystring) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnyOfarraystring(val *AnyOfarraystring) *NullableAnyOfarraystring {
	return &NullableAnyOfarraystring{value: val, isSet: true}
}

func (v NullableAnyOfarraystring) MarshalJSON() ([]byte, error) {
	return v.value.MarshalJSON()
}

func (v *NullableAnyOfarraystring) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
