/*
 * Object Storage
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api_client

// GetCredentialsResponseCredential Storage credentials for a user
type GetCredentialsResponseCredential struct {
	// The ID for the access key
	AccessKey string `json:"accessKey,omitempty"`
}