/*
 * NGCP API
 *
 * Sipwise NGCP API (role admin)
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// AutoAttendant struct for AutoAttendant
type AutoAttendant struct {
	// IVR Slots. A list of items each containing the keys slot and destination. \"slot\" from 0-9. \"destination\" can be a number, username or full SIP URI.
	Slots []AutoAttendantSlots `json:"slots"`
}
