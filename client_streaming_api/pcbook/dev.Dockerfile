# Dockerfile for golang(version:1.15.2)
FROM golang:1.15.2
MAINTAINER jihoon.kim <jihoon.kim@42dot.ai>

# install protobuf from source
RUN apt-get update && \
    apt-get -y install git unzip build-essential autoconf libtool vim net-tools openssh-server
RUN git clone https://github.com/protocolbuffers/protobuf && \
    cd protobuf && \
    ./autogen.sh && \
    ./configure && \
    make && \
    make install && \
    ldconfig && \
    make clean && \
    cd .. && \
    rm -r protobuf

# NOTE: for now, this docker image always builds the current HEAD version of gRPC
# After gRPC's beta release, the Dockerfile versions will be updated to
# build a specific version.

# Get the source from GitHub
RUN go get google.golang.org/grpc
# Install protoc-gen-go
RUN go get github.com/golang/protobuf/protoc-gen-go

#Source Code Copy for Develop
COPY ./ /go/src/VNet/
WORKDIR /go/src/VNet/

# Install all Dependancies in VNet(TBD)
#RUN go get ./...

# set password
RUN echo 'root:root' |chpasswd

# replace sshd_config
RUN sed -ri 's/^#?PermitRootLogin\s+.*/PermitRootLogin yes/' /etc/ssh/sshd_config
RUN sed -ri 's/UsePAM yes/#UsePAM yes/g' /etc/ssh/sshd_config

# make .ssh
RUN cat /etc/ssh/sshd_config
RUN mkdir /root/.ssh

RUN chown -R root:root /root/.ssh;chmod -R 700 /root/.ssh

RUN echo “StrictHostKeyChecking=no” >> /etc/ssh/ssh_config

RUN mkdir /var/run/sshd

RUN apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

EXPOSE 22

WORKDIR /root
