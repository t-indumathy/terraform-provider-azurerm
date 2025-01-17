package attacheddatanetwork

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttachedDataNetworkClient struct {
	Client *resourcemanager.Client
}

func NewAttachedDataNetworkClientWithBaseURI(api environments.Api) (*AttachedDataNetworkClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "attacheddatanetwork", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AttachedDataNetworkClient: %+v", err)
	}

	return &AttachedDataNetworkClient{
		Client: client,
	}, nil
}
