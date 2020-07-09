// Code generated by go-swagger; DO NOT EDIT.

package invoicetemplates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api1/models"
)

// GetInvoiceTemplateitemsReader is a Reader for the GetInvoiceTemplateitems structure.
type GetInvoiceTemplateitemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvoiceTemplateitemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInvoiceTemplateitemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInvoiceTemplateitemsOK creates a GetInvoiceTemplateitemsOK with default headers values
func NewGetInvoiceTemplateitemsOK() *GetInvoiceTemplateitemsOK {
	return &GetInvoiceTemplateitemsOK{}
}

/*GetInvoiceTemplateitemsOK handles this case with default header values.

GetInvoiceTemplateitemsOK get invoice templateitems o k
*/
type GetInvoiceTemplateitemsOK struct {
	Payload []*models.InvoiceTemplate
}

func (o *GetInvoiceTemplateitemsOK) Error() string {
	return fmt.Sprintf("[GET /invoicetemplates][%d] getInvoiceTemplateitemsOK  %+v", 200, o.Payload)
}

func (o *GetInvoiceTemplateitemsOK) GetPayload() []*models.InvoiceTemplate {
	return o.Payload
}

func (o *GetInvoiceTemplateitemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
