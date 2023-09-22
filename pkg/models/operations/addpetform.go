// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/golx_pets/pkg/models/shared"
	"net/http"
)

type AddPetFormResponse struct {
	Body        []byte
	ContentType string
	// Successful operation
	Pet         *shared.Pet
	StatusCode  int
	RawResponse *http.Response
}

func (o *AddPetFormResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *AddPetFormResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AddPetFormResponse) GetPet() *shared.Pet {
	if o == nil {
		return nil
	}
	return o.Pet
}

func (o *AddPetFormResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AddPetFormResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
