# Proxmox SPICE CLI

A compact Golang program that enables connection to a virtual machine (VM) on a Proxmox server using SPICE from the command line interface (CLI) across Linux, Windows, and Mac platforms.

Inspired by https://github.com/Elbandi/proxmox-spice-quickconnect

Proxmox SPICE CLI allows users to easily start, stop, and connect to a VM on a Proxmox host using the command line interface. The project is written in Golang.

# Configuration

An example configuration file can be found inside the release archive or in the git repository. Configure the program by placing the configuration file in the following locations:

## Linux

`~/.proxmox-spice-cli`

## Windows

`C:\Users\$env:USERNAME\.proxmox-spice-cli`

For Windows, please modify the SPICE path from:
```json
    "path": "/usr/bin/remote-viewer"
```
to the path of the remote-viewer.exe, making sure to escape the slashes like this: `\\`
