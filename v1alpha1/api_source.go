/*
 * v1alpha1/proto/grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

// Source describes the location of the source used for the build.
type ApiSource struct {

	// If provided, get the source from this location in in Google Cloud Storage.
	StorageSource *ApiStorageSource `json:"storage_source,omitempty"`

	// If provided, get source from this location in a Cloud Repo.
	RepoSource *ApiRepoSource `json:"repo_source,omitempty"`

	// If provided, the input binary artifacts for the build came from this location.
	ArtifactStorageSource *ApiStorageSource `json:"artifact_storage_source,omitempty"`

	// Hash(es) of the build source, which can be used to verify that the original source integrity was maintained in the build.  The keys to this map are file paths used as build source and the values contain the hash values for those files.  If the build source came in a single package such as a gzipped tarfile (.tar.gz), the FileHash will be for the single path to that file.
	FileHashes map[string]ApiFileHashes `json:"file_hashes,omitempty"`

	// If provided, the source code used for the build came from this location.
	Context *ApiSourceContext `json:"context,omitempty"`

	// If provided, some of the source code used for the build may be found in these locations, in the case where the source repository had multiple remotes or submodules. This list will not include the context specified in the context field.
	AdditionalContexts []ApiSourceContext `json:"additional_contexts,omitempty"`
}
