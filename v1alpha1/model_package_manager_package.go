/*
 * Grafeas API
 *
 * An API to insert and retrieve annotations on cloud artifacts.
 *
 * API version: v1alpha1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

// This represents a particular package that is distributed over various channels. e.g. glibc (aka libc6) is distributed by many, at various versions.
type PackageManagerPackage struct {
	// The name of the package.
	Name string `json:"name,omitempty"`
	// The various channels by which a package is distributed.
	Distribution []PackageManagerDistribution `json:"distribution,omitempty"`
}
