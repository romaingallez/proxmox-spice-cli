package power

import (
	"crypto/tls"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Telmate/proxmox-api-go/proxmox"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Power(vmID int, Status bool) {

	vmr := proxmox.NewVmRef(vmID)

	tlsConf := &tls.Config{InsecureSkipVerify: true}
	client, err := proxmox.NewClient(fmt.Sprintf("https://%s:8006/api2/json", viper.GetString("host")), nil, "", tlsConf, "", 300)
	if err != nil {
		log.Println(err)
	}
	err = client.Login(viper.GetString("login.username"), viper.GetString("login.password"), "")
	if err != nil {
		log.Println(err)
	}

	state, err := client.GetVmState(vmr)
	if err != nil {
		log.Println(err)
	}
	VmStatus, ok := state["status"].(string)
	if !ok {
		log.Println("Error converting status to string")
	}

	// log.Println(Status)
	if Status {
		if strings.Contains(VmStatus, "running") {
			fmt.Printf("VM %d is already running\n", vmID)
		} else {
			fmt.Printf("VM %d is not running, starting now\n", vmID)
			client.StartVm(vmr)
		}

	} else {
		if !strings.Contains(VmStatus, "false") {
			fmt.Printf("VM %d is not running\n", vmID)
		} else {
			fmt.Printf("VM %d is running, stopping now\n", vmID)
			client.StopVm(vmr)
		}
	}
}

func On(cmd *cobra.Command, args []string) {
	vmID, err := strconv.Atoi(args[0])
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Stoping the vmid %d\n", vmID)

	Power(vmID, true)

}
func Off(cmd *cobra.Command, args []string) {
	vmID, err := strconv.Atoi(args[0])
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Stoping the vmid %d\n", vmID)

	Power(vmID, false)

}
