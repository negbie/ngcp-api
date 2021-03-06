// Code generated by go-swagger; DO NOT EDIT.

package nr_id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api/models"
)

// GetaspecificCustomerFraudPreferenceReader is a Reader for the GetaspecificCustomerFraudPreference structure.
type GetaspecificCustomerFraudPreferenceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetaspecificCustomerFraudPreferenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetaspecificCustomerFraudPreferenceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetaspecificCustomerFraudPreferenceOK creates a GetaspecificCustomerFraudPreferenceOK with default headers values
func NewGetaspecificCustomerFraudPreferenceOK() *GetaspecificCustomerFraudPreferenceOK {
	return &GetaspecificCustomerFraudPreferenceOK{}
}

/*GetaspecificCustomerFraudPreferenceOK handles this case with default header values.

GetaspecificCustomerFraudPreferenceOK getaspecific customer fraud preference o k
*/
type GetaspecificCustomerFraudPreferenceOK struct {
	Payload *models.CustomerFraudPreferences
}

func (o *GetaspecificCustomerFraudPreferenceOK) Error() string {
	return fmt.Sprintf("[GET /customerfraudpreferences/{id}][%d] getaspecificCustomerFraudPreferenceOK  %+v", 200, o.Payload)
}

func (o *GetaspecificCustomerFraudPreferenceOK) GetPayload() *models.CustomerFraudPreferences {
	return o.Payload
}

func (o *GetaspecificCustomerFraudPreferenceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomerFraudPreferences)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
