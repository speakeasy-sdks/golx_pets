// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/golx_pets/pkg/models/shared"
	"net/http"
)

type UploadFileSecurity struct {
	PetstoreAuth string `security:"scheme,type=oauth2,name=Authorization"`
}

func (o *UploadFileSecurity) GetPetstoreAuth() string {
	if o == nil {
		return ""
	}
	return o.PetstoreAuth
}

type UploadFileRequest struct {
	RequestBody []byte `request:"mediaType=application/octet-stream"`
	// Additional Metadata
	AdditionalMetadata *string `queryParam:"style=form,explode=true,name=additionalMetadata"`
	// ID of pet to update
	PetID int64 `pathParam:"style=simple,explode=false,name=petId"`
}

func (o *UploadFileRequest) GetRequestBody() []byte {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *UploadFileRequest) GetAdditionalMetadata() *string {
	if o == nil {
		return nil
	}
	return o.AdditionalMetadata
}

func (o *UploadFileRequest) GetPetID() int64 {
	if o == nil {
		return 0
	}
	return o.PetID
}

type UploadFileResponse struct {
	// successful operation
	APIResponse *shared.APIResponse
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *UploadFileResponse) GetAPIResponse() *shared.APIResponse {
	if o == nil {
		return nil
	}
	return o.APIResponse
}

func (o *UploadFileResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UploadFileResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UploadFileResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
