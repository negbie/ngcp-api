/*
 * NGCP API
 *
 * Sipwise NGCP API (role admin)
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PbxDeviceLines struct for PbxDeviceLines
type PbxDeviceLines struct {
	// The linerange name to use.
	Linerange string `json:"linerange,omitempty"`
	// The line/key to use (starting from 0)
	KeyNum float32 `json:"key_num,omitempty"`
	// The subscriber to use on this line/key
	SubscriberId float32 `json:"subscriber_id,omitempty"`
	// The type of feature to use on this line/key
	Type string `json:"type,omitempty"`
}