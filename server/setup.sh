#!/bin/bash

apt update
apt-get update

apt-get install -y apt-transport-https ca-certificates curl gnupg-agent software-properties-common

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -

add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"

apt-get update

apt-get install -y docker-ce docker-ce-cli containerd.io

apt install -y git upstart upstart-sysv

cd /opt
git clone https://github.com/chekalskiy/compilebox /opt/compilebox

cp /opt/compilebox/server/service_compilebox.conf /etc/init/service_compilebox.conf

docker build -t virtual_machine /opt/compilebox/bin

reboot
