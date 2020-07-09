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

// GetaspecificBannedUserReader is a Reader for the GetaspecificBannedUser structure.
type GetaspecificBannedUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetaspecificBannedUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetaspecificBannedUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetaspecificBannedUserOK creates a GetaspecificBannedUserOK with default headers values
func NewGetaspecificBannedUserOK() *GetaspecificBannedUserOK {
	return &GetaspecificBannedUserOK{}
}

/*GetaspecificBannedUserOK handles this case with default header values.

GetaspecificBannedUserOK getaspecific banned user o k
*/
type GetaspecificBannedUserOK struct {
	Payload string
}

func (o *GetaspecificBannedUserOK) Error() string {
	return fmt.Sprintf("[GET /bannedusers/{id}][%d] getaspecificBannedUserOK  %+v", 200, o.Payload)
}

func (o *GetaspecificBannedUserOK) GetPayload() string {
	return o.Payload
}

func (o *GetaspecificBannedUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}