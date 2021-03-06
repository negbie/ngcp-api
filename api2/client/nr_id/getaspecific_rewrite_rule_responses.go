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

// GetaspecificRewriteRuleReader is a Reader for the GetaspecificRewriteRule structure.
type GetaspecificRewriteRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetaspecificRewriteRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetaspecificRewriteRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetaspecificRewriteRuleOK creates a GetaspecificRewriteRuleOK with default headers values
func NewGetaspecificRewriteRuleOK() *GetaspecificRewriteRuleOK {
	return &GetaspecificRewriteRuleOK{}
}

/*GetaspecificRewriteRuleOK handles this case with default header values.

GetaspecificRewriteRuleOK getaspecific rewrite rule o k
*/
type GetaspecificRewriteRuleOK struct {
	Payload *models.RewriteRules
}

func (o *GetaspecificRewriteRuleOK) Error() string {
	return fmt.Sprintf("[GET /rewriterules/{id}][%d] getaspecificRewriteRuleOK  %+v", 200, o.Payload)
}

func (o *GetaspecificRewriteRuleOK) GetPayload() *models.RewriteRules {
	return o.Payload
}

func (o *GetaspecificRewriteRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RewriteRules)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
