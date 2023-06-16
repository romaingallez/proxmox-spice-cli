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
