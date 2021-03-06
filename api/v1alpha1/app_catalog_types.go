package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/k8smetadata/pkg/annotation"
)

const (
	kindAppCatalog              = "AppCatalog"
	appCatalogDocumentationLink = "https://docs.giantswarm.io/ui-api/management-api/crd/appcatalogs.application.giantswarm.io/"
)

func NewAppCatalogTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       kindAppCatalog,
	}
}

// NewAppCatalogCR returns an AppCatalog Custom Resource.
func NewAppCatalogCR() *AppCatalog {
	return &AppCatalog{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				annotation.Docs: appCatalogDocumentationLink,
			},
		},
		TypeMeta: NewAppCatalogTypeMeta(),
	}
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=common;giantswarm,scope=Cluster
// +kubebuilder:storageversion
// +k8s:openapi-gen=true
// Deprecated, use Catalog CRD instead.
// AppCatalog represents a catalog of managed apps. It stores general information for potential apps to install.
// It is reconciled by app-operator.
type AppCatalog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppCatalogSpec `json:"spec"`
}

// +k8s:openapi-gen=true
type AppCatalogSpec struct {
	// Title is the name of the app catalog for this CR
	// e.g. Catalog of Apps by Giant Swarm
	Title       string `json:"title"`
	Description string `json:"description"`
	// +kubebuilder:validation:Optional
	// +nullable
	// Config is the config to be applied when apps belonging to this
	// catalog are deployed.
	Config AppCatalogSpecConfig `json:"config,omitempty"`
	// LogoURL contains the links for logo image file for this app catalog
	LogoURL string `json:"logoURL"`
	// Storage references a map containing values that should be applied to
	// the appcatalog.
	Storage AppCatalogSpecStorage `json:"storage"`
}

// +k8s:openapi-gen=true
type AppCatalogSpecConfig struct {
	// +kubebuilder:validation:Optional
	// +nullable
	// ConfigMap references a config map containing catalog values that
	// should be applied to apps in this catalog.
	ConfigMap AppCatalogSpecConfigConfigMap `json:"configMap,omitempty"`
	// +kubebuilder:validation:Optional
	// +nullable
	// Secret references a secret containing catalog values that should be
	// applied to apps in this catalog.
	Secret AppCatalogSpecConfigSecret `json:"secret,omitempty"`
}

// +k8s:openapi-gen=true
type AppCatalogSpecConfigConfigMap struct {
	// Name is the name of the config map containing catalog values to
	// apply, e.g. app-catalog-values.
	Name string `json:"name"`
	// Namespace is the namespace of the catalog values config map,
	// e.g. giantswarm.
	Namespace string `json:"namespace"`
}

// +k8s:openapi-gen=true
type AppCatalogSpecConfigSecret struct {
	// Name is the name of the secret containing catalog values to apply,
	// e.g. app-catalog-secret.
	Name string `json:"name"`
	// Namespace is the namespace of the secret,
	// e.g. giantswarm.
	Namespace string `json:"namespace"`
}

// +k8s:openapi-gen=true
type AppCatalogSpecStorage struct {
	// Type indicates which repository type would be used for this AppCatalog.
	// e.g. helm
	Type string `json:"type"`
	// URL is the link to where this AppCatalog's repository is located
	// e.g. https://example.com/app-catalog/
	URL string `json:"URL"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AppCatalogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []AppCatalog `json:"items"`
}
