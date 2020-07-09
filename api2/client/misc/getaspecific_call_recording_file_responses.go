// Code generated by go-swagger; DO NOT EDIT.

package misc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api/models"
)

// GetaspecificCallRecordingFileReader is a Reader for the GetaspecificCallRecordingFile structure.
type GetaspecificCallRecordingFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetaspecificCallRecordingFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetaspecificCallRecordingFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetaspecificCallRecordingFileOK creates a GetaspecificCallRecordingFileOK with default headers values
func NewGetaspecificCallRecordingFileOK() *GetaspecificCallRecordingFileOK {
	return &GetaspecificCallRecordingFileOK{}
}

/*GetaspecificCallRecordingFileOK handles this case with default header values.

GetaspecificCallRecordingFileOK getaspecific call recording file o k
*/
type GetaspecificCallRecordingFileOK struct {
	Payload *models.CallRecordingFiles
}

func (o *GetaspecificCallRecordingFileOK) Error() string {
	return fmt.Sprintf("[GET /callrecordingfiles/{id}][%d] getaspecificCallRecordingFileOK  %+v", 200, o.Payload)
}

func (o *GetaspecificCallRecordingFileOK) GetPayload() *models.CallRecordingFiles {
	return o.Payload
}

func (o *GetaspecificCallRecordingFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CallRecordingFiles)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}