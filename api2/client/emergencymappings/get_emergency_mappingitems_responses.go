// Code generated by go-swagger; DO NOT EDIT.

package emergencymappings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api/models"
)

// GetEmergencyMappingitemsReader is a Reader for the GetEmergencyMappingitems structure.
type GetEmergencyMappingitemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEmergencyMappingitemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEmergencyMappingitemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEmergencyMappingitemsOK creates a GetEmergencyMappingitemsOK with default headers values
func NewGetEmergencyMappingitemsOK() *GetEmergencyMappingitemsOK {
	return &GetEmergencyMappingitemsOK{}
}

/*GetEmergencyMappingitemsOK handles this case with default header values.

GetEmergencyMappingitemsOK get emergency mappingitems o k
*/
type GetEmergencyMappingitemsOK struct {
	Payload []*models.EmergencyMappings
}

func (o *GetEmergencyMappingitemsOK) Error() string {
	return fmt.Sprintf("[GET /emergencymappings][%d] getEmergencyMappingitemsOK  %+v", 200, o.Payload)
}

func (o *GetEmergencyMappingitemsOK) GetPayload() []*models.EmergencyMappings {
	return o.Payload
}

func (o *GetEmergencyMappingitemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
