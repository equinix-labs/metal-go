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
	"fmt"
)

// FindIPAvailabilitiesCidrParameter the model 'FindIPAvailabilitiesCidrParameter'
type FindIPAvailabilitiesCidrParameter string

// List of findIPAvailabilities_cidr_parameter
const (
	FINDIPAVAILABILITIESCIDRPARAMETER__20  FindIPAvailabilitiesCidrParameter = "20"
	FINDIPAVAILABILITIESCIDRPARAMETER__21  FindIPAvailabilitiesCidrParameter = "21"
	FINDIPAVAILABILITIESCIDRPARAMETER__22  FindIPAvailabilitiesCidrParameter = "22"
	FINDIPAVAILABILITIESCIDRPARAMETER__23  FindIPAvailabilitiesCidrParameter = "23"
	FINDIPAVAILABILITIESCIDRPARAMETER__24  FindIPAvailabilitiesCidrParameter = "24"
	FINDIPAVAILABILITIESCIDRPARAMETER__25  FindIPAvailabilitiesCidrParameter = "25"
	FINDIPAVAILABILITIESCIDRPARAMETER__26  FindIPAvailabilitiesCidrParameter = "26"
	FINDIPAVAILABILITIESCIDRPARAMETER__27  FindIPAvailabilitiesCidrParameter = "27"
	FINDIPAVAILABILITIESCIDRPARAMETER__28  FindIPAvailabilitiesCidrParameter = "28"
	FINDIPAVAILABILITIESCIDRPARAMETER__29  FindIPAvailabilitiesCidrParameter = "29"
	FINDIPAVAILABILITIESCIDRPARAMETER__30  FindIPAvailabilitiesCidrParameter = "30"
	FINDIPAVAILABILITIESCIDRPARAMETER__31  FindIPAvailabilitiesCidrParameter = "31"
	FINDIPAVAILABILITIESCIDRPARAMETER__32  FindIPAvailabilitiesCidrParameter = "32"
	FINDIPAVAILABILITIESCIDRPARAMETER__33  FindIPAvailabilitiesCidrParameter = "33"
	FINDIPAVAILABILITIESCIDRPARAMETER__34  FindIPAvailabilitiesCidrParameter = "34"
	FINDIPAVAILABILITIESCIDRPARAMETER__35  FindIPAvailabilitiesCidrParameter = "35"
	FINDIPAVAILABILITIESCIDRPARAMETER__36  FindIPAvailabilitiesCidrParameter = "36"
	FINDIPAVAILABILITIESCIDRPARAMETER__37  FindIPAvailabilitiesCidrParameter = "37"
	FINDIPAVAILABILITIESCIDRPARAMETER__38  FindIPAvailabilitiesCidrParameter = "38"
	FINDIPAVAILABILITIESCIDRPARAMETER__39  FindIPAvailabilitiesCidrParameter = "39"
	FINDIPAVAILABILITIESCIDRPARAMETER__40  FindIPAvailabilitiesCidrParameter = "40"
	FINDIPAVAILABILITIESCIDRPARAMETER__41  FindIPAvailabilitiesCidrParameter = "41"
	FINDIPAVAILABILITIESCIDRPARAMETER__42  FindIPAvailabilitiesCidrParameter = "42"
	FINDIPAVAILABILITIESCIDRPARAMETER__43  FindIPAvailabilitiesCidrParameter = "43"
	FINDIPAVAILABILITIESCIDRPARAMETER__44  FindIPAvailabilitiesCidrParameter = "44"
	FINDIPAVAILABILITIESCIDRPARAMETER__45  FindIPAvailabilitiesCidrParameter = "45"
	FINDIPAVAILABILITIESCIDRPARAMETER__46  FindIPAvailabilitiesCidrParameter = "46"
	FINDIPAVAILABILITIESCIDRPARAMETER__47  FindIPAvailabilitiesCidrParameter = "47"
	FINDIPAVAILABILITIESCIDRPARAMETER__48  FindIPAvailabilitiesCidrParameter = "48"
	FINDIPAVAILABILITIESCIDRPARAMETER__49  FindIPAvailabilitiesCidrParameter = "49"
	FINDIPAVAILABILITIESCIDRPARAMETER__50  FindIPAvailabilitiesCidrParameter = "50"
	FINDIPAVAILABILITIESCIDRPARAMETER__51  FindIPAvailabilitiesCidrParameter = "51"
	FINDIPAVAILABILITIESCIDRPARAMETER__52  FindIPAvailabilitiesCidrParameter = "52"
	FINDIPAVAILABILITIESCIDRPARAMETER__53  FindIPAvailabilitiesCidrParameter = "53"
	FINDIPAVAILABILITIESCIDRPARAMETER__54  FindIPAvailabilitiesCidrParameter = "54"
	FINDIPAVAILABILITIESCIDRPARAMETER__55  FindIPAvailabilitiesCidrParameter = "55"
	FINDIPAVAILABILITIESCIDRPARAMETER__56  FindIPAvailabilitiesCidrParameter = "56"
	FINDIPAVAILABILITIESCIDRPARAMETER__57  FindIPAvailabilitiesCidrParameter = "57"
	FINDIPAVAILABILITIESCIDRPARAMETER__58  FindIPAvailabilitiesCidrParameter = "58"
	FINDIPAVAILABILITIESCIDRPARAMETER__59  FindIPAvailabilitiesCidrParameter = "59"
	FINDIPAVAILABILITIESCIDRPARAMETER__60  FindIPAvailabilitiesCidrParameter = "60"
	FINDIPAVAILABILITIESCIDRPARAMETER__61  FindIPAvailabilitiesCidrParameter = "61"
	FINDIPAVAILABILITIESCIDRPARAMETER__62  FindIPAvailabilitiesCidrParameter = "62"
	FINDIPAVAILABILITIESCIDRPARAMETER__63  FindIPAvailabilitiesCidrParameter = "63"
	FINDIPAVAILABILITIESCIDRPARAMETER__64  FindIPAvailabilitiesCidrParameter = "64"
	FINDIPAVAILABILITIESCIDRPARAMETER__65  FindIPAvailabilitiesCidrParameter = "65"
	FINDIPAVAILABILITIESCIDRPARAMETER__66  FindIPAvailabilitiesCidrParameter = "66"
	FINDIPAVAILABILITIESCIDRPARAMETER__67  FindIPAvailabilitiesCidrParameter = "67"
	FINDIPAVAILABILITIESCIDRPARAMETER__68  FindIPAvailabilitiesCidrParameter = "68"
	FINDIPAVAILABILITIESCIDRPARAMETER__69  FindIPAvailabilitiesCidrParameter = "69"
	FINDIPAVAILABILITIESCIDRPARAMETER__70  FindIPAvailabilitiesCidrParameter = "70"
	FINDIPAVAILABILITIESCIDRPARAMETER__71  FindIPAvailabilitiesCidrParameter = "71"
	FINDIPAVAILABILITIESCIDRPARAMETER__72  FindIPAvailabilitiesCidrParameter = "72"
	FINDIPAVAILABILITIESCIDRPARAMETER__73  FindIPAvailabilitiesCidrParameter = "73"
	FINDIPAVAILABILITIESCIDRPARAMETER__74  FindIPAvailabilitiesCidrParameter = "74"
	FINDIPAVAILABILITIESCIDRPARAMETER__75  FindIPAvailabilitiesCidrParameter = "75"
	FINDIPAVAILABILITIESCIDRPARAMETER__76  FindIPAvailabilitiesCidrParameter = "76"
	FINDIPAVAILABILITIESCIDRPARAMETER__77  FindIPAvailabilitiesCidrParameter = "77"
	FINDIPAVAILABILITIESCIDRPARAMETER__78  FindIPAvailabilitiesCidrParameter = "78"
	FINDIPAVAILABILITIESCIDRPARAMETER__79  FindIPAvailabilitiesCidrParameter = "79"
	FINDIPAVAILABILITIESCIDRPARAMETER__80  FindIPAvailabilitiesCidrParameter = "80"
	FINDIPAVAILABILITIESCIDRPARAMETER__81  FindIPAvailabilitiesCidrParameter = "81"
	FINDIPAVAILABILITIESCIDRPARAMETER__82  FindIPAvailabilitiesCidrParameter = "82"
	FINDIPAVAILABILITIESCIDRPARAMETER__83  FindIPAvailabilitiesCidrParameter = "83"
	FINDIPAVAILABILITIESCIDRPARAMETER__84  FindIPAvailabilitiesCidrParameter = "84"
	FINDIPAVAILABILITIESCIDRPARAMETER__85  FindIPAvailabilitiesCidrParameter = "85"
	FINDIPAVAILABILITIESCIDRPARAMETER__86  FindIPAvailabilitiesCidrParameter = "86"
	FINDIPAVAILABILITIESCIDRPARAMETER__87  FindIPAvailabilitiesCidrParameter = "87"
	FINDIPAVAILABILITIESCIDRPARAMETER__88  FindIPAvailabilitiesCidrParameter = "88"
	FINDIPAVAILABILITIESCIDRPARAMETER__89  FindIPAvailabilitiesCidrParameter = "89"
	FINDIPAVAILABILITIESCIDRPARAMETER__90  FindIPAvailabilitiesCidrParameter = "90"
	FINDIPAVAILABILITIESCIDRPARAMETER__91  FindIPAvailabilitiesCidrParameter = "91"
	FINDIPAVAILABILITIESCIDRPARAMETER__92  FindIPAvailabilitiesCidrParameter = "92"
	FINDIPAVAILABILITIESCIDRPARAMETER__93  FindIPAvailabilitiesCidrParameter = "93"
	FINDIPAVAILABILITIESCIDRPARAMETER__94  FindIPAvailabilitiesCidrParameter = "94"
	FINDIPAVAILABILITIESCIDRPARAMETER__95  FindIPAvailabilitiesCidrParameter = "95"
	FINDIPAVAILABILITIESCIDRPARAMETER__96  FindIPAvailabilitiesCidrParameter = "96"
	FINDIPAVAILABILITIESCIDRPARAMETER__97  FindIPAvailabilitiesCidrParameter = "97"
	FINDIPAVAILABILITIESCIDRPARAMETER__98  FindIPAvailabilitiesCidrParameter = "98"
	FINDIPAVAILABILITIESCIDRPARAMETER__99  FindIPAvailabilitiesCidrParameter = "99"
	FINDIPAVAILABILITIESCIDRPARAMETER__100 FindIPAvailabilitiesCidrParameter = "100"
	FINDIPAVAILABILITIESCIDRPARAMETER__101 FindIPAvailabilitiesCidrParameter = "101"
	FINDIPAVAILABILITIESCIDRPARAMETER__102 FindIPAvailabilitiesCidrParameter = "102"
	FINDIPAVAILABILITIESCIDRPARAMETER__103 FindIPAvailabilitiesCidrParameter = "103"
	FINDIPAVAILABILITIESCIDRPARAMETER__104 FindIPAvailabilitiesCidrParameter = "104"
	FINDIPAVAILABILITIESCIDRPARAMETER__105 FindIPAvailabilitiesCidrParameter = "105"
	FINDIPAVAILABILITIESCIDRPARAMETER__106 FindIPAvailabilitiesCidrParameter = "106"
	FINDIPAVAILABILITIESCIDRPARAMETER__107 FindIPAvailabilitiesCidrParameter = "107"
	FINDIPAVAILABILITIESCIDRPARAMETER__108 FindIPAvailabilitiesCidrParameter = "108"
	FINDIPAVAILABILITIESCIDRPARAMETER__109 FindIPAvailabilitiesCidrParameter = "109"
	FINDIPAVAILABILITIESCIDRPARAMETER__110 FindIPAvailabilitiesCidrParameter = "110"
	FINDIPAVAILABILITIESCIDRPARAMETER__111 FindIPAvailabilitiesCidrParameter = "111"
	FINDIPAVAILABILITIESCIDRPARAMETER__112 FindIPAvailabilitiesCidrParameter = "112"
	FINDIPAVAILABILITIESCIDRPARAMETER__113 FindIPAvailabilitiesCidrParameter = "113"
	FINDIPAVAILABILITIESCIDRPARAMETER__114 FindIPAvailabilitiesCidrParameter = "114"
	FINDIPAVAILABILITIESCIDRPARAMETER__115 FindIPAvailabilitiesCidrParameter = "115"
	FINDIPAVAILABILITIESCIDRPARAMETER__116 FindIPAvailabilitiesCidrParameter = "116"
	FINDIPAVAILABILITIESCIDRPARAMETER__117 FindIPAvailabilitiesCidrParameter = "117"
	FINDIPAVAILABILITIESCIDRPARAMETER__118 FindIPAvailabilitiesCidrParameter = "118"
	FINDIPAVAILABILITIESCIDRPARAMETER__119 FindIPAvailabilitiesCidrParameter = "119"
	FINDIPAVAILABILITIESCIDRPARAMETER__120 FindIPAvailabilitiesCidrParameter = "120"
	FINDIPAVAILABILITIESCIDRPARAMETER__121 FindIPAvailabilitiesCidrParameter = "121"
	FINDIPAVAILABILITIESCIDRPARAMETER__122 FindIPAvailabilitiesCidrParameter = "122"
	FINDIPAVAILABILITIESCIDRPARAMETER__123 FindIPAvailabilitiesCidrParameter = "123"
	FINDIPAVAILABILITIESCIDRPARAMETER__124 FindIPAvailabilitiesCidrParameter = "124"
	FINDIPAVAILABILITIESCIDRPARAMETER__125 FindIPAvailabilitiesCidrParameter = "125"
	FINDIPAVAILABILITIESCIDRPARAMETER__126 FindIPAvailabilitiesCidrParameter = "126"
	FINDIPAVAILABILITIESCIDRPARAMETER__127 FindIPAvailabilitiesCidrParameter = "127"
	FINDIPAVAILABILITIESCIDRPARAMETER__128 FindIPAvailabilitiesCidrParameter = "128"
)

