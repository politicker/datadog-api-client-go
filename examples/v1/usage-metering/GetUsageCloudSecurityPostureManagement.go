// Get hourly usage for CSPM returns "OK" response

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
	resp, r, err := api.GetUsageCloudSecurityPostureManagement(ctx, time.Now().AddDate(0, 0, -3), *datadogV1.NewGetUsageCloudSecurityPostureManagementOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageMeteringApi.GetUsageCloudSecurityPostureManagement`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `UsageMeteringApi.GetUsageCloudSecurityPostureManagement`:\n%s\n", responseContent)
}
