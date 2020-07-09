// Code generated by go-swagger; DO NOT EDIT.

package topuplogs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api1/models"
)

// GetaspecificTopupLogReader is a Reader for the GetaspecificTopupLog structure.
type GetaspecificTopupLogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetaspecificTopupLogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetaspecificTopupLogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetaspecificTopupLogOK creates a GetaspecificTopupLogOK with default headers values
func NewGetaspecificTopupLogOK() *GetaspecificTopupLogOK {
	return &GetaspecificTopupLogOK{}
}

/*GetaspecificTopupLogOK handles this case with default header values.

GetaspecificTopupLogOK getaspecific topup log o k
*/
type GetaspecificTopupLogOK struct {
	Payload *models.TopupLogs
}

func (o *GetaspecificTopupLogOK) Error() string {
	return fmt.Sprintf("[GET /topuplogs/{id}][%d] getaspecificTopupLogOK  %+v", 200, o.Payload)
}

func (o *GetaspecificTopupLogOK) GetPayload() *models.TopupLogs {
	return o.Payload
}

func (o *GetaspecificTopupLogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TopupLogs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
