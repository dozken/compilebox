############################################################
# Dockerfile to build sandbox for executing user code
# Based on Ubuntu
############################################################

FROM ubuntu:16.04

RUN apt-get update
RUN apt-get install -y gcc g++ gobjc curl sudo bc
RUN apt-get install -y mono-xsp2 mono-xsp2-base mono-vbnc
RUN apt-get install -y python