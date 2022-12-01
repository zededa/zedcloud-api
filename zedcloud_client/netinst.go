// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"

	"github.com/spf13/cobra"
	zedcloudapi "github.com/zededa/zedcloud-api"
	"github.com/zededa/zedcloud-api/swagger_models"
)

func getNetInstConfig(client *zedcloudapi.Client,
	name, id string) (*swagger_models.NetInstConfig, error) {
	fmt.Printf("\nSending GET NetInst Config Request\n")
	client.XRequestIdPrefix = "zedcloudapi-client-get-nwinst"
	rspData := &swagger_models.NetInstConfig{}
	_, err := client.GetObj("netinsts", name, id, false, rspData)
	if err != nil {
		return nil, err
	}
	return rspData, nil
}

func createNetInst(client *zedcloudapi.Client, name, edgeNodeId string) {
	title := "zedcloud_client test NetInst"
	kind := swagger_models.NetworkInstanceKindNETWORKINSTANCEKINDLOCAL
	dhcpType := swagger_models.NetworkInstanceDhcpTypeNETWORKINSTANCEDHCPTYPEV4
	cfg := &swagger_models.NetInstConfig{
		Name:          &name,
		Title:         &title,
		Description:   title,
		DeviceDefault: "false",
		DeviceID:      &edgeNodeId,
		Kind:          &kind,
		Type:          &dhcpType,
	}
	resp, rspData, err := client.CreateObj("netinsts", cfg)
	if err != nil || resp == nil {
		fmt.Printf("\n***[ERROR] Failed to create network instance. err: %s, resp: %+v",
			err.Error(), resp)
		return
	}
	if resp.StatusCode != 200 {
		fmt.Printf("\n***[ERROR] Failed to create network instance. rsp.Status: %s"+
			"\nresp: %+v", resp.Status, resp)
		return
	}
	fmt.Printf("Network Instance %s (ID: %s) Successfully created\n",
		rspData.ObjectName, rspData.ObjectID)
	return
}

func netInstCmdHndlr(cmd *cobra.Command) {
	client := getClient(cmd)
	if client == nil {
		return
	}
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		fmt.Printf("***ERROR: failed to get Network Instance name. err: %s ", err.Error())
		return
	}
	if name == "" {
		fmt.Printf("***[ERROR]: App Name must be specified")
		return
	}
	cfg, err := getNetInstConfig(client, name, "")
	if cfg == nil {
		edgeNodeId, err := cmd.Flags().GetString("edgenode-id")
		if err != nil {
			fmt.Printf("***ERROR: failed to get EdgeNode ID. err: %s ", err.Error())
			return
		}
		if edgeNodeId == "" {
			fmt.Printf("***[ERROR]: EdgeNodeID must be specified")
			return
		}
		createNetInst(client, name, edgeNodeId)
		fmt.Printf("\nNetwork Instance %s Not Found. Creating It.\n", name)
	}
	fmt.Printf("Found NetInst. cfg: %+v\n", *cfg)
}

var netInstCmd = &cobra.Command{
	Use:   "netinst",
	Short: "Network Instance related commands",
	Long:  "Network Instance related commands",
	Run: func(cmd *cobra.Command, args []string) {
		netInstCmdHndlr(cmd)
	},
}
