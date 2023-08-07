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
		valuesFiles     []string
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
		{
			name: "autoReloadEnabled",
			values: map[string]string{
				"reloader.enabled": "true",
			},
			golden:          "autoReloadEnabled.yaml",
			renderTemplates: []string{"templates/deployment.yaml"},
		},
		{
			name: "autoReloadDisabled",
			values: map[string]string{
				"reloader.enabled": "false",
			},
			golden:          "autoReloadDisabled.yaml",
			renderTemplates: []string{"templates/deployment.yaml"},
		},
		{
			name: "envVarsWithKeyValuePairs",
			values: map[string]string{
				"env.KEY.value": "my-valye",
			},
			golden:          "envVarsWithKeyValuePairs.yaml",
			renderTemplates: []string{"templates/deployment.yaml"},
		},
		{
			name: "envVarsWithKeyValuePairsAndValueFrom",
			values: map[string]string{
				"env.KEY.valueFrom.configMapKeyRef.name": "confit-map-name",
				"env.KEY.valueFrom.configMapKeyRef.key":  "confit-map-key",
			},
			golden:          "envVarsWithKeyValuePairsAndValueFrom.yaml",
			renderTemplates: []string{"templates/deployment.yaml"},
		},
		{
			name: "createPvc",
			values: map[string]string{
				"pvcs.my-pvc.storageClassName":           "manual",
				"pvcs.my-pvc.resources.requests.storage": "3Gi",
			},
			golden:          "createPvc.yaml",
			renderTemplates: []string{"templates/pvc.yaml"},
		},
		{
			name:        "createAndMountPVC",
			golden:      "createAndMountPVC.yaml",
			valuesFiles: []string{"testdata/valueFiles/createAndMountPVC.yaml"},
			renderTemplates: []string{
				"templates/pvc.yaml",
				"templates/deployment.yaml",
			},
		},
		{
			name:        "createAndMountPVCWithStaticPV",
			golden:      "createAndMountPVCWithStaticPV.yaml",
			valuesFiles: []string{"testdata/valueFiles/createAndMountPVCWithStaticPV.yaml"},
			renderTemplates: []string{
				"templates/pvc.yaml",
				"templates/pvs.yaml",
				"templates/deployment.yaml",
			},
		},
		{
			name:        "renderHTTPRoute",
			golden:      "renderHTTPRoute.yaml",
			valuesFiles: []string{"testdata/valueFiles/renderHTTPRoute.yaml"},
			renderTemplates: []string{
				"templates/http-route.yaml",
			},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(subT *testing.T) {
			subT.Parallel()

			options := &helm.Options{
				SetValues:   testCase.values,
				ValuesFiles: testCase.valuesFiles,
			}

			output := helm.RenderTemplate(t, options, helmChartPath, releaseName, testCase.renderTemplates)

			golden.Assert(t, output, testCase.golden)
		})
	}
}
