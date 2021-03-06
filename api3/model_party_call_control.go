/*
 * NGCP API
 *
 * Sipwise NGCP API (role admin)
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PartyCallControl struct for PartyCallControl
type PartyCallControl struct {
	// Caller number
	Caller string `json:"caller,omitempty"`
	// External call control request type
	Type string `json:"type,omitempty"`
	// Call id
	Callid string `json:"callid,omitempty"`
	// Call status
	Status string `json:"status,omitempty"`
	// Session related token
	Token string `json:"token,omitempty"`
	// Callee number
	Callee string `json:"callee,omitempty"`
}
