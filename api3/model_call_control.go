/*
 * NGCP API
 *
 * Sipwise NGCP API (role admin)
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CallControl struct for CallControl
type CallControl struct {
	// The ID of the calling subscriber
	SubscriberId float32 `json:"subscriber_id,omitempty"`
	// The destination URI, user or number as dialed by the end user
	Destination string `json:"destination,omitempty"`
}
