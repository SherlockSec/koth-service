# King of The Hill Service

## An improved KoTH service including support for flag randomization and cross platform.

### What to do

```bash
# Clone the repo
git clone https://github.com/SherlockSec/koth-service
cd koth-service

# Edit main.go and change the following constants:
# const kingPath = "king.txt" 	// Path to king file
# const mapPath = "map.txt" 	// Path to map file
# const flags = 4 	        // amount of flags

# Then build:

# Windows
env GOOS=windows GOARCH=amd64 go build .
# Linux
env GOOS=linux GOARCH=amd64 go build .

# Then create the service config for your OS, e.g. use systemctl for Linux and sc for Windows
```

### Credits

[NinjaJc01](https://github.com/NinjaJc01), for the base `http.serve` code for Port 9999  
[TryHackMe](https://tryhackme.com), for actually making the gamemode and implementing this on the backend  
