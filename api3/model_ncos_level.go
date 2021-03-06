/*
 * NGCP API
 *
 * Sipwise NGCP API (role admin)
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// NcosLevel struct for NcosLevel
type NcosLevel struct {
	// Whether to include check for intra pbx calls within same customer
	IntraPbx bool `json:"intra_pbx"`
	// The level name
	Level string `json:"level,omitempty"`
	// The reseller this level belongs to.
	ResellerId float32 `json:"reseller_id,omitempty"`
	// The level mode (one of blacklist, whitelist)
	Mode string `json:"mode,omitempty"`
	// Whether to include check for calls to local area code
	LocalAc bool `json:"local_ac"`
	// The description of the level
	Description string `json:"description"`
}
