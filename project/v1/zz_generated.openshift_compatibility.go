// +build !ignore_autogenerated

// Code generated by openshift-compatibility-gen. DO NOT EDIT.

package v1

// CompatibilityLevel is an autogenerated function, returning the OpenShift API compatibility level.
// It is controlled by the "openshift:compatibility-gen:level" tag in types.go.
// Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).
func (in *Project) CompatibilityLevel() (level int) {
	return 1
}

// Exposed is an autogenerated function, returning the true if the type is exposed as an OpenShift API.
// It is controlled by the "openshift:compatibility-gen:exposed" tag in types.go.
func (in *Project) Exposed() bool {
	return true
}

// CompatibilityLevel is an autogenerated function, returning the OpenShift API compatibility level.
// It is controlled by the "openshift:compatibility-gen:level" tag in types.go.
// Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).
func (in *ProjectList) CompatibilityLevel() (level int) {
	return 1
}

// Exposed is an autogenerated function, returning the true if the type is exposed as an OpenShift API.
// It is controlled by the "openshift:compatibility-gen:exposed" tag in types.go.
func (in *ProjectList) Exposed() bool {
	return true
}

// CompatibilityLevel is an autogenerated function, returning the OpenShift API compatibility level.
// It is controlled by the "openshift:compatibility-gen:level" tag in types.go.
// Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).
func (in *ProjectRequest) CompatibilityLevel() (level int) {
	return 1
}

// Exposed is an autogenerated function, returning the true if the type is exposed as an OpenShift API.
// It is controlled by the "openshift:compatibility-gen:exposed" tag in types.go.
func (in *ProjectRequest) Exposed() bool {
	return true
}
