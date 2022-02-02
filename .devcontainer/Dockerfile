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