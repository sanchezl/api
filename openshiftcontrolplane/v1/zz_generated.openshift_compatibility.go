// +build !ignore_autogenerated

// Code generated by openshift-compatibility-gen. DO NOT EDIT.

package v1

// CompatibilityLevel is an autogenerated function, returning the OpenShift API compatibility level.
// It is controlled by the "openshift:compatibility-gen:level" tag in types.go.
func (in *BuildDefaultsConfig) CompatibilityLevel() (level int) {
	return 4
}

// Internal is an autogenerated function, returning the true if the type is internal to OpenShift and not exposed as a supported API .
// It is controlled by the "openshift:compatibility-gen:internal" tag in types.go.
func (in *BuildDefaultsConfig) Internal() bool {
	return true
}

// CompatibilityLevel is an autogenerated function, returning the OpenShift API compatibility level.
// It is controlled by the "openshift:compatibility-gen:level" tag in types.go.
func (in *BuildOverridesConfig) CompatibilityLevel() (level int) {
	return 4
}

// Internal is an autogenerated function, returning the true if the type is internal to OpenShift and not exposed as a supported API .
// It is controlled by the "openshift:compatibility-gen:internal" tag in types.go.
func (in *BuildOverridesConfig) Internal() bool {
	return true
}

// CompatibilityLevel is an autogenerated function, returning the OpenShift API compatibility level.
// It is controlled by the "openshift:compatibility-gen:level" tag in types.go.
func (in *OpenShiftAPIServerConfig) CompatibilityLevel() (level int) {
	return 4
}

// Internal is an autogenerated function, returning the true if the type is internal to OpenShift and not exposed as a supported API .
// It is controlled by the "openshift:compatibility-gen:internal" tag in types.go.
func (in *OpenShiftAPIServerConfig) Internal() bool {
	return true
}

// CompatibilityLevel is an autogenerated function, returning the OpenShift API compatibility level.
// It is controlled by the "openshift:compatibility-gen:level" tag in types.go.
func (in *OpenShiftControllerManagerConfig) CompatibilityLevel() (level int) {
	return 4
}

// Internal is an autogenerated function, returning the true if the type is internal to OpenShift and not exposed as a supported API .
// It is controlled by the "openshift:compatibility-gen:internal" tag in types.go.
func (in *OpenShiftControllerManagerConfig) Internal() bool {
	return true
}
