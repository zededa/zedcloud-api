// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"github.com/spf13/cobra"
	zedcloudapi "github.com/zededa/zedcloud-api"
	"github.com/zededa/zedcloud-api/swagger_models"
)

func getEdgeAppConfig(client *zedcloudapi.Client,
	name, id string) (*swagger_models.AppConfig, error) {
	fmt.Printf("\nSending getEdgeAppConfig Request\n")
	rspData := &swagger_models.AppConfig{}
	_, err := client.GetObj("apps", name, id, false, rspData)
	if err != nil {
		return nil, err
	}
	return rspData, nil
}

func getEdgeAppConfigHndler(params *CmdParams) {
	cfg, err := getEdgeAppConfig(params.client, params.name, "")
	if cfg == nil {
		fmt.Printf("\n***Failed to get edgeApp config for %s. err: %s\n",
			params.name, err.Error())
		return
	}
	fmt.Printf("getEdgeAppConfig SUCCESS. name: %s, ID: %s\n",
		*cfg.Name, cfg.ID)
}

func createEdgeAppConfigHndler(params *CmdParams) {
	fmt.Printf("createEdgeAppConfigHndler")
}
func updateEdgeAppConfigHndler(params *CmdParams) {
	fmt.Printf("updateEdgeAppConfigHndler")
}

func deleteEdgeAppConfigHndler(params *CmdParams) {
	fmt.Printf("deleteEdgeAppConfigHndler")
}

var edgeAppCmd = &CmdInfo{
	hndlrs: map[string]HndlrFunc{
		"create": createEdgeAppConfigHndler,
		"get":    getEdgeAppConfigHndler,
		"update": updateEdgeAppConfigHndler,
		"delete": deleteEdgeAppConfigHndler,
	},
	cobraCmd: &cobra.Command{
		Use:   "edge-app",
		Short: "Edge App related commands",
		Long:  "Edge App related commands",
	},
}
