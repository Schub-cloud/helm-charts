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
// Test case for automatic deployment restart on secret and configmap update (test if given an auto reload config is given the right annotations are added to deployment template)
		{
			name: "autoReloadEnabled",
			values: map[string]string{
				"autoReload.enabled": "true",
			},
			golden: "autoReloadEnabled.yaml",
		},
		{
			name: "autoReloadDisabled",
			values: map[string]string{
				"autoReload.enabled": "false",
			},
			golden: "autoReloadDisabled.yaml",
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
