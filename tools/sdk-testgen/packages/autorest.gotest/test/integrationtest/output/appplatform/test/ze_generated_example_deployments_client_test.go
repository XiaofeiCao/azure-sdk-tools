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

// x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Deployments_Get.json
func ExampleDeploymentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := test.NewDeploymentsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<app-name>",
		"<deployment-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DeploymentsClientGetResult)
}

// x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Deployments_CreateOrUpdate.json
func ExampleDeploymentsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := test.NewDeploymentsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<app-name>",
		"<deployment-name>",
		test.DeploymentResource{
			Properties: &test.DeploymentResourceProperties{
				DeploymentSettings: &test.DeploymentSettings{
					CPU: to.Int32Ptr(1),
					EnvironmentVariables: map[string]*string{
						"env": to.StringPtr("test"),
					},
					JvmOptions:     to.StringPtr("<jvm-options>"),
					MemoryInGB:     to.Int32Ptr(3),
					RuntimeVersion: test.RuntimeVersionJava8.ToPtr(),
				},
				Source: &test.UserSourceInfo{
					Type:             test.UserSourceTypeSource.ToPtr(),
					ArtifactSelector: to.StringPtr("<artifact-selector>"),
					RelativePath:     to.StringPtr("<relative-path>"),
					Version:          to.StringPtr("<version>"),
				},
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
	log.Printf("Response result: %#v\n", res.DeploymentsClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Deployments_Delete.json
func ExampleDeploymentsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := test.NewDeploymentsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<app-name>",
		"<deployment-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Deployments_Update.json
func ExampleDeploymentsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := test.NewDeploymentsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<app-name>",
		"<deployment-name>",
		test.DeploymentResource{
			Properties: &test.DeploymentResourceProperties{
				Source: &test.UserSourceInfo{
					Type:             test.UserSourceTypeSource.ToPtr(),
					ArtifactSelector: to.StringPtr("<artifact-selector>"),
					RelativePath:     to.StringPtr("<relative-path>"),
					Version:          to.StringPtr("<version>"),
				},
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
	log.Printf("Response result: %#v\n", res.DeploymentsClientUpdateResult)
}

// x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Deployments_List.json
func ExampleDeploymentsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := test.NewDeploymentsClient("<subscription-id>", cred, nil)
	pager := client.List("<resource-group-name>",
		"<service-name>",
		"<app-name>",
		&test.DeploymentsClientListOptions{Version: []string{}})
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

// x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Deployments_ListForCluster.json
func ExampleDeploymentsClient_ListForCluster() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := test.NewDeploymentsClient("<subscription-id>", cred, nil)
	pager := client.ListForCluster("<resource-group-name>",
		"<service-name>",
		&test.DeploymentsClientListForClusterOptions{Version: []string{}})
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

// x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Deployments_Start.json
func ExampleDeploymentsClient_BeginStart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := test.NewDeploymentsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginStart(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<app-name>",
		"<deployment-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Deployments_Stop.json
func ExampleDeploymentsClient_BeginStop() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := test.NewDeploymentsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginStop(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<app-name>",
		"<deployment-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2020-11-01-preview/examples/Deployments_Restart.json
func ExampleDeploymentsClient_BeginRestart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := test.NewDeploymentsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginRestart(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<app-name>",
		"<deployment-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
