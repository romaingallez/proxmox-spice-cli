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
package spice

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/Telmate/proxmox-api-go/proxmox"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Spice(cmd *cobra.Command, args []string) {
	// log.Println(args)
	vmID, err := strconv.Atoi(args[0])
	if err != nil {
		log.Println(err)
	}

	// spicePath := viper.Get("spice.path")

	host := viper.GetString("host")

	tlsConf := &tls.Config{InsecureSkipVerify: true}
	client, err := proxmox.NewClient(fmt.Sprintf("https://%s:8006/api2/json", host), nil, "", tlsConf, "", 300)
	if err != nil {
		log.Println(err)
	}
	err = client.Login(viper.GetString("login.username"), viper.GetString("login.password"), "")
	if err != nil {
		log.Println(err)
	}

	// log.Println(client.GetVmList())

	vmr := proxmox.NewVmRef(vmID)
	config, err := client.GetVmSpiceProxy(vmr)
	if err != nil {
		log.Println(err)
	}

	// convert config["proxy"] to a string

	proxy := config["proxy"].(string)

	log.Println(proxy)

	// log.Panicln(proxy, host)

	proxyURL, err := url.Parse(proxy)
	if err != nil {
		log.Println(err)
	}
	// extract the port from the host
	port := proxyURL.Port()
	log.Println(port)

	// log.Println(hostURL)

	if !strings.Contains(proxy, host) {
		log.Println(config, "does not contains", host)
		proxy = fmt.Sprintf("http://%s:%s", host, port)
	}

	log.Println(proxy)

	config["proxy"] = proxy

	// wait for user keypress to continue
	log.Println("Press any key to continue...")
	var input string
	fmt.Scanln(&input)

	// log.Println(config)
	// log.Println(config["tls-port"], config["delete-this-file"], config["title"], config["proxy"], config["toggle-fullscreen"],
	// 	config["type"], config["release-cursor"], config["host-subject"], config["password"], config["secure-attention"],
	// 	config["host"], config["ca"])

	// subProcess := exec.Command(viper.GetString("spice.path"), "--debug", "-")
	subProcess := exec.Command(viper.GetString("spice.path"), "-")

	stdin, err := subProcess.StdinPipe()
	if err != nil {
		log.Println(err)
	}
	defer stdin.Close()
	devnull, err := cmd.Flags().GetBool("devnull")
	if err != nil {
		log.Println(err)
	}
	if devnull {
		devnull, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0755)
		if err != nil {
			log.Println(err)
		}

		subProcess.Stderr = devnull
		subProcess.Stdout = devnull
	} else {
		subProcess.Stderr = os.Stderr
		subProcess.Stdout = os.Stdout
	}

	err = subProcess.Start()

	if err != nil {
		log.Println(err)
	}

	command := fmt.Sprintf(
		"[virt-viewer]\n"+
			"tls-port=%.0f\n"+
			"delete-this-file=%.0f\n"+
			"title=%s\n"+
			"proxy=%s\n"+
			"toggle-fullscreen=%s\n"+
			"type=%s\n"+
			"release-cursor=%s\n"+
			"host-subject=%s\n"+
			"password=%s\n"+
			"secure-attention=%s\n"+
			"host=%s\n"+
			"ca=%s\n",
		config["tls-port"], config["delete-this-file"], config["title"], config["proxy"], config["toggle-fullscreen"],
		config["type"], config["release-cursor"], config["host-subject"], config["password"], config["secure-attention"],
		config["host"], config["ca"],
	)
	_, err = fmt.Fprint(stdin, command)

	if err != nil {
		log.Println(err)
	}

	log.Println(command)

	go func() {
		err = subProcess.Wait()
		log.Printf("Command finished with error: %v", err)
	}()

}
