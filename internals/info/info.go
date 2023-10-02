/*
Copyright © 2023 Romain GALLEZ

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
package info

import (
	"crypto/tls"
	"fmt"
	"log"

	"github.com/Telmate/proxmox-api-go/proxmox"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func List(cmd *cobra.Command, args []string) {
	// log.Println("List")

	tlsConf := &tls.Config{InsecureSkipVerify: true}
	client, err := proxmox.NewClient(fmt.Sprintf("https://%s:8006/api2/json", viper.GetString("host")), nil, "", tlsConf, "", 300)
	if err != nil {
		log.Println(err)
	}
	err = client.Login(viper.GetString("login.username"), viper.GetString("login.password"), "")
	if err != nil {
		log.Println(err)
	}

	guests, err := proxmox.ListGuests(client)
	if err != nil {
		log.Panic(err)
	}

	for _, guest := range guests {
		log.Printf("VmID %d Name: %s Status: %s\n", guest.Id, guest.Name, guest.Status)
	}
}
