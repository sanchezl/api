// +build !ignore_autogenerated

// Code generated by openshift-compatibility-gen. DO NOT EDIT.

package v1

// CompatibilityLevel is an autogenerated function, returning the OpenShift API compatibility level.
// It is controlled by the "openshift:compatibility-gen:level" tag in types.go.
func (in *RangeAllocation) CompatibilityLevel() (level int) {
	return 1
}

// Internal is an autogenerated function, returning the true if the type is internal to OpenShift and not exposed as a supported API .
// It is controlled by the "openshift:compatibility-gen:internal" tag in types.go.
func (in *RangeAllocation) Internal() bool {
	return false
}

// CompatibilityLevel is an autogenerated function, returning the OpenShift API compatibility level.
// It is controlled by the "openshift:compatibility-gen:level" tag in types.go.
func (in *RangeAllocationList) CompatibilityLevel() (level int) {
	return 1
}

// Internal is an autogenerated function, returning the true if the type is internal to OpenShift and not exposed as a supported API .
// It is controlled by the "openshift:compatibility-gen:internal" tag in types.go.
func (in *RangeAllocationList) Internal() bool {
	return false
}
