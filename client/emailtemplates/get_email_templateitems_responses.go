// Code generated by go-swagger; DO NOT EDIT.

package emailtemplates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api1/models"
)

// GetEmailTemplateitemsReader is a Reader for the GetEmailTemplateitems structure.
type GetEmailTemplateitemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEmailTemplateitemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEmailTemplateitemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEmailTemplateitemsOK creates a GetEmailTemplateitemsOK with default headers values
func NewGetEmailTemplateitemsOK() *GetEmailTemplateitemsOK {
	return &GetEmailTemplateitemsOK{}
}

/*GetEmailTemplateitemsOK handles this case with default header values.

GetEmailTemplateitemsOK get email templateitems o k
*/
type GetEmailTemplateitemsOK struct {
	Payload []*models.EmailTemplates
}

func (o *GetEmailTemplateitemsOK) Error() string {
	return fmt.Sprintf("[GET /emailtemplates][%d] getEmailTemplateitemsOK  %+v", 200, o.Payload)
}

func (o *GetEmailTemplateitemsOK) GetPayload() []*models.EmailTemplates {
	return o.Payload
}

func (o *GetEmailTemplateitemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
