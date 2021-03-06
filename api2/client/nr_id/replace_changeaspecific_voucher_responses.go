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

// ReplaceChangeaspecificVoucherReader is a Reader for the ReplaceChangeaspecificVoucher structure.
type ReplaceChangeaspecificVoucherReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceChangeaspecificVoucherReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceChangeaspecificVoucherOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceChangeaspecificVoucherOK creates a ReplaceChangeaspecificVoucherOK with default headers values
func NewReplaceChangeaspecificVoucherOK() *ReplaceChangeaspecificVoucherOK {
	return &ReplaceChangeaspecificVoucherOK{}
}

/*ReplaceChangeaspecificVoucherOK handles this case with default header values.

ReplaceChangeaspecificVoucherOK replace changeaspecific voucher o k
*/
type ReplaceChangeaspecificVoucherOK struct {
	Payload *models.Vouchers
}

func (o *ReplaceChangeaspecificVoucherOK) Error() string {
	return fmt.Sprintf("[PUT /vouchers/{id}][%d] replaceChangeaspecificVoucherOK  %+v", 200, o.Payload)
}

func (o *ReplaceChangeaspecificVoucherOK) GetPayload() *models.Vouchers {
	return o.Payload
}

func (o *ReplaceChangeaspecificVoucherOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Vouchers)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
