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

// ChangeaspecificUpnRewriteSetReader is a Reader for the ChangeaspecificUpnRewriteSet structure.
type ChangeaspecificUpnRewriteSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeaspecificUpnRewriteSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeaspecificUpnRewriteSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeaspecificUpnRewriteSetOK creates a ChangeaspecificUpnRewriteSetOK with default headers values
func NewChangeaspecificUpnRewriteSetOK() *ChangeaspecificUpnRewriteSetOK {
	return &ChangeaspecificUpnRewriteSetOK{}
}

/*ChangeaspecificUpnRewriteSetOK handles this case with default header values.

ChangeaspecificUpnRewriteSetOK changeaspecific upn rewrite set o k
*/
type ChangeaspecificUpnRewriteSetOK struct {
	Payload string
}

func (o *ChangeaspecificUpnRewriteSetOK) Error() string {
	return fmt.Sprintf("[PATCH /upnrewritesets/{id}][%d] changeaspecificUpnRewriteSetOK  %+v", 200, o.Payload)
}

func (o *ChangeaspecificUpnRewriteSetOK) GetPayload() string {
	return o.Payload
}

func (o *ChangeaspecificUpnRewriteSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
