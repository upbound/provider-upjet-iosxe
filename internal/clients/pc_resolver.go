package clients

import (
	"context"

	"github.com/crossplane/crossplane-runtime/v2/pkg/errors"
	xpresource "github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/apis/configuration/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ReconciliationPolicy returns the reconciliation policy configured in the
// ProviderConfig referenced by the given managed resource. The generated
// controllers use it to configure the rate limiting of the resources that keep
// failing.
func ReconciliationPolicy(ctx context.Context, client client.Client, mg xpresource.Managed) (*v1alpha1.ReconciliationPolicy, error) {
	spec, err := resolveProviderConfig(ctx, client, mg)
	if err != nil {
		return nil, errors.Wrap(err, "cannot resolve the referenced ProviderConfig")
	}
	return spec.ReconciliationPolicy, nil
}
