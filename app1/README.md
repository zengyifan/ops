# Overview
This project `app1` serves as one of the backends

# Dependencies
The project requires the following dependencies:
- Make
- Go (>=1.23)
- Redis
- MySQL
- Docker (Optional)

# Quickstart

```bash
# Make a copy of the config file and setup the Redis & MySQL service details
cp configs/config.yaml configs/config.dev.yaml
vim configs/config.dev.yaml

# Build the project to a binary
make build

# Run the binary with the created config file
./_output/platforms/linux/amd64/app1 -c configs/config.dev.yaml
```

For more detailed steps of setting up Redis, MySQL and testing of the services,
refer to the <a href="#installation">Installation</a> section

# Installation
The Installation section covers the following:
1. Setup - Setting up dependencies/external services and pre-requisites for the
   application to run properly
2. Running - Running the project
3. Building - Building the project
4. Testing - Test if the APIs are working as intended

## 1. Setup
Before building and running this project, ensure that you have an instance of
Redis and MySQL running.

**Local Deployment**

For local deployments, install Redis and MySQL on your system and start them

For the full details, refer to the `deploy/local/README.md` guide

**Cloud Deployment**

For cloud deployments (e.g. Tencent Cloud), please refer to the deployment guide
at `deploy/tencentcloud/manual/04Backend1部署.md`

### Configuration File
After setting up the external services (Redis & MySQL) of the project, we will
need to configure `app1` to use these external services through the config file.

Make a copy of the file given example config file at `configs/config.yaml` and
edit the copy.

```bash
# Copy the config file and edit it
cp configs/config.yaml configs/config.dev.yaml
vim configs/config.dev.yaml
```

```yaml
# Change the placeholders appropriately
mysql:
  host: <mysql-hostname>
  port: 3306
  user: root
  password: P@ssw0rd
  dbname: app1

redis:
  addr: <redis-hostname>:6379
  password: P@ssw0rd
  db: 0

log:
  filePath: "/tmp/app1.log"
```

## 2. Running
This project uses `make` as its build system with its build build
processes/recipes documented in the `Makefile`

To run the project, you can use either of the following commands:
```bash
# Using Makefile
make run

# Or using go to run directly (optional)
go mod tidy
go run cmd/main.go -c configs/config.dev.yaml
```

## 3. Building
Similar to building the project, you can run either of the following commands:
```bash
# Using Makefile
make build

# Or using go to build directly (optional)
export GOOS=linux
export GOARCH=amd64
go build -o _output/platforms/linux/amd64/app1 cmd/main.go
```

The build produces a binary located at
`_output/platforms/<platform-name>/<platform-arch>/app1`

This binary can then be run with the config file:
```bash
./_output/platforms/<platform-name>/<platform-arch>/app1 -c config.dev.yaml
```

## 4. Testing
After running the project, you can test if it works as intended with the
`make test` command:
```bash
# Using Makefile
make test-api

# Or manually curl the endpoints
export APP1_IP="<my-app1-server-ip>
curl http://$APP1_IP:8888/hello 
curl http://$APP1_IP:8888/users 
curl http://$APP1_IP:8888/groups
curl http://$APP1_IP:8888/auth?user=admin&pwd=P@ssw0rd
```


# Deploy with `systemd`
If the host you're deploying on is a GNU/Linux based system (e.g. Centos/Ubuntu),
you can also deploy the application as a `systemd` service.

After building the application, we will need to install the application to
server by running the following commands:
```bash
# Create a directory for the config file and binary
mkdir -p /usr/local/bin/app1/configs

# Copy the app1 binary and config file to the created directory
cp _output/platforms/linux/amd64/app1 /usr/local/bin/app1/
chmod +x /usr/local/bin/app1/app1
cp configs/config.yaml /usr/local/bin/app1/configs/
```

After you have installed the binary and config file, you can the the following
commands to create the `systemd` unit files
```bash
# Create the Systemd Unit files
cat >  /etc/systemd/system/app1.service <<EOF
[Unit]
Description=Go Application Service

[Service]
ExecStart=/bin/bash -c '/usr/local/bin/app1/app1 -c /usr/local/bin/app1/configs/config.yaml >> /tmp/app1.log 2>&1'
WorkingDirectory=/usr/local/bin/
User=root
Restart=always
Type=simple
KillMode=mixed

[Install]
WantedBy=multi-user.target
EOF

# Reload the systemd daemon to let it read the new units created
systemctl daemon-reload

# Restart the app1.service and view the status
systemctl enable --now app1.service
systemctl status app1.service

# When not needed, you can either disable or stop the app1.service
systemctl stop app1.service
systemctl restart app1.service
```
