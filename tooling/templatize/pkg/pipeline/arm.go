package pipeline

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/go-logr/logr"
)

type armClient struct {
	creds          *azidentity.DefaultAzureCredential
	SubscriptionID string
	Region         string
}

func newArmClient(subscriptionID, region string) *armClient {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil
	}
	return &armClient{
		creds:          cred,
		SubscriptionID: subscriptionID,
		Region:         region,
	}
}

func (a *armClient) runArmStep(ctx context.Context, options *PipelineRunOptions, deploymentName string, rgName string, paramterFile string, input map[string]output) (output, error) {
	logger := logr.FromContextOrDiscard(ctx)

	// Transform Bicep to ARM
	deploymentProperties, err := transformBicepToARM(ctx, paramterFile, options.Vars)
	if err != nil {
		return nil, fmt.Errorf("failed to transform Bicep to ARM: %w", err)
	}

	// Create the deployment
	deployment := armresources.Deployment{
		Properties: deploymentProperties,
	}

	// Ensure resourcegroup exists
	err = a.ensureResourceGroupExists(ctx, rgName)
	if err != nil {
		return nil, fmt.Errorf("failed to ensure resource group exists: %w", err)
	}

	// TODO handle dry-run

	// Run deployment
	client, err := armresources.NewDeploymentsClient(a.SubscriptionID, a.creds, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create deployments client: %w", err)
	}

	poller, err := client.BeginCreateOrUpdate(ctx, rgName, deploymentName, deployment, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create deployment: %w", err)
	}
	logger.Info("Deployment started", "deployment", deploymentName)

	// Wait for completion
	resp, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to wait for deployment completion: %w", err)
	}
	logger.Info("Deployment finished successfully", "deployment", deploymentName, "responseId", *resp.ID)

	if resp.Properties.Outputs != nil {
		if outputMap, ok := resp.Properties.Outputs.(map[string]any); ok {
			returnMap := armOutput{}
			for k, v := range outputMap {
				returnMap[k] = v
			}
			return returnMap, nil
		}
	}
	return nil, nil
}

func (a *armClient) ensureResourceGroupExists(ctx context.Context, rgName string) error {
	// Create a new Azure identity client
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return fmt.Errorf("failed to obtain a credential: %w", err)
	}

	// Create a new ARM client
	client, err := armresources.NewResourceGroupsClient(a.SubscriptionID, cred, nil)
	if err != nil {
		return fmt.Errorf("failed to create ARM client: %w", err)
	}

	// Check if the resource group exists
	// todo fill tags properly
	tags := map[string]*string{
		"persist": to.Ptr("true"),
	}
	_, err = client.Get(ctx, rgName, nil)
	if err != nil {
		// Create the resource group
		resourceGroup := armresources.ResourceGroup{
			Location: to.Ptr(a.Region),
			Tags:     tags,
		}
		_, err = client.CreateOrUpdate(ctx, rgName, resourceGroup, nil)
		if err != nil {
			return fmt.Errorf("failed to create resource group: %w", err)
		}
	} else {
		patchResourceGroup := armresources.ResourceGroupPatchable{
			Tags: tags,
		}
		_, err = client.Update(ctx, rgName, patchResourceGroup, nil)
		if err != nil {
			return fmt.Errorf("failed to update resource group: %w", err)
		}
	}
	return nil
}
