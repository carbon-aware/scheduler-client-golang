// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package carbonaware_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/carbon-aware/scheduler-client-golang"
	"github.com/carbon-aware/scheduler-client-golang/internal/testutil"
	"github.com/carbon-aware/scheduler-client-golang/option"
)

func TestHealthCheck(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := carbonaware.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Health.Check(context.TODO())
	if err != nil {
		var apierr *carbonaware.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
