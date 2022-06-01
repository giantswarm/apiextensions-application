package v1alpha1

import (
	v1alpha2 "github.com/giantswarm/apiextensions-application/api/v1alpha2"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// Converts AppCatalog to Hub version (v1alpha2)
func (c *Catalog) ConvertTo(hubRaw conversion.Hub) error {
	hub := hubRaw.(*v1alpha2.AppCatalog)

	// actual conversion
	hub.Spec.Storage = []v1alpha2.CatalogSpecStorage{
		{
			Type: c.Spec.Storage.Type,
			URL:  c.Spec.Storage.URL,
		},
	}

	// copy rest of the properties
	hub.ObjectMeta = c.ObjectMeta
	hub.Spec.Title = c.Spec.Title
	hub.Spec.Description = c.Spec.Description
	hub.Spec.Config = c.Spec.Config
	hub.Spec.LogoURL = c.Spec.LogoURL

	return nil
}

// Converts AppCatalog from Hub (v1alpha2) to Spoke version (v1alpha1)
func (c *Catalog) ConvertFrom(hubRaw conversion.Hub) error {
	hub := hubRaw.(*v1alpha2.AppCatalog)

	// actual conversion
	c.Spec.Storage.Type = hub.Spec.Storage[0].Type
	c.Spec.Storage.URL = hub.Spec.Storage[0].URL

	// copy rest of the properties
	c.ObjectMeta = hub.ObjectMeta
	c.Spec.Title = hub.Spec.Title
	c.Spec.Description = hub.Spec.Description
	c.Spec.Config = hub.Spec.Config
	c.Spec.LogoURL = hub.Spec.LogoURL

	return nil
}
