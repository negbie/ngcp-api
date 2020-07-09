// Code generated by go-swagger; DO NOT EDIT.

package nr_id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api1/models"
)

// GetaspecificEmailTemplateReader is a Reader for the GetaspecificEmailTemplate structure.
type GetaspecificEmailTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetaspecificEmailTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetaspecificEmailTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetaspecificEmailTemplateOK creates a GetaspecificEmailTemplateOK with default headers values
func NewGetaspecificEmailTemplateOK() *GetaspecificEmailTemplateOK {
	return &GetaspecificEmailTemplateOK{}
}

/*GetaspecificEmailTemplateOK handles this case with default header values.

GetaspecificEmailTemplateOK getaspecific email template o k
*/
type GetaspecificEmailTemplateOK struct {
	Payload *models.EmailTemplates
}

func (o *GetaspecificEmailTemplateOK) Error() string {
	return fmt.Sprintf("[GET /emailtemplates/{id}][%d] getaspecificEmailTemplateOK  %+v", 200, o.Payload)
}

func (o *GetaspecificEmailTemplateOK) GetPayload() *models.EmailTemplates {
	return o.Payload
}

func (o *GetaspecificEmailTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EmailTemplates)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}