package v1alpha1

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/yaml"
)

func Test_GenerateChartYAML(t *testing.T) {
	testCases := []struct {
		category string
		name     string
		resource runtime.Object
	}{
		{
			category: "cr",
			name:     fmt.Sprintf("%s_%s_chart.yaml", group, version),
			resource: newChartExampleCR(),
		},
	}

	docs := filepath.Join(root, "..", "..", "docs")
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("case %d: generates %s successfully", i, tc.name), func(t *testing.T) {
			rendered, err := yaml.Marshal(tc.resource)
			if err != nil {
				t.Fatal(err)
			}
			directory := filepath.Join(docs, tc.category)
			path := filepath.Join(directory, tc.name)

			// We don't want a status in the docs YAML for the CR and CRD so that they work with `kubectl create -f <file>.yaml`.
			// This just strips off the top level `status:` and everything following.
			statusRegex := regexp.MustCompile(`(?ms)^status:.*$`)
			rendered = statusRegex.ReplaceAll(rendered, []byte(""))

			if *update {
				err := os.WriteFile(path, rendered, 0644) // nolint
				if err != nil {
					t.Fatal(err)
				}
			}
			goldenFile, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}

			if !bytes.Equal(rendered, goldenFile) {
				t.Fatalf("\n\n%s\n", cmp.Diff(string(goldenFile), string(rendered)))
			}
		})
	}
}

func newChartExampleCR() *Chart {
	cr := NewChartCR()

	cr.ObjectMeta = metav1.ObjectMeta{
		Name:      "prometheus",
		Namespace: "default",
		Labels: map[string]string{
			"chart-operator.giantswarm.io/version": "1.0.0",
		},
	}
	cr.Spec = ChartSpec{
		Name:      "prometheus",
		Namespace: "monitoring",
		Install: ChartSpecInstall{
			SkipCRDs: true,
			Timeout:  &metav1.Duration{Duration: 360 * time.Second},
		},
		Config: ChartSpecConfig{
			ConfigMap: ChartSpecConfigConfigMap{
				Name:      "f2def-chart-values",
				Namespace: "f2def",
			},
			Secret: ChartSpecConfigSecret{
				Name:      "f2def-chart-values",
				Namespace: "f2def",
			},
		},
		NamespaceConfig: ChartSpecNamespaceConfig{
			Annotations: map[string]string{
				"ownership.my-org.com/responsible": "my-team",
			},
		},
		Rollback: ChartSpecRollback{
			Timeout: &metav1.Duration{Duration: 420 * time.Second},
		},
		TarballURL: "prometheus-1.0.1.tgz",
		Uninstall: ChartSpecUninstall{
			Timeout: &metav1.Duration{Duration: 480 * time.Second},
		},
		Upgrade: ChartSpecUpgrade{
			Timeout: &metav1.Duration{Duration: 540 * time.Second},
		},
		Version: "1.0.1",
	}

	return cr
}
