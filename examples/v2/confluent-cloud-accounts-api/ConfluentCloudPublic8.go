// Add resource to Confluent account returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.ConfluentCloudResourceRequest{
		Data: datadogV2.ConfluentCloudResource{
			Id:           "resource_id_abc123",
			ResourceType: "kafka",
			Tags: []string{
				"myTag",
				"myTag2:myValue",
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewConfluentCloudAccountsAPIApi(apiClient)
	resp, r, err := api.ConfluentCloudPublic8(ctx, "account_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfluentCloudAccountsAPIApi.ConfluentCloudPublic8`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ConfluentCloudAccountsAPIApi.ConfluentCloudPublic8`:\n%s\n", responseContent)
}
