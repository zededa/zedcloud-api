// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
	zedcloudapi "github.com/zededa/zedcloud-api"
	"github.com/zededa/zedcloud-api/swagger_models"
)

func setDeviceTitle(client *zedcloudapi.Client, devCfg *swagger_models.DeviceConfig) {
	title := "Device Title Updated " + time.Now().String()
	devCfg.Title = &title
	fmt.Printf("\n\nUpdating Device Title to %s\n", *devCfg.Title)
	client.SendReq("PUT", "devices/id/cf1c0382-81b4-4d9d-8166-ed4db05ec489",
		devCfg, nil)
}

func getEdgeNodeConfig(client *zedcloudapi.Client, name, id string) *swagger_models.DeviceConfig {
	fmt.Printf("\nSending GetDevice Request\n")
	rspData := &swagger_models.DeviceConfig{}
	_, err := client.GetObj("devices", name, id, false, rspData)
	if err != nil {
		return rspData
	}
	return nil
}

func edgeNodeCmdHndlr(cmd *cobra.Command) {
	client := getClient(cmd)
	if client == nil {
		return
	}
	name, _ := cmd.Flags().GetString("name")
	devCfg := getEdgeNodeConfig(client, name, "")
	if devCfg == nil {
		fmt.Printf("\n***Failed to get edge-node config for %s\n", name)
		return
	}
	setDeviceTitle(client, devCfg)
}

var edgeNodeCmd = &cobra.Command{
	Use:   "edge-node",
	Short: "Edge Node related commands",
	Long:  "Edge Node related commands",
	Run: func(cmd *cobra.Command, args []string) {
		edgeNodeCmdHndlr(cmd)
	},
}
