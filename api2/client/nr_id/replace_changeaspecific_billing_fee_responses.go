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

// ReplaceChangeaspecificBillingFeeReader is a Reader for the ReplaceChangeaspecificBillingFee structure.
type ReplaceChangeaspecificBillingFeeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceChangeaspecificBillingFeeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewReplaceChangeaspecificBillingFeeNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceChangeaspecificBillingFeeNoContent creates a ReplaceChangeaspecificBillingFeeNoContent with default headers values
func NewReplaceChangeaspecificBillingFeeNoContent() *ReplaceChangeaspecificBillingFeeNoContent {
	return &ReplaceChangeaspecificBillingFeeNoContent{}
}

/*ReplaceChangeaspecificBillingFeeNoContent handles this case with default header values.

ReplaceChangeaspecificBillingFeeNoContent replace changeaspecific billing fee no content
*/
type ReplaceChangeaspecificBillingFeeNoContent struct {
	Payload interface{}
}

func (o *ReplaceChangeaspecificBillingFeeNoContent) Error() string {
	return fmt.Sprintf("[PUT /billingfees/{id}][%d] replaceChangeaspecificBillingFeeNoContent  %+v", 204, o.Payload)
}

func (o *ReplaceChangeaspecificBillingFeeNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ReplaceChangeaspecificBillingFeeNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
