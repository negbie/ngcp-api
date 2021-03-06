/*
 * NGCP API
 *
 * Sipwise NGCP API (role admin)
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SoundFileRecording struct for SoundFileRecording
type SoundFileRecording struct {
	// The sound set the sound file belongs to.
	SetId float32 `json:"set_id,omitempty"`
	// The filename of the sound file (for informational purposes only).
	Filename string `json:"filename,omitempty"`
	// The sound handle when to play this sound file.
	Handle string `json:"handle,omitempty"`
	// Play file in a loop.
	Loopplay bool `json:"loopplay,omitempty"`
}
