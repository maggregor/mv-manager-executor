
# Use the offical golang image to create a binary.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.16-buster as builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
RUN go build -v -o server

# Use the official Debian slim image for a lean production container.
# https://hub.docker.com/_/debian
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates software-properties-common curl wget unzip && \
    rm -rf /var/lib/apt/lists/*

# Install Terraform
ENV TERRAFORM_VERSION 0.15.4
ENV TERRAFORM_DIR /terraform
RUN wget https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip
RUN unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip -d /bin && rm -f terraform_${TERRAFORM_VERSION}_linux_amd64.zip
COPY --from=builder /app/terraform ${TERRAFORM_DIR}
WORKDIR ${TERRAFORM_DIR}

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/server /server
COPY --from=builder /app/docker-entrypoint.sh /docker-entrypoint.sh

# Run the web service on container startup.
CMD ["/docker-entrypoint.sh"]