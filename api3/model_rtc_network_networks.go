/*
 * NGCP API
 *
 * Sipwise NGCP API (role admin)
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// RtcNetworkNetworks struct for RtcNetworkNetworks
type RtcNetworkNetworks struct {
	// One of the available options. This defines, to which networks rtc subscribers will be able to connect to.
	Connector string `json:"connector"`
	// An arbitrary hash of additional config contents; e.g. {\"xms\": false}
	Config map[string]interface{} `json:"config"`
	// An arbitrary name, to address that network instance
	Tag string `json:"tag"`
}
