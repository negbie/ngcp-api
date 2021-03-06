/*
 * NGCP API
 *
 * Sipwise NGCP API (role admin)
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PeeringServer struct for PeeringServer
type PeeringServer struct {
	// Gewicht
	Weight float32 `json:"weight,omitempty"`
	// The peering group this server belongs to.
	GroupId float32 `json:"group_id,omitempty"`
	// Via Route
	ViaRoute string `json:"via_route"`
	// IP-Adresse
	Ip string `json:"ip,omitempty"`
	// Protokoll
	Transport string `json:"transport"`
	// Host-Name
	Host string `json:"host"`
	// Server enabled state.
	Enabled bool `json:"enabled"`
	// Port
	Port float32 `json:"port,omitempty"`
	// Name
	Name string `json:"name,omitempty"`
	// Perform probing via SIP pings to check reachability. If disabled, calls are always sent to this peer since the status is not known.
	Probe bool `json:"probe"`
}
