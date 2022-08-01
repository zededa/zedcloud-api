// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"github.com/spf13/cobra"
	zedcloudapi "github.com/zededa/zedcloud-api"
	"github.com/zededa/zedcloud-api/swagger_models"
)

func getAppInstConfig(client *zedcloudapi.Client,
	name, id string) (*swagger_models.AppInstance, error) {
	fmt.Printf("\nSending GET AppInst Config Request\n")
	client.XRequestIdPrefix = "zedcloudapi-client-get-appinst"
	rspData := &swagger_models.AppInstance{}
	_, err := client.GetObj("apps/instances", name, id, false, rspData)
	if err != nil {
		fmt.Printf("Failed to get AppInst. err: %+v", err)
		return nil, err
	}
	return rspData, nil
}

func deleteAppInst(client *zedcloudapi.Client, appID string) {
	// Warning or errors can be collected in a slice type
	resp, _, _ := client.DeleteObj("apps/instances", appID)
	if resp.StatusCode == 200 {
		fmt.Printf("[INFO] App Instance Delete Successful.")
		return
	}
	fmt.Printf("***[ERROR] App Instance Delete Failed. Status: %s\n resp: %++v",
		resp.Status, resp)
}

func createAppInst(client *zedcloudapi.Client, params *CmdParams) {
	title := "zedcloud_client test AppInst"
	activate := "true"
	appType := swagger_models.AppTypeAPPTYPEVM
	var cpus int64 = 1
	var memory int64 = 512
	var vnc bool = true
	hvMode := swagger_models.HvModeHVHVM

	vmInfo := swagger_models.VM{
		Cpus:   &cpus,
		Memory: &memory,
		Mode:   &hvMode,
		Vnc:    &vnc,
	}
	cfg := &swagger_models.AppInstance{
		Name:        &params.name,
		Title:       &title,
		Description: title,
		DeviceID:    &params.edgeNodeId,
		Activate:    &activate,
		AppID:       &params.edgeAppId,
		AppType:     &appType,
		Vminfo:      &vmInfo,
		Interfaces: []*swagger_models.AppInterface{
			&swagger_models.AppInterface{
				Intfname:    &params.intfName,
				Netinstname: &params.netInstName,
			},
		},
	}
	rspData := &swagger_models.ZsrvResponse{}
	client.XRequestIdPrefix = "zedcloudapi-client-create-appinst"
	resp, err := client.SendReq("POST", "apps/instances", cfg, rspData)
	if err != nil {
		fmt.Printf("\n***[ERROR] Failed to create app instance. err: %s",
			err.Error())
		return
	}
	if resp.StatusCode != 200 {
		fmt.Printf("\n***[ERROR] Failed to create app instance. rsp.Status: %s"+
			"\nresp: %+v", resp.Status, resp)
		return
	}
	fmt.Printf("App Instance %s (ID: %s) Successfully created\n",
		rspData.ObjectName, rspData.ObjectID)
	return
}

func appInstCmdHndlr(cmd *cobra.Command) {
	client := getClient(cmd)
	if client == nil {
		return
	}
	params := getParams(cmd)
	switch params.op {
	case "get":
		cfg, _ := getAppInstConfig(client, params.name, "")
		if cfg == nil {
			fmt.Printf("\nApp Instance %s Not Found.\n", params.name)
		} else {
			fmt.Printf("Found AppInst. cfg: %+v\n", *cfg)
		}
	case "create":
		fmt.Printf("\nCreating App Instance %s\n", params.name)
		createAppInst(client, params)
	case "delete":
		cfg, _ := getAppInstConfig(client, params.name, "")
		if cfg == nil {
			fmt.Printf("\nApp Instance %s Not Found.\n", params.name)
		} else {
			fmt.Printf("Found AppInst. cfg: %+v\n", *cfg)
			fmt.Printf("\nDeleting App Instance %s\n", params.name)
			deleteAppInst(client, cfg.ID)
		}
	}
}

var appInstCmd = &cobra.Command{
	Use:   "app-inst",
	Short: "App Instance related commands",
	Long:  "App Instance related commands",
	Run: func(cmd *cobra.Command, args []string) {
		appInstCmdHndlr(cmd)
	},
}
