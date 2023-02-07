package services

import (
	"context"
	"os"
	"testing"

	mlp "github.com/gojek/mlp/api/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2/google"

	"github.com/caraml-dev/timber/dataset-service/services/mocks"
)

func TestNewMLPClient(t *testing.T) {
	reset := testSetupEnvForGoogleCredentials(t)
	defer reset()

	// Create test Google client
	gc, err := google.DefaultClient(context.Background(), "https://www.googleapis.com/auth/userinfo.email")
	require.NoError(t, err)
	// Create expected MLP config
	cfg := mlp.NewConfiguration()
	cfg.BasePath = "base-path"
	cfg.HTTPClient = gc

	// Test
	resultClient := newMLPClient(gc, "base-path")
	require.NotNil(t, resultClient)
	assert.Equal(t, mlp.NewAPIClient(cfg), resultClient.api)
}

func TestMLPServiceGetProject(t *testing.T) {
	mlpSvc := &mocks.MLPService{}
	expectedProject := &mlp.Project{Id: 1}

	projectID := int64(1)
	mlpSvc.On("GetProject", projectID).Return(expectedProject, nil)

	project, err := mlpSvc.GetProject(projectID)
	assert.NoError(t, err)
	assert.Equal(t, project, expectedProject)
}

// testSetupEnvForGoogleCredentials creates a temporary file containing dummy service account JSON
// then set the environment variable GOOGLE_APPLICATION_CREDENTIALS to point to the the file.
// This is useful for tests that assume Google Cloud Client libraries can automatically find
// the service account credentials in any environment.
// At the end of the test, the returned function can be called to perform cleanup.
func testSetupEnvForGoogleCredentials(t *testing.T) (reset func()) {
	serviceAccountKey := []byte(`{
		"type": "service_account",
		"project_id": "foo",
		"private_key_id": "bar",
		"private_key": "baz",
		"client_email": "foo@example.com",
		"client_id": "bar_client_id",
		"auth_uri": "https://oauth2.googleapis.com/auth",
		"token_uri": "https://oauth2.googleapis.com/token"
	}`)

	file, err := os.CreateTemp("", "dummy-service-account")
	if err != nil {
		t.Fatal(err)
	}

	err = os.WriteFile(file.Name(), serviceAccountKey, 0644)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", file.Name())
	if err != nil {
		t.Fatal(err)
	}

	return func() {
		err := os.Unsetenv("GOOGLE_APPLICATION_CREDENTIALS")
		if err != nil {
			t.Log("Cleanup failed", err)
		}
		err = os.Remove(file.Name())
		if err != nil {
			t.Log("Cleanup failed", err)
		}
	}
}
