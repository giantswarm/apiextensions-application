package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/k8smetadata/pkg/annotation"
)

const (
	kindApp              = "App"
	appDocumentationLink = "https://docs.giantswarm.io/ui-api/management-api/crd/apps.application.giantswarm.io/"

	// NOTE: These should match the kubebuilder annotations set on AppExtraConfig
	ConfigPriorityDistance = 50
	ConfigPriorityCatalog  = 0
	ConfigPriorityDefault  = ConfigPriorityCatalog + ConfigPriorityDistance/2 //nolint
	ConfigPriorityCluster  = ConfigPriorityCatalog + ConfigPriorityDistance
	ConfigPriorityUser     = ConfigPriorityCluster + ConfigPriorityDistance
	ConfigPriorityMaximum  = ConfigPriorityUser + ConfigPriorityDistance //nolint
)

func NewAppTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       kindApp,
	}
}

// NewAppCR returns an App Custom Resource.
func NewAppCR() *App {
	return &App{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				annotation.Docs: appDocumentationLink,
			},
		},
		TypeMeta: NewAppTypeMeta(),
	}
}

// +kubebuilder:printcolumn:name="Desired Version",type=string,priority=1,JSONPath=`.spec.version`,description="Desired version of the app"
// +kubebuilder:printcolumn:name="Installed Version",type=string,JSONPath=`.status.version`,description="Installed version of the app"
// +kubebuilder:printcolumn:name="Created At",type=date,JSONPath=`.metadata.creationTimestamp`,description="Time of app creation"
// +kubebuilder:printcolumn:name="Last Deployed",type=date,JSONPath=`.status.release.lastDeployed`,description="Time since last deployment"
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.status.release.status`,description="Deployment status of the app"
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=common;giantswarm
// +k8s:openapi-gen=true
// App represents a managed app which a user intended to install. It is reconciled by app-operator.
type App struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppSpec `json:"spec"`
	// +kubebuilder:validation:Optional
	// Status Spec part of the App resource.
	// Initially, it would be left as empty until the operator successfully reconciles the helm release.
	Status AppStatus `json:"status,omitempty"`
}

// +k8s:openapi-gen=true
type AppSpec struct {
	// Catalog is the name of the app catalog this app belongs to.
	// e.g. giantswarm
	Catalog string `json:"catalog"`
	// +kubebuilder:validation:Optional
	// +nullable
	// CatalogNamespace is the namespace of the Catalog CR this app belongs to.
	// e.g. giantswarm
	CatalogNamespace string `json:"catalogNamespace,omitempty"`
	// +kubebuilder:validation:Optional
	// +nullable
	// Config is the config to be applied when the app is deployed.
	Config AppSpecConfig `json:"config,omitempty"`
	// +kubebuilder:validation:Optional
	// +nullable
	// +listType:=map
	// +listMapKey:=kind
	// +listMapKey:=name
	// +listMapKey:=namespace
	// ExtraConfigs is a list of configurations to merge together based on the priority and order in the list.
	// See: https://github.com/giantswarm/rfc/tree/main/multi-layer-app-config#enhancing-app-cr
	ExtraConfigs []AppExtraConfig `json:"extraConfigs,omitempty"`
	// +kubebuilder:validation:Optional
	// +nullable
	// Install is the config used when installing the app.
	Install AppSpecInstall `json:"install,omitempty"`
	// KubeConfig is the kubeconfig to connect to the cluster when deploying
	// the app.
	KubeConfig AppSpecKubeConfig `json:"kubeConfig"`
	// Name is the name of the app to be deployed.
	// e.g. kubernetes-prometheus
	Name string `json:"name"`
	// Namespace is the target namespace where the app should be deployed
	// e.g. monitoring, it cannot be changed.
	Namespace string `json:"namespace"`
	// +kubebuilder:validation:Optional
	// +nullable
	// NamespaceConfig is the namespace config to be applied to the target namespace when the app is deployed.
	NamespaceConfig AppSpecNamespaceConfig `json:"namespaceConfig,omitempty"`
	// +kubebuilder:validation:Optional
	// +nullable
	// Rollback is the config used when rolling back the app.
	Rollback AppSpecRollback `json:"rollback,omitempty"`
	// +kubebuilder:validation:Optional
	// +nullable
	// Upgrade is the config used when upgrading the app.
	Upgrade AppSpecUpgrade `json:"upgrade,omitempty"`
	// +kubebuilder:validation:Optional
	// +nullable
	// Uninstall is the config used when uninstalling the app.
	Uninstall AppSpecUninstall `json:"uninstall,omitempty"`
	// +kubebuilder:validation:Optional
	// +nullable
	// UserConfig is the user config to be applied when the app is deployed.
	UserConfig AppSpecUserConfig `json:"userConfig,omitempty"`
	// Version is the version of the app that should be deployed.
	// e.g. 1.0.0
	Version string `json:"version"`
}

// +k8s:openapi-gen=true
type AppSpecConfig struct {
	// +kubebuilder:validation:Optional
	// +nullable
	// ConfigMap references a config map containing values that should be
	// applied to the app.
	ConfigMap AppSpecConfigConfigMap `json:"configMap,omitempty"`
	// +kubebuilder:validation:Optional
	// +nullable
	// Secret references a secret containing secret values that should be
	// applied to the app.
	Secret AppSpecConfigSecret `json:"secret,omitempty"`
}

// +k8s:openapi-gen=true
type AppExtraConfig struct {
	// +optional
	// +kubebuilder:validation:Enum=configMap;secret
	// +kubebuilder:default:=configMap
	// Kind of configuration to look up that should be applied to the app when deployed.
	Kind string `json:"kind"`
	// Name of the resource of the given kind to look up.
	Name string `json:"name"`
	// Namespace where the resource with the given name and kind to look up is located.
	Namespace string `json:"namespace"`
	// +optional
	// +kubebuilder:validation:Minimum:=1
	// +kubebuilder:validation:Maximum:=150
	// +kubebuilder:default:=25
	// Priority is used to indicate at which stage the extra configuration should be merged.
	// See: https://github.com/giantswarm/rfc/tree/main/multi-layer-app-config#enhancing-app-cr
	Priority int `json:"priority"`
}

// +k8s:openapi-gen=true
type AppSpecInstall struct {
	// +kubebuilder:validation:Optional
	// +nullable
	// SkipCRDs when true decides that CRDs which are supplied with the chart are not installed. Default: false.
	SkipCRDs bool `json:"skipCRDs,omitempty"`
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern="^([0-9]+(\\.[0-9]+)?(ms|s|m))+$"
	// +optional
	// Timeout for the Helm install. When not set the default timeout of 5 minutes is being enforced.
	Timeout *metav1.Duration `json:"timeout,omitempty"`
}

// +k8s:openapi-gen=true
type AppSpecRollback struct {
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern="^([0-9]+(\\.[0-9]+)?(ms|s|m))+$"
	// +optional
	// Timeout for the Helm rollback. When not set the default timeout of 5 minutes is being enforced.
	Timeout *metav1.Duration `json:"timeout,omitempty"`
}

// +k8s:openapi-gen=true
type AppSpecUninstall struct {
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern="^([0-9]+(\\.[0-9]+)?(ms|s|m))+$"
	// +optional
	// Timeout for the Helm uninstall. When not set the default timeout of 5 minutes is being enforced.
	Timeout *metav1.Duration `json:"timeout,omitempty"`
}

// +k8s:openapi-gen=true
type AppSpecUpgrade struct {
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern="^([0-9]+(\\.[0-9]+)?(ms|s|m))+$"
	// +optional
	// Timeout for the Helm upgrade. When not set the default timeout of 5 minutes is being enforced.
	Timeout *metav1.Duration `json:"timeout,omitempty"`
}

// +k8s:openapi-gen=true
type AppSpecConfigConfigMap struct {
	// Name is the name of the config map containing app values to apply,
	// e.g. prometheus-values.
	Name string `json:"name" `
	// Namespace is the namespace of the values config map,
	// e.g. monitoring.
	Namespace string `json:"namespace"`
}

type AppSpecConfigSecret struct {
	// Name is the name of the secret containing app values to apply,
	// e.g. prometheus-secret.
	Name string `json:"name"`
	// Namespace is the namespace of the secret,
	// e.g. kube-system.
	Namespace string `json:"namespace"`
}

// +k8s:openapi-gen=true
type AppSpecKubeConfig struct {
	// InCluster is a flag for whether to use InCluster credentials. When true the
	// context name and secret should not be set.
	InCluster bool `json:"inCluster"`
	// +kubebuilder:validation:Optional
	// +nullable
	// Deprecated: this field is no longer used.
	Context AppSpecKubeConfigContext `json:"context,omitempty"`
	// +kubebuilder:validation:Optional
	// +nullable
	// Secret references a secret containing the kubconfig.
	Secret AppSpecKubeConfigSecret `json:"secret,omitempty"`
}

// +k8s:openapi-gen=true
type AppSpecKubeConfigContext struct {
	// Name is the name of the kubeconfig context
	// e.g. giantswarm-12345.
	Name string `json:"name"`
}

// +k8s:openapi-gen=true
type AppSpecKubeConfigSecret struct {
	// Name is the name of the secret containing the kubeconfig,
	// e.g. app-operator-kubeconfig.
	Name string `json:"name"`
	// Namespace is the namespace of the secret containing the kubeconfig,
	// e.g. giantswarm.
	Namespace string `json:"namespace"`
}

// +k8s:openapi-gen=true
type AppSpecNamespaceConfig struct {
	// +kubebuilder:validation:Optional
	// +nullable
	// Annotations is a string map of annotations to apply to the target namespace.
	Annotations map[string]string `json:"annotations,omitempty"`
	// +kubebuilder:validation:Optional
	// +nullable
	// Labels is a string map of labels to apply to the target namespace.
	Labels map[string]string `json:"labels,omitempty"`
}

// +k8s:openapi-gen=true
type AppSpecUserConfig struct {
	// +kubebuilder:validation:Optional
	// +nullable
	// ConfigMap references a config map containing user values that should be
	// applied to the app.
	ConfigMap AppSpecUserConfigConfigMap `json:"configMap,omitempty"`
	// +kubebuilder:validation:Optional
	// +nullable
	// Secret references a secret containing user secret values that should be
	// applied to the app.
	Secret AppSpecUserConfigSecret `json:"secret,omitempty"`
}

// +k8s:openapi-gen=true
type AppSpecUserConfigConfigMap struct {
	// Name is the name of the config map containing user values to apply,
	// e.g. prometheus-user-values.
	Name string `json:"name"`
	// Namespace is the namespace of the user values config map on the management cluster,
	// e.g. 123ab.
	Namespace string `json:"namespace"`
}

// +k8s:openapi-gen=true
type AppSpecUserConfigSecret struct {
	// Name is the name of the secret containing user values to apply,
	// e.g. prometheus-user-secret.
	Name string `json:"name"`
	// Namespace is the namespace of the secret,
	// e.g. kube-system.
	Namespace string `json:"namespace"`
}

// +k8s:openapi-gen=true
type AppStatus struct {
	// AppVersion is the value of the AppVersion field in the Chart.yaml of the
	// deployed app. This is an optional field with the version of the
	// component being deployed.
	// e.g. 0.21.0.
	// https://helm.sh/docs/topics/charts/#the-chartyaml-file
	AppVersion string `json:"appVersion"`
	// Release is the status of the Helm release for the deployed app.
	Release AppStatusRelease `json:"release"`
	// Version is the value of the Version field in the Chart.yaml of the
	// deployed app.
	// e.g. 1.0.0.
	Version string `json:"version"`
}

// +k8s:openapi-gen=true
type AppStatusRelease struct {
	// +kubebuilder:validation:Optional
	// +nullable
	// LastDeployed is the time when the app was last deployed.
	LastDeployed metav1.Time `json:"lastDeployed,omitempty"`
	// Reason is the description of the last status of helm release when the app is
	// not installed successfully, e.g. deploy resource already exists.
	Reason string `json:"reason,omitempty"`
	// Status is the status of the deployed app,
	// e.g. DEPLOYED.
	Status string `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []App `json:"items"`
}
