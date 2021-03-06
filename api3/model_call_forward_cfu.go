/*
 * NGCP API
 *
 * Sipwise NGCP API (role admin)
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CallForwardCfu Call Forward Unconditional, Contains the keys \"destinations\", \"times\" and \"sources\". \"destinations\" is an Array of Objects having a \"destination\", \"priority\" and \"timeout\" field. \"times\" is an Array of Objects having the fields \"minute\", \"hour\", \"wday\", \"mday\", \"month\", \"year\". \"times\" can be empty, then the CF is applied always. \"sources\" is an Array of Objects having one field \"source\". \"sources\" can be empty.
type CallForwardCfu struct {
	// Sources
	Sources []CallForwardCftSources `json:"sources"`
	// Sources mode
	SourcesMode string `json:"sources_mode"`
	// Bnumbers
	Bnumbers []CallForwardCftBnumbers `json:"bnumbers"`
	// Times
	Times []CallForwardCfuTimes `json:"times"`
	// Bnumbers is regex
	BnumbersIsRegex bool `json:"bnumbers_is_regex"`
	// Bnumbers mode
	BnumbersMode string `json:"bnumbers_mode"`
	// Destinations
	Destinations []CallForwardCfuDestinations `json:"destinations"`
	// Sources is regex
	SourcesIsRegex bool `json:"sources_is_regex"`
}
