/*
 * NGCP API
 *
 * Sipwise NGCP API (role admin)
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// Interception struct for Interception
type Interception struct {
	// The IP address of the X-2 interface.
	X2Host string `json:"x2_host,omitempty"`
	// The username for authenticating on the X-2 interface.
	X2User string `json:"x2_user"`
	// The password for authenticating on the X-2 interface.
	X2Password string `json:"x2_password"`
	// The port of the X-2 interface.
	X2Port float32 `json:"x2_port,omitempty"`
	// Whether to also intercept call content via X-3 interface (false by default).
	X3Required bool `json:"x3_required"`
	// The port of the X-3 interface.
	X3Port float32 `json:"x3_port"`
	// The number to intercept.
	Number string `json:"number,omitempty"`
	// The IP address of the X-3 interface.
	X3Host string `json:"x3_host"`
	// The LI ID for this interception.
	Liid float32 `json:"liid,omitempty"`
}
