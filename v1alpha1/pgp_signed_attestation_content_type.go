/*
 * v1alpha1/proto/grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas
// PgpSignedAttestationContentType : Type (e.g. schema) of the attestation payload that was signed.   - CONTENT_TYPE_UNSPECIFIED: ContentType is not set.  - SIMPLE_SIGNING_JSON: Atomic format attestation signature. See https://github.com/containers/image/blob/8a5d2f82a6e3263290c8e0276c3e0f64e77723e7/docs/atomic-signature.md The payload extracted from `signature` is a JSON blob conforming to the linked schema.
type PgpSignedAttestationContentType string

// List of PgpSignedAttestationContentType
const (
	CONTENT_TYPE_UNSPECIFIED PgpSignedAttestationContentType = "CONTENT_TYPE_UNSPECIFIED"
	SIMPLE_SIGNING_JSON PgpSignedAttestationContentType = "SIMPLE_SIGNING_JSON"
)
