// Package trees provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package trees

import (
	"time"
)

const (
	JWTAuthScopes = "JWTAuth.Scopes"
)

// Error defines model for Error.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// Tree defines model for Tree.
type Tree struct {
	Comment              *string    `json:"comment,omitempty"`
	CreateTime           time.Time  `json:"create_time"`
	Creator              int32      `json:"creator"`
	Description          *string    `json:"description,omitempty"`
	ExternalId           *int32     `json:"external_id,omitempty"`
	Id                   int32      `json:"id"`
	IdValidator          *int32     `json:"id_validator,omitempty"`
	InactivationReason   *string    `json:"inactivation_reason,omitempty"`
	InactivationTime     *time.Time `json:"inactivation_time,omitempty"`
	IsActive             bool       `json:"is_active"`
	IsValidated          *bool      `json:"is_validated,omitempty"`
	LastModificationTime *time.Time `json:"last_modification_time,omitempty"`
	LastModificationUser *int32     `json:"last_modification_user,omitempty"`
	Name                 string     `json:"name"`
}

// TreeList defines model for TreeList.
type TreeList struct {
	CreateTime  time.Time `json:"create_time"`
	Creator     int32     `json:"creator"`
	Description *string   `json:"description,omitempty"`
	ExternalId  *int32    `json:"external_id,omitempty"`
	Id          int32     `json:"id"`
	IsActive    bool      `json:"is_active"`
	Name        string    `json:"name"`
}

// ListParams defines parameters for List.
type ListParams struct {
	// maximum number of results to return
	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}

// CreateJSONBody defines parameters for Create.
type CreateJSONBody = Tree

// CreateJSONRequestBody defines body for Create for application/json ContentType.
type CreateJSONRequestBody = CreateJSONBody
