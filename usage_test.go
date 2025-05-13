// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package carbonaware_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/carbon-aware/scheduler-client-golang"
	"github.com/carbon-aware/scheduler-client-golang/internal/testutil"
	"github.com/carbon-aware/scheduler-client-golang/option"
)

func TestUsage(t *testing.T) {
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
	schedule, err := client.Schedule.New(context.TODO(), carbonaware.ScheduleNewParams{
		Duration: "PT1H",
		Windows: []carbonaware.ScheduleNewParamsWindow{{
			End:   time.Now(),
			Start: time.Now(),
		}},
		Zones: []carbonaware.CloudZoneParam{{
			Provider: carbonaware.CloudZoneProviderAws,
			Region:   carbonaware.CloudZoneRegionUsEast1,
		}},
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", schedule.Ideal)
}
