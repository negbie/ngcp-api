/*
 * NGCP API
 *
 * Sipwise NGCP API (role admin)
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// TopupCash struct for TopupCash
type TopupCash struct {
	// The subscriber for which to topup the balance.
	SubscriberId float32 `json:"subscriber_id,omitempty"`
	// The amount to top up in cents of Euro/USD/etc.
	Amount float32 `json:"amount,omitempty"`
	// An external ID to identify the top-up request in the top-up log.
	RequestToken string `json:"request_token"`
	// The billing package to switch to after topup.
	PackageId float32 `json:"package_id"`
}
