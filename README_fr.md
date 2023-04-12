# Proxmox SPICE CLI

Un programme Golang compact qui permet la connexion à une machine virtuelle (VM) sur un serveur Proxmox en utilisant SPICE depuis l'interface de ligne de commande (CLI) sur les plates-formes Linux, Windows et Mac.

Inspiré de https://github.com/Elbandi/proxmox-spice-quickconnect

Proxmox SPICE CLI permet aux utilisateurs de démarrer, d'arrêter et de se connecter facilement à une VM sur un hôte Proxmox en utilisant l'interface de ligne de commande. Le projet est écrit en Golang.

# Configuration

Un fichier de configuration exemple peut être trouvé à l'intérieur de l'archive de sortie ou dans le dépôt Git.

En utilisant le drapeau --config, vous pouvez utiliser une configuration placée dans n'importe quel dossier.

Configurez le programme en plaçant le fichier de configuration aux emplacements suivants :

## Linux

`~/.proxmox-spice-cli`

## Windows

`C:\Users\$env:USERNAME\.proxmox-spice-cli`

Pour Windows, veuillez modifier le chemin SPICE de :
```json
    "path": "/usr/bin/remote-viewer"
```
au chemin de remote-viewer.exe, en veillant à échapper les barres obliques comme ceci : `\\`

# Spice
Activez Spice dans la configuration de la VM dans proxmox

Sur Windows, n'oubliez pas d'installer l'addon invité Spice 
![](docs/spice_windows.png)