############################################################
# Dockerfile to build sandbox for executing user code
# Based on Ubuntu
############################################################

FROM ubuntu:16.04

RUN apt-get update
RUN apt-get install -y gcc g++ gobjc curl sudo bc
RUN apt-get install -y python build-essential libssl-dev libffi-dev python-dev

# c#
RUN apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 3FA7E0328081BFF6A14DA29AA6A19B38D3D831EF
RUN apt install apt-transport-https
RUN echo "deb https://download.mono-project.com/repo/ubuntu stable-xenial main" > /etc/apt/sources.list.d/mono-official-stable.list
RUN apt update
RUN apt install -y mono-devel mono-xsp4
