// Get hourly usage for CI visibility returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewUsageMeteringApi(apiClient)
	resp, r, err := api.GetUsageCIApp(ctx, time.Date(2021, 11, 11, 11, 11, 11, 111000, time.UTC), *datadogV1.NewGetUsageCIAppOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageCIApp`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsageMeteringApi.GetUsageCIApp`:\n%s\n", responseContent)
}
