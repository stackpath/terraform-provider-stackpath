/*
 * Object Storage
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api_client

// StackpathapiStatus struct for StackpathapiStatus
type StackpathapiStatus struct {
	Code    int32                      `json:"code,omitempty"`
	Error   string                     `json:"error,omitempty"`
	Details []InlineResponse401Details `json:"details,omitempty"`
}