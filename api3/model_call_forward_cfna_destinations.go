/*
 * NGCP API
 *
 * Sipwise NGCP API (role admin)
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CallForwardCfnaDestinations struct for CallForwardCfnaDestinations
type CallForwardCfnaDestinations struct {
	// Timeout
	Timeout float32 `json:"timeout"`
	// Announcement id
	AnnouncementId float32 `json:"announcement_id"`
	// Destination
	Destination string `json:"destination"`
	// Simple destination
	SimpleDestination string `json:"simple_destination"`
	// Priority
	Priority float32 `json:"priority"`
}
