FROM golang:1.17

# Add Zscaler stuff
ADD zscaler.crt /tmp/
ADD zscaler-bundle.pem /tmp/
ENV ZSCALER_BUNDLE=/tmp/zscaler-bundle.pem
ENV ZSCALER_CERT=/tmp/zscaler-bundle.pem

# Make go (and go get, etc.) play nicely with zscaler
ENV SSL_CERT_FILE=$ZSCALER_BUNDLE

## utils
RUN apt-get update
RUN apt-get install -y netcat
RUN apt-get install -y man
RUN apt-get install -y less
RUN apt-get install -y vim
RUN apt-get install -y build-essential
RUN apt-get install -y tree
RUN apt-get install -y bat
RUN echo 'alias bat=batcat' >> ~/.bashrc
RUN apt-get install -y silversearcher-ag
RUN apt-get install -y net-tools
RUN apt-get install -y file
RUN apt-get install -y unzip

# terraform
RUN wget https://releases.hashicorp.com/terraform/1.1.4/terraform_1.1.4_linux_amd64.zip
RUN unzip terraform_1.1.4_linux_amd64.zip
RUN mv terraform /usr/local/bin/
RUN terraform --version 