// Package xoapicodegenonlyhonourgoname provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/CeskyPane/oapi-codegen/v2 version v2.0.0-00010101000000-000000000000 DO NOT EDIT.
package xoapicodegenonlyhonourgoname

// TypeWithUnexportedField A struct will be output where one of the fields is not exported
type TypeWithUnexportedField struct {
	accountIdentifier *string `json:"-"`
	Name              *string `json:"name,omitempty"`
}
