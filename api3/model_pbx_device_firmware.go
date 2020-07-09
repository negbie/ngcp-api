/*
 * NGCP API
 *
 * Sipwise NGCP API (role admin)
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PbxDeviceFirmware struct for PbxDeviceFirmware
type PbxDeviceFirmware struct {
	// The pbx device model id this firmware belongs to.
	DeviceId float32 `json:"device_id,omitempty"`
	// The version number of this firmware.
	Version string `json:"version,omitempty"`
	// A custom tag for this firmware.
	Tag string `json:"tag"`
	// The filename of this firmware.
	Filename string `json:"filename,omitempty"`
}