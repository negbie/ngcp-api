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

// GetaspecificRewriteRuleSetReader is a Reader for the GetaspecificRewriteRuleSet structure.
type GetaspecificRewriteRuleSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetaspecificRewriteRuleSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetaspecificRewriteRuleSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetaspecificRewriteRuleSetOK creates a GetaspecificRewriteRuleSetOK with default headers values
func NewGetaspecificRewriteRuleSetOK() *GetaspecificRewriteRuleSetOK {
	return &GetaspecificRewriteRuleSetOK{}
}

/*GetaspecificRewriteRuleSetOK handles this case with default header values.

GetaspecificRewriteRuleSetOK getaspecific rewrite rule set o k
*/
type GetaspecificRewriteRuleSetOK struct {
	Payload *models.RewriteRuleSets
}

func (o *GetaspecificRewriteRuleSetOK) Error() string {
	return fmt.Sprintf("[GET /rewriterulesets/{id}][%d] getaspecificRewriteRuleSetOK  %+v", 200, o.Payload)
}

func (o *GetaspecificRewriteRuleSetOK) GetPayload() *models.RewriteRuleSets {
	return o.Payload
}

func (o *GetaspecificRewriteRuleSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RewriteRuleSets)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
