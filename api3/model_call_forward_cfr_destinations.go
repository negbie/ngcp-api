/*
 * NGCP API
 *
 * Sipwise NGCP API (role admin)
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CallForwardCfrDestinations struct for CallForwardCfrDestinations
type CallForwardCfrDestinations struct {
	// Priority
	Priority float32 `json:"priority"`
	// Simple destination
	SimpleDestination string `json:"simple_destination"`
	// Destination
	Destination string `json:"destination"`
	// Announcement id
	AnnouncementId float32 `json:"announcement_id"`
	// Timeout
	Timeout float32 `json:"timeout"`
}
