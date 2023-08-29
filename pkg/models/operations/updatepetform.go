// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/golx_pets/pkg/models/shared"
	"net/http"
)

type UpdatePetFormSecurity struct {
	PetstoreAuth string `security:"scheme,type=oauth2,name=Authorization"`
}

func (o *UpdatePetFormSecurity) GetPetstoreAuth() string {
	if o == nil {
		return ""
	}
	return o.PetstoreAuth
}

type UpdatePetFormResponse struct {
	Body        []byte
	ContentType string
	// Successful operation
	Pet         *shared.Pet
	StatusCode  int
	RawResponse *http.Response
}

func (o *UpdatePetFormResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *UpdatePetFormResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdatePetFormResponse) GetPet() *shared.Pet {
	if o == nil {
		return nil
	}
	return o.Pet
}

func (o *UpdatePetFormResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdatePetFormResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
