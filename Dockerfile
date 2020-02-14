FROM golang:alpine

# Install bash to use at compose commands
RUN apk update && apk add bash
#RUN apk add --no-cache bash # for alpine 3+

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GO_ENV=docker

# Move to working directory /build
WORKDIR /build

# Copy the code into the container
COPY . .

# Download dependency using go mod
RUN go mod download

# Build the application
RUN go build -o main .

# Move wait for scipt for root folder
RUN mkdir /scripts
RUN mv wait-for-it.sh /scripts

# Make app folder
RUN mkdir /app

# Copy binary from build to app folder
RUN cp main /app

# Command to run when starting the container (called by wait-for-it.sh into app docker compose file)
#ENTRYPOINT /app/main

# Export necessary port
EXPOSE 3000
