/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/romaingallez/proxmox-spice-cli/internals/power"
	"github.com/spf13/cobra"
)

// powerCmd represents the power command
var powerCmd = &cobra.Command{
	Use:   "power",
	Short: "Control the state of the VM",
	Long:  `This command allows you to turn a virtual machine on or off.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please us the subcommand command on or off")
	},
}

func init() {
	rootCmd.AddCommand(powerCmd)

	powerCmd.AddCommand(&cobra.Command{
		Use:   "on",
		Short: "Turn on a virtual machine",
		Run:   power.On,
	})

	powerCmd.AddCommand(&cobra.Command{
		Use:   "off",
		Short: "Turn off a virtual machine",
		Run:   power.Off,
	})

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	powerCmd.PersistentFlags().String("type", "", "on/off")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// powerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
