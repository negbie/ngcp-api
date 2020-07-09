// Code generated by go-swagger; DO NOT EDIT.

package managersecretary

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api1/models"
)

// GetManagerSecretaryitemsReader is a Reader for the GetManagerSecretaryitems structure.
type GetManagerSecretaryitemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetManagerSecretaryitemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetManagerSecretaryitemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetManagerSecretaryitemsOK creates a GetManagerSecretaryitemsOK with default headers values
func NewGetManagerSecretaryitemsOK() *GetManagerSecretaryitemsOK {
	return &GetManagerSecretaryitemsOK{}
}

/*GetManagerSecretaryitemsOK handles this case with default header values.

GetManagerSecretaryitemsOK get manager secretaryitems o k
*/
type GetManagerSecretaryitemsOK struct {
	Payload []*models.ManagerSecretary
}

func (o *GetManagerSecretaryitemsOK) Error() string {
	return fmt.Sprintf("[GET /managersecretary][%d] getManagerSecretaryitemsOK  %+v", 200, o.Payload)
}

func (o *GetManagerSecretaryitemsOK) GetPayload() []*models.ManagerSecretary {
	return o.Payload
}

func (o *GetManagerSecretaryitemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
