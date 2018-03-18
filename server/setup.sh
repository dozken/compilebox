#!/bin/bash

apt update
apt install -y git upstart upstart-sysv docker.io

cd /opt
git clone https://github.com/chekalskiy/compilebox /opt/compilebox

cp /opt/compilebox/server/service_compilebox.conf /etc/init/service_compilebox.conf

docker build -t virtual_machine /opt/compilebox/bin

reboot
