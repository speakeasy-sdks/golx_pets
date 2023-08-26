// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/golx_pets/pkg/models/shared"
	"net/http"
)

type UpdateUserFormRequest struct {
	// Update an existent user in the store
	User *shared.User `request:"mediaType=application/x-www-form-urlencoded"`
	// name that need to be deleted
	Username string `pathParam:"style=simple,explode=false,name=username"`
}

func (o *UpdateUserFormRequest) GetUser() *shared.User {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *UpdateUserFormRequest) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

type UpdateUserFormResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *UpdateUserFormResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateUserFormResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateUserFormResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
