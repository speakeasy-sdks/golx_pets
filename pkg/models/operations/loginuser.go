// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type LoginUserRequest struct {
	// The password for login in clear text
	Password *string `queryParam:"style=form,explode=true,name=password"`
	// The user name for login
	Username *string `queryParam:"style=form,explode=true,name=username"`
}

func (o *LoginUserRequest) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *LoginUserRequest) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

type LoginUserResponse struct {
	// successful operation
	TwoHundredApplicationJSONRes *string
	// successful operation
	TwoHundredApplicationXMLRes *string
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *LoginUserResponse) GetTwoHundredApplicationJSONRes() *string {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONRes
}

func (o *LoginUserResponse) GetTwoHundredApplicationXMLRes() *string {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationXMLRes
}

func (o *LoginUserResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *LoginUserResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *LoginUserResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *LoginUserResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
