FROM golang:1.17

# Add Zscaler stuff
ADD zscaler.crt /tmp/
ADD zscaler-bundle.pem /tmp/
ENV ZSCALER_BUNDLE=/tmp/zscaler-bundle.pem
ENV ZSCALER_CERT=/tmp/zscaler-bundle.pem

# Make go (and go get, etc.) play nicely with zscaler
ENV SSL_CERT_FILE=$ZSCALER_BUNDLE