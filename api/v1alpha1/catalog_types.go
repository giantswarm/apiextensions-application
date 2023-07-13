package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/k8smetadata/pkg/annotation"
)

const (
	kindCatalog              = "Catalog"
	catalogDocumentationLink = "https://docs.giantswarm.io/ui-api/management-api/crd/catalogs.application.giantswarm.io/"
)

func NewCatalogTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       kindCatalog,
	}
}

// NewCatalogCR returns an Catalog Custom Resource.
func NewCatalogCR() *Catalog {
	return &Catalog{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				annotation.Docs: catalogDocumentationLink,
			},
		},
		TypeMeta: NewCatalogTypeMeta(),
	}
}

// +kubebuilder:printcolumn:name="Catalog URL",type=string,JSONPath=`.spec.storage.URL`,description="URL of the catalog"
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`,description="Time since created"
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=common;giantswarm
// +kubebuilder:storageversion
// +k8s:openapi-gen=true
// Catalog represents a catalog of managed apps. It stores general information for potential apps to install.
// It is reconciled by app-operator.
type Catalog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CatalogSpec `json:"spec"`
	// +kubebuilder:validation:Optional
	Status CatalogStatus `json:"status"`
}

// +k8s:openapi-gen=true
type CatalogSpec struct {
	// Title is the name of the catalog for this CR
	// e.g. Catalog of Apps by Giant Swarm
	Title       string `json:"title"`
	Description string `json:"description"`
	// +kubebuilder:validation:Optional
	// +nullable
	// Config is the config to be applied when apps belonging to this
	// catalog are deployed.
	Config *CatalogSpecConfig `json:"config,omitempty"`
	// LogoURL contains the links for logo image file for this catalog
	LogoURL string `json:"logoURL"`
	// Storage references an object defining catalog repository.
	// This field is deprecated and replaced by Repositories.
	Storage CatalogSpecStorage `json:"storage"`
	// +kubebuilder:validation:MinItems=1
	// Repositories is an array of objects defining catalog repositories.
	Repositories []CatalogSpecRepository `json:"repositories"`
}

// +k8s:openapi-gen=true
type CatalogSpecConfig struct {
	// +kubebuilder:validation:Optional
	// +nullable
	// ConfigMap references a config map containing catalog values that
	// should be applied to apps in this catalog.
	ConfigMap *CatalogSpecConfigConfigMap `json:"configMap,omitempty"`
	// +kubebuilder:validation:Optional
	// +nullable
	// Secret references a secret containing catalog values that should be
	// applied to apps in this catalog.
	Secret *CatalogSpecConfigSecret `json:"secret,omitempty"`
}

// +k8s:openapi-gen=true
type CatalogSpecConfigConfigMap struct {
	// Name is the name of the config map containing catalog values to
	// apply, e.g. app-catalog-values.
	Name string `json:"name"`
	// Namespace is the namespace of the catalog values config map,
	// e.g. giantswarm.
	Namespace string `json:"namespace"`
}

// +k8s:openapi-gen=true
type CatalogSpecConfigSecret struct {
	// Name is the name of the secret containing catalog values to apply,
	// e.g. app-catalog-secret.
	Name string `json:"name"`
	// Namespace is the namespace of the secret,
	// e.g. giantswarm.
	Namespace string `json:"namespace"`
}

// +k8s:openapi-gen=true
type CatalogSpecStorage struct {
	// Type indicates which repository type would be used for this Catalog.
	// e.g. helm
	Type string `json:"type"`
	// URL is the link to where this Catalog's repository is located
	// e.g. https://example.com/app-catalog/
	URL string `json:"URL"`
}

// +k8s:openapi-gen=true
type CatalogSpecRepository struct {
	// Type indicates which repository type would be used for this Catalog.
	// e.g. helm
	Type string `json:"type"`
	// URL is the link to where this Catalog's repository is located
	// e.g. https://example.com/app-catalog/
	URL string `json:"URL"`
}

// CatalogStatus represents the current state of the catalog.
type CatalogStatus struct {
	// HelmRepositoryList contains the list of Flux HelmRepository custom resources
	// that have been successfully created from the Catalog object.
	// +optional
	HelmRepositoryList *HelmRepositoryList `json:"helmRepositoryList,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CatalogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Catalog `json:"items"`
}

// HelmRepositoryList carries the list of Flux HelmRepository custom resources
// that have been successfully created from the Catalog object.
type HelmRepositoryList struct {
	// Entries of HelmRepository custom resources.
	Entries []HelmRepositoryRef `json:"entries"`
}

// HelmRepositoryRef represents a basic HelmRepository custom resource information.
type HelmRepositoryRef struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}
