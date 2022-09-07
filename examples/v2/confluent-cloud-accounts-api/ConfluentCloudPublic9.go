// Update resource in Confluent account returns "OK" response

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
	body := datadogV2.UpdateResourceRequest{
		Data: datadogV2.UpdateResource{
			Attributes: &datadogV2.UpdateResourceAttributes{
				ResourceType: datadog.PtrString("kafka"),
				Tags: []string{
					"myTag",
					"myTag2:myValue",
				},
			},
			Id:   "resource-id-123",
			Type: datadogV2.UPDATERESOURCETYPE_CONFLUENT_CLOUD_RESOURCES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewConfluentCloudAccountsAPIApi(apiClient)
	resp, r, err := api.ConfluentCloudPublic9(ctx, "account_id", "resource_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfluentCloudAccountsAPIApi.ConfluentCloudPublic9`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ConfluentCloudAccountsAPIApi.ConfluentCloudPublic9`:\n%s\n", responseContent)
}
