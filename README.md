# Setup EC2 instance

* Create and launch EC2
* Set security group's inbound rule to accept SSH (22) and HTTP(80)

# Connect to EC2 instance and install go

```bash
./connect-to-ec2.sh
sudo apt-get update
sudo apt-get install golang
# ctrl+c to return to local machine
```

# Copy our source code to ec2

```bash
./connect-to-ec2.sh
mkdir -p app
# ctrl+c to return to local machine
./copy-to-ec2.sh
```

# Build and run the program

```bash
./connect-to-ec2.sh
cd ~/app
./build-and-run.sh
```

# Deploy in EC2 using github workflow

```bash
./connect-to-ec2.sh
sudo apt-get docker.io
# make docker accessible by every user (don't do this actually)
sudo chmod 777 /var/run/docker.sock
# ctrl+c to return to local machine
```

- Pull this repository from EC2
- Remove existing docker container if exist
- Docker build
- Docker run/start