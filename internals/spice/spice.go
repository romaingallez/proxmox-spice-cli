package spice

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"

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

	tlsConf := &tls.Config{InsecureSkipVerify: true}
	client, err := proxmox.NewClient(fmt.Sprintf("https://%s:8006/api2/json", viper.GetString("host")), nil, "", tlsConf, "", 300)
	if err != nil {
		log.Println(err)
	}
	err = client.Login(viper.GetString("login.username"), viper.GetString("login.password"), "")
	if err != nil {
		log.Println(err)
	}

	// log.Println(client.GetVmList())

	vmr := proxmox.NewVmRef(vmID)
	log.Println(vmr)
	config, err := client.GetVmSpiceProxy(vmr)
	if err != nil {
		log.Println(err)
	}

	// log.Println(config)
	log.Println(config["tls-port"], config["delete-this-file"], config["title"], config["proxy"], config["toggle-fullscreen"],
		config["type"], config["release-cursor"], config["host-subject"], config["password"], config["secure-attention"],
		config["host"], config["ca"])

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

	_, err = fmt.Fprintf(stdin, "[virt-viewer]\n"+
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
		config["host"], config["ca"])

	if err != nil {
		log.Println(err)
	}

	go func() {
		err = subProcess.Wait()
		fmt.Printf("Command finished with error: %v", err)
	}()

}
