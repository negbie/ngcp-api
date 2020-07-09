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

// ChangeaspecificEmergencyMappingReader is a Reader for the ChangeaspecificEmergencyMapping structure.
type ChangeaspecificEmergencyMappingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeaspecificEmergencyMappingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewChangeaspecificEmergencyMappingNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeaspecificEmergencyMappingNoContent creates a ChangeaspecificEmergencyMappingNoContent with default headers values
func NewChangeaspecificEmergencyMappingNoContent() *ChangeaspecificEmergencyMappingNoContent {
	return &ChangeaspecificEmergencyMappingNoContent{}
}

/*ChangeaspecificEmergencyMappingNoContent handles this case with default header values.

ChangeaspecificEmergencyMappingNoContent changeaspecific emergency mapping no content
*/
type ChangeaspecificEmergencyMappingNoContent struct {
	Payload interface{}
}

func (o *ChangeaspecificEmergencyMappingNoContent) Error() string {
	return fmt.Sprintf("[PATCH /emergencymappings/{id}][%d] changeaspecificEmergencyMappingNoContent  %+v", 204, o.Payload)
}

func (o *ChangeaspecificEmergencyMappingNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ChangeaspecificEmergencyMappingNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