// All allowed values of FindIPAvailabilitiesCidrParameter enum
var AllowedFindIPAvailabilitiesCidrParameterEnumValues = []FindIPAvailabilitiesCidrParameter{
	"20",
	"21",
	"22",
	"23",
	"24",
	"25",
	"26",
	"27",
	"28",
	"29",
	"30",
	"31",
	"32",
	"33",
	"34",
	"35",
	"36",
	"37",
	"38",
	"39",
	"40",
	"41",
	"42",
	"43",
	"44",
	"45",
	"46",
	"47",
	"48",
	"49",
	"50",
	"51",
	"52",
	"53",
	"54",
	"55",
	"56",
	"57",
	"58",
	"59",
	"60",
	"61",
	"62",
	"63",
	"64",
	"65",
	"66",
	"67",
	"68",
	"69",
	"70",
	"71",
	"72",
	"73",
	"74",
	"75",
	"76",
	"77",
	"78",
	"79",
	"80",
	"81",
	"82",
	"83",
	"84",
	"85",
	"86",
	"87",
	"88",
	"89",
	"90",
	"91",
	"92",
	"93",
	"94",
	"95",
	"96",
	"97",
	"98",
	"99",
	"100",
	"101",
	"102",
	"103",
	"104",
	"105",
	"106",
	"107",
	"108",
	"109",
	"110",
	"111",
	"112",
	"113",
	"114",
	"115",
	"116",
	"117",
	"118",
	"119",
	"120",
	"121",
	"122",
	"123",
	"124",
	"125",
	"126",
	"127",
	"128",
}

func (v *FindIPAvailabilitiesCidrParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FindIPAvailabilitiesCidrParameter(value)
	for _, existing := range AllowedFindIPAvailabilitiesCidrParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FindIPAvailabilitiesCidrParameter", value)
}

// NewFindIPAvailabilitiesCidrParameterFromValue returns a pointer to a valid FindIPAvailabilitiesCidrParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFindIPAvailabilitiesCidrParameterFromValue(v string) (*FindIPAvailabilitiesCidrParameter, error) {
	ev := FindIPAvailabilitiesCidrParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FindIPAvailabilitiesCidrParameter: valid values are %v", v, AllowedFindIPAvailabilitiesCidrParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FindIPAvailabilitiesCidrParameter) IsValid() bool {
	for _, existing := range AllowedFindIPAvailabilitiesCidrParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to findIPAvailabilities_cidr_parameter value
func (v FindIPAvailabilitiesCidrParameter) Ptr() *FindIPAvailabilitiesCidrParameter {
	return &v
}

type NullableFindIPAvailabilitiesCidrParameter struct {
	value *FindIPAvailabilitiesCidrParameter
	isSet bool
}

func (v NullableFindIPAvailabilitiesCidrParameter) Get() *FindIPAvailabilitiesCidrParameter {
	return v.value
}

func (v *NullableFindIPAvailabilitiesCidrParameter) Set(val *FindIPAvailabilitiesCidrParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableFindIPAvailabilitiesCidrParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableFindIPAvailabilitiesCidrParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindIPAvailabilitiesCidrParameter(val *FindIPAvailabilitiesCidrParameter) *NullableFindIPAvailabilitiesCidrParameter {
	return &NullableFindIPAvailabilitiesCidrParameter{value: val, isSet: true}
}

func (v NullableFindIPAvailabilitiesCidrParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindIPAvailabilitiesCidrParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
