package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAutoReloadAnnotations(t *testing.T) {
	// Prepare test input data
	autoReloadConfig := true

	// Call the function that adds the annotations for automatic deployment restart on secret and configmap update
	annotations := addAutoReloadAnnotations(autoReloadConfig)

	// Verify that the expected annotations have been added
	assert.Equal(t, annotations["checksum/config"], "", "Annotation for checksum/config not added")
	assert.Equal(t, annotations["checksum/secrets"], "", "Annotation for checksum/secrets not added")
	assert.Equal(t, annotations["reload.strategy"], "rollingUpdate", "Annotation for reload.strategy not added")
}

func addAutoReloadAnnotations(autoReloadConfig bool) map[string]string {
	annotations := make(map[string]string)
	if autoReloadConfig {
		annotations["checksum/config"] = ""
		annotations["checksum/secrets"] = ""
		annotations["reload.strategy"] = "rollingUpdate"
	}
	return annotations
}
