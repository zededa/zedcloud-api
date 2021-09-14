// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	zedcloudapi "github.com/zededa/zedcloud-api"
)

func getParam(cmd *cobra.Command, param string) (string, error) {
	val, err := cmd.Flags().GetString(param)
	if err != nil {
		fmt.Printf("***ERROR: failed to get %s. err: %s ", param, err.Error())
		return "", err
	}
	if val == "" {
		err = fmt.Errorf("***[ERROR]: %s must be specified", param)
		fmt.Printf(err.Error())
		return "", err
	}
	return val, nil
}

func getClient(cmd *cobra.Command) *zedcloudapi.Client {
	cloudUrl, _ := cmd.Flags().GetString("zedcloud-url")
	user, _ := cmd.Flags().GetString("user")
	password, _ := cmd.Flags().GetString("password")
	baseUrl := fmt.Sprintf("https://%s", cloudUrl)
	client, err := zedcloudapi.NewClient(baseUrl, "", user, password)
	if err != nil {
		return nil
	}
	return client
}

type CmdParams struct {
	client      *zedcloudapi.Client
	cmd         *cobra.Command
	op          string
	name        string
	id          string
	edgeNodeId  string
	edgeAppId   string
	appManifest string
	netInstName string
	intfName    string
}

func getParams(cmd *cobra.Command) *CmdParams {
	params := &CmdParams{}
	params.op, _ = cmd.Flags().GetString("op")
	params.name, _ = cmd.Flags().GetString("name")
	params.id, _ = cmd.Flags().GetString("id")
	params.edgeNodeId, _ = cmd.Flags().GetString("edgenode-id")
	params.edgeAppId, _ = cmd.Flags().GetString("edgeapp-id")
	params.appManifest, _ = cmd.Flags().GetString("app-manifest")
	params.intfName, _ = cmd.Flags().GetString("interface")
	params.netInstName, _ = cmd.Flags().GetString("network-instance")
	return params
}

func cmdHndlr(cmd *cobra.Command, args []string) {
	cmdInfo, ok := cmdMap[cmd.Use]
	if !ok {
		fmt.Printf("***[ERROR] - Command Handler for %s not found\n",
			cmd.Use)
		return
	}
	params := getParams(cmd)
	client := getClient(cmd)
	if client == nil {
		return
	}
	params.client = client
	hndlr, ok := cmdInfo.hndlrs[params.op]
	if hndlr == nil {
		fmt.Printf("***[ERROR] Handler for %s not defined for %s",
			params.op, cmd.Use)
		return
	}
	hndlr(params)
}

type HndlrFunc func(params *CmdParams)

type CmdInfo struct {
	hndlrs   map[string]HndlrFunc
	cobraCmd *cobra.Command
}

var cmdArray = []*CmdInfo{
	//edgeNodeCmd,
	edgeAppCmd,
	//netInstCmd,
	//appInstCmd,
}

var cmdMap map[string]*CmdInfo

var rootCmd = &cobra.Command{
	Use:   "zedcloud-client",
	Short: "zedcloud-client is a Go-lang based CLI Client for testing purposes",
	Long:  `zedcloud-client is a sample Go-lang based client for zedcloud`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Ex: go run zedcloud_client.go -c alpha -u user -p password edge-app --name=ubuntu-xenial-azure-iotedge-v1
func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringP("zedcloud-url", "", "",
		"Zedcloud url. (Ex: zedcontrol.<xxx>.zededa.net )")
	viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))
	rootCmd.PersistentFlags().StringP("user", "u", "", "Zedcontrol User")
	viper.BindPFlag("user", rootCmd.PersistentFlags().Lookup("user"))
	rootCmd.PersistentFlags().StringP("password", "p", "",
		"zedcontrol password")
	viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))
	rootCmd.PersistentFlags().String("api-token", "",
		"api-token to be used instead of username and passwords")
	viper.BindPFlag("api-token", rootCmd.PersistentFlags().Lookup("api-token"))
	rootCmd.PersistentFlags().String("name", "", "Object Name")
	viper.BindPFlag("name", rootCmd.PersistentFlags().Lookup("name"))
	rootCmd.PersistentFlags().String("edgenode-id", "", "EdgeNode ID")
	viper.BindPFlag("edgenode-id", rootCmd.PersistentFlags().Lookup("edgenode-id"))
	rootCmd.PersistentFlags().String("edgeapp-id", "", "EdgeApp ID")
	viper.BindPFlag("edgeapp-id", rootCmd.PersistentFlags().Lookup("edgeapp-id"))
	rootCmd.PersistentFlags().String("app-manifest", "", "App Manifest Json File")
	viper.BindPFlag("app-manifest", rootCmd.PersistentFlags().Lookup("app-manifest"))
	rootCmd.PersistentFlags().String("interface", "", "Interface name")
	viper.BindPFlag("intf", rootCmd.PersistentFlags().Lookup("intf"))
	rootCmd.PersistentFlags().String("network-instance", "", "network instance name")
	viper.BindPFlag("netinst-name", rootCmd.PersistentFlags().Lookup("netinst-name"))
	rootCmd.PersistentFlags().String("op", "", "Operation to Perform - "+
		"get / create / delete")
	viper.BindPFlag("op", rootCmd.PersistentFlags().Lookup("op"))
}

func initConfig() {
}

func init() {
	for _, cmdInfo := range cmdArray {
		cmdInfo.cobraCmd.Run = cmdHndlr
		rootCmd.AddCommand(cmdInfo.cobraCmd)
	}
	rootCmd.AddCommand(edgeNodeCmd)
	//rootCmd.AddCommand(edgeAppCmd)
	rootCmd.AddCommand(netInstCmd)
	rootCmd.AddCommand(appInstCmd)
}

func main() {
	Execute()
}
