/*
 * v1alpha1/proto/grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas
// PackageManagerArchitecture : Instruction set architectures supported by various package managers.   - ARCHITECTURE_UNSPECIFIED: Unknown architecture  - X86: X86 architecture  - X64: X64 architecture
type PackageManagerArchitecture string

// List of PackageManagerArchitecture
const (
	ARCHITECTURE_UNSPECIFIED PackageManagerArchitecture = "ARCHITECTURE_UNSPECIFIED"
	X86 PackageManagerArchitecture = "X86"
	X64 PackageManagerArchitecture = "X64"
)
