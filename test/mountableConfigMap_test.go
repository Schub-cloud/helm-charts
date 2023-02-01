package deployment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMountableConfigMap(t *testing.T) {
	// Define test data
	values := map[string]interface{}{
		"config": map[string]interface{}{
			"enabled": true,
			"volumes": []map[string]interface{}{
				{
					"mount": map[string]interface{}{
						"name":      "test-config-map",
						"mountPath": "/etc/config",
					},
					"config": map[string]interface{}{
						"name": "test-config-map",
					},
				},
			},
		},
	}

	// Render the deployment template
	deployment, err := RenderDeployment(values)
	if err != nil {
		t.Fatalf("Failed to render deployment: %s", err)
	}

	// Check if the config map is properly mounted in the container
	assert.Contains(t, deployment.Spec.Template.Spec.Containers[0].VolumeMounts,
		VolumeMount{
			Name:      "test-config-map",
			MountPath: "/etc/config",
		})

	assert.Contains(t, deployment.Spec.Template.Spec.Volumes,
		Volume{
			Name: "test-config-map",
			VolumeSource: VolumeSource{
				ConfigMap: &ConfigMapVolumeSource{
					LocalObjectReference: LocalObjectReference{
						Name: "test-config-map",
					},
				},
			},
		})
}
