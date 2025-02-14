//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package test_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

// x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Services_Get.json
func ExampleServicesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := test.NewServicesClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<service-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ServicesClientGetResult)
}

// x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Services_CreateOrUpdate.json
func ExampleServicesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := test.NewServicesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<service-name>",
		test.ServiceResource{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"key1": to.StringPtr("value1"),
			},
			Properties: &test.ClusterResourceProperties{},
			SKU: &test.SKU{
				Name: to.StringPtr("<name>"),
				Tier: to.StringPtr("<tier>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ServicesClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Services_Delete.json
func ExampleServicesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := test.NewServicesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<service-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Services_Update.json
func ExampleServicesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := test.NewServicesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<service-name>",
		test.ServiceResource{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"key1": to.StringPtr("value1"),
			},
			Properties: &test.ClusterResourceProperties{},
			SKU: &test.SKU{
				Name: to.StringPtr("<name>"),
				Tier: to.StringPtr("<tier>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ServicesClientUpdateResult)
}

// x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Services_ListTestKeys.json
func ExampleServicesClient_ListTestKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := test.NewServicesClient("<subscription-id>", cred, nil)
	res, err := client.ListTestKeys(ctx,
		"<resource-group-name>",
		"<service-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ServicesClientListTestKeysResult)
}

// x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Services_RegenerateTestKey.json
func ExampleServicesClient_RegenerateTestKey() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := test.NewServicesClient("<subscription-id>", cred, nil)
	res, err := client.RegenerateTestKey(ctx,
		"<resource-group-name>",
		"<service-name>",
		test.RegenerateTestKeyRequestPayload{
			KeyType: test.TestKeyTypePrimary.ToPtr(),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ServicesClientRegenerateTestKeyResult)
}

// x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Services_DisableTestEndpoint.json
func ExampleServicesClient_DisableTestEndpoint() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := test.NewServicesClient("<subscription-id>", cred, nil)
	_, err = client.DisableTestEndpoint(ctx,
		"<resource-group-name>",
		"<service-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Services_EnableTestEndpoint.json
func ExampleServicesClient_EnableTestEndpoint() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := test.NewServicesClient("<subscription-id>", cred, nil)
	res, err := client.EnableTestEndpoint(ctx,
		"<resource-group-name>",
		"<service-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ServicesClientEnableTestEndpointResult)
}

// x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Services_CheckNameAvailability.json
func ExampleServicesClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := test.NewServicesClient("<subscription-id>", cred, nil)
	res, err := client.CheckNameAvailability(ctx,
		"<location>",
		test.NameAvailabilityParameters{
			Name: to.StringPtr("<name>"),
			Type: to.StringPtr("<type>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ServicesClientCheckNameAvailabilityResult)
}

// x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Services_ListBySubscription.json
func ExampleServicesClient_ListBySubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := test.NewServicesClient("<subscription-id>", cred, nil)
	pager := client.ListBySubscription(nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Services_List.json
func ExampleServicesClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := test.NewServicesClient("<subscription-id>", cred, nil)
	pager := client.List("<resource-group-name>",
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}
