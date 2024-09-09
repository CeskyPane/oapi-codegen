// Package xomitempty provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/CeskyPane/oapi-codegen/v2 version v2.0.0-00010101000000-000000000000 DO NOT EDIT.
package xomitempty

// Client defines model for Client.
type Client struct {
	Id   *float32 `json:"id,omitempty"`
	Name string   `json:"name"`
}

// ClientWithExtension defines model for ClientWithExtension.
type ClientWithExtension struct {
	Id   *float32 `json:"id,omitempty"`
	Name string   `json:"name,omitempty"`
}
