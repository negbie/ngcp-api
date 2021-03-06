/*
 * NGCP API
 *
 * Sipwise NGCP API (role admin)
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// NcosPattern struct for NcosPattern
type NcosPattern struct {
	// The ncos level this pattern belongs to.
	NcosLevelId float32 `json:"ncos_level_id,omitempty"`
	// Beschreibung
	Description string `json:"description"`
	// Pattern
	Pattern string `json:"pattern,omitempty"`
}
