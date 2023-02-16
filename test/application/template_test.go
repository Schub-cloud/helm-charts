package application

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"gotest.tools/v3/golden"

	"github.com/gruntwork-io/terratest/modules/helm"
)

func TestTemplateRender(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../../application")
	releaseName := "application"
	require.NoError(t, err)

	testCases := []struct {
		name            string
		values          map[string]string
		golden          string
		renderTemplates []string
	}{
		{
			name: "priorityClassLow",
			values: map[string]string{
				"priorityClass.enabled": "true",
				"priorityClass.name":    "low",
			},
			golden:          "priorityClass.yaml",
			renderTemplates: []string{"templates/deployment.yaml"},
		},
		{
			name: "priorityClassHigh",
			values: map[string]string{
				"priorityClass.enabled": "true",
				"priorityClass.name":    "high",
			},
			golden:          "priorityClassHigh.yaml",
			renderTemplates: []string{"templates/deployment.yaml"},
		},
		
// Test case of mountable config maps and secrets (test if when values provide a mountable config map it renders it mounted in the container of the deployment)
		{
			name: "mountConfigMap",
			values: map[string]string{
				"configMap.enabled": "true",
				"configMap.name":    "test-configmap",
			},
			golden:          "mountConfigMap.yaml",
			renderTemplates: []string{"templates/deployment.yaml"},
		},
		{
			name: "mountSecret",
			values: map[string]string{

				"secret.name":    "my-secret",
			},
			golden:          "mountSecret.yaml",
			renderTemplates: []string{"templates/deployment.yaml"},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(subT *testing.T) {
			subT.Parallel()

			options := &helm.Options{
				SetValues: testCase.values,
			}

			output := helm.RenderTemplate(t, options, helmChartPath, releaseName, testCase.renderTemplates)

			golden.Assert(t, output, testCase.golden)
		})
	}
}
