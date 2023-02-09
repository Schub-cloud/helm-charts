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

// Unit test of customizable liveness and readiness probes (test if the right liveness and readiness probes are rendered)
func TestRightLivenessReadinessProbes(t *testing.T) {
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
			name: "defaultProbes",
			values: map[string]string{
				"livenessProbe.enabled":  "false",
				"readinessProbe.enabled": "false",
			},
			golden:          "defaultProbes.yaml",
			renderTemplates: []string{"templates/deployment.yaml"},
		},
		{
			name: "customProbes",
			values: map[string]string{
				"livenessProbe.enabled":  "true",
				"livenessProbe.path":     "/health",
				"livenessProbe.port":     "8080",
				"livenessProbe.interval": "10s",
				"livenessProbe.timeout":  "5s",

				"readinessProbe.enabled": "true",
				"readinessProbe.path":    "/ready",
				"readinessProbe.port":    "8080",
				"readinessProbe.interval": "5s",
				"readinessProbe.timeout":  "2s",
			},
			golden:          "customProbes.yaml",
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


// Unit test of customizable liveness and readiness probes (test if the right liveness probes are rendered)
func TestRightLivenessProbes(t *testing.T) {
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
			name: "LivenessProbeEnabled",
			values: map[string]string{
				"livenessProbe.enabled":             "true",
				"livenessProbe.path":                "/health",
				"livenessProbe.port":                "8080",
				"livenessProbe.initialDelaySeconds": "5",
			},
			golden:          "livenessProbeEnabled.yaml",
			renderTemplates: []string{"templates/deployment.yaml"},
		},
		{
			name: "LivenessProbeDisabled",
			values: map[string]string{
				"livenessProbe.enabled":             "false",
				"livenessProbe.path":                "/health",
				"livenessProbe.port":                "8080",
				"livenessProbe.initialDelaySeconds": "5",
			},
			golden:          "livenessProbeDisabled.yaml",
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

// Unit test of customizable liveness and readiness probes (test if both liveness and readiness probes are rendered)
func TestBothLivenessAndReadinessProbes(t *testing.T) {
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
			name: "livenessProbe",
			values: map[string]string{
				"probes.liveness.enabled": "true",
				"probes.liveness.probe.httpGet.path": "/health",
				"probes.liveness.probe.httpGet.port": "8080",
				"probes.liveness.probe.initialDelaySeconds": "5",
			},
			golden:          "livenessProbe.yaml",
			renderTemplates: []string{"templates/deployment.yaml"},
		},
		{
			name: "readinessProbe",
			values: map[string]string{
				"probes.readiness.enabled": "true",
				"probes.readiness.probe.httpGet.path": "/readiness",
				"probes.readiness.probe.httpGet.port": "8080",
				"probes.readiness.probe.initialDelaySeconds": "5",
			},
			golden:          "readinessProbe.yaml",
			renderTemplates: []string{"templates/deployment.yaml"},
		},
		{
			name: "livenessAndReadinessProbes",
			values: map[string]string{
				"probes.liveness.enabled": "true",
				"probes.liveness.probe.httpGet.path": "/health",
				"probes.liveness.probe.httpGet.port": "8080",
				"probes.liveness.probe.initialDelaySeconds": "5",
				"probes.readiness.enabled": "true",
				"probes.readiness.probe.httpGet.path": "/readiness",
				"probes.readiness.probe.httpGet.port": "8080",
				"probes.readiness.probe.initialDelaySeconds": "5",
			},
			golden:          "livenessAndReadinessProbes.yaml",
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

