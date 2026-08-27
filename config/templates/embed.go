package templates

import _ "embed"

// ControllerTemplate is the template used to generate the managed resource
// controllers. It differs from the upjet default by wiring the
// ReconciliationPolicy of the ProviderConfig into the reconciler, so that the
// back-off of a resource that keeps failing can be configured.
//
//go:embed controller.go.tmpl
var ControllerTemplate string
