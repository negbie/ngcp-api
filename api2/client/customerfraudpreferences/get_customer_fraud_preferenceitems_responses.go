// Code generated by go-swagger; DO NOT EDIT.

package customerfraudpreferences

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api/models"
)

// GetCustomerFraudPreferenceitemsReader is a Reader for the GetCustomerFraudPreferenceitems structure.
type GetCustomerFraudPreferenceitemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomerFraudPreferenceitemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomerFraudPreferenceitemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCustomerFraudPreferenceitemsOK creates a GetCustomerFraudPreferenceitemsOK with default headers values
func NewGetCustomerFraudPreferenceitemsOK() *GetCustomerFraudPreferenceitemsOK {
	return &GetCustomerFraudPreferenceitemsOK{}
}

/*GetCustomerFraudPreferenceitemsOK handles this case with default header values.

GetCustomerFraudPreferenceitemsOK get customer fraud preferenceitems o k
*/
type GetCustomerFraudPreferenceitemsOK struct {
	Payload []*models.CustomerFraudPreferences
}

func (o *GetCustomerFraudPreferenceitemsOK) Error() string {
	return fmt.Sprintf("[GET /customerfraudpreferences][%d] getCustomerFraudPreferenceitemsOK  %+v", 200, o.Payload)
}

func (o *GetCustomerFraudPreferenceitemsOK) GetPayload() []*models.CustomerFraudPreferences {
	return o.Payload
}

func (o *GetCustomerFraudPreferenceitemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
