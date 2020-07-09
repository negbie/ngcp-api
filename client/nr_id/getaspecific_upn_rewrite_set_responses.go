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

// GetaspecificUpnRewriteSetReader is a Reader for the GetaspecificUpnRewriteSet structure.
type GetaspecificUpnRewriteSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetaspecificUpnRewriteSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetaspecificUpnRewriteSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetaspecificUpnRewriteSetOK creates a GetaspecificUpnRewriteSetOK with default headers values
func NewGetaspecificUpnRewriteSetOK() *GetaspecificUpnRewriteSetOK {
	return &GetaspecificUpnRewriteSetOK{}
}

/*GetaspecificUpnRewriteSetOK handles this case with default header values.

GetaspecificUpnRewriteSetOK getaspecific upn rewrite set o k
*/
type GetaspecificUpnRewriteSetOK struct {
	Payload string
}

func (o *GetaspecificUpnRewriteSetOK) Error() string {
	return fmt.Sprintf("[GET /upnrewritesets/{id}][%d] getaspecificUpnRewriteSetOK  %+v", 200, o.Payload)
}

func (o *GetaspecificUpnRewriteSetOK) GetPayload() string {
	return o.Payload
}

func (o *GetaspecificUpnRewriteSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
