# This is the base image of the services.
# In this stage we prepare all required common data to dockerize our services.
FROM golang:1.15 as preparer

RUN apt-get update                                                        && \
  DEBIAN_FRONTEND=noninteractive apt-get install -yq --no-install-recommends \
    curl git unzip                                                           \
  && rm -rf /var/lib/apt/lists/*

# Install solidity stuff
# RUN add-apt-repository ppa:ethereum/ethereum
# RUN apt-get update
# RUN apt-get install solc ethereum

# Install Protobuf.
ARG PROTOBUF_VERSION=3.12.4
RUN curl -sOL "https://github.com/google/protobuf/releases/download/v${PROTOBUF_VERSION}/protoc-${PROTOBUF_VERSION}-linux-x86_64.zip" && \
  unzip protoc-*.zip              && \
  mv bin/protoc /usr/local/bin    && \
  mv include/* /usr/local/include && \
  rm -f protoc-*.zip

# Go to the root of the project.
WORKDIR /go/src/github.com/begmaroman/beaconspot/

# Switch on the Go modules. Go modules will be switched on by default from Go 1.13 onwards.
ENV GO111MODULE=on

# Copy module files
COPY go.mod .
COPY go.sum .

# Download project dependencies
RUN go mod download

# Install go tools.
RUN go install \
  github.com/go-openapi/runtime \
  github.com/tylerb/graceful \
  github.com/jessevdk/go-flags \
  github.com/golang/protobuf/protoc-gen-go \
  github.com/go-swagger/go-swagger/cmd/swagger

FROM preparer as builder

# Copy the project on build trigger.
COPY . .

 # Remove vendored deps
RUN rm -rf ./vendor

# Install the service binary.
RUN CGO_CFLAGS_ALLOW="-D__BLST_PORTABLE__" CGO_CFLAGS="-D__BLST_PORTABLE__" CGO_ENABLED=1 GOOS=linux go install -a -tags blst_enabled -ldflags "-linkmode external -extldflags \"-static -lm  -msoft-float -Wl,-y,__sigsetjmp_aux\"" ./cmd/beaconspot

# Stage 2: Prepare all required data to run the service.
FROM alpine:3.9 as runner

# Install ca-certificates, bash
RUN apk -v --update add ca-certificates bash

# Copy entrypoint and service executable.
COPY --from=builder /go/bin/beaconspot /go/bin/beaconspot

# This is needed for healthcheck
EXPOSE 5678

ENTRYPOINT ["/go/bin/beaconspot"]
