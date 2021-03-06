// Code generated by go-swagger; DO NOT EDIT.

package nr_id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ReplaceChangeaspecificManagerSecretaryReader is a Reader for the ReplaceChangeaspecificManagerSecretary structure.
type ReplaceChangeaspecificManagerSecretaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceChangeaspecificManagerSecretaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewReplaceChangeaspecificManagerSecretaryNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceChangeaspecificManagerSecretaryNoContent creates a ReplaceChangeaspecificManagerSecretaryNoContent with default headers values
func NewReplaceChangeaspecificManagerSecretaryNoContent() *ReplaceChangeaspecificManagerSecretaryNoContent {
	return &ReplaceChangeaspecificManagerSecretaryNoContent{}
}

/*ReplaceChangeaspecificManagerSecretaryNoContent handles this case with default header values.

ReplaceChangeaspecificManagerSecretaryNoContent replace changeaspecific manager secretary no content
*/
type ReplaceChangeaspecificManagerSecretaryNoContent struct {
	Payload interface{}
}

func (o *ReplaceChangeaspecificManagerSecretaryNoContent) Error() string {
	return fmt.Sprintf("[PUT /managersecretary/{id}][%d] replaceChangeaspecificManagerSecretaryNoContent  %+v", 204, o.Payload)
}

func (o *ReplaceChangeaspecificManagerSecretaryNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ReplaceChangeaspecificManagerSecretaryNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
