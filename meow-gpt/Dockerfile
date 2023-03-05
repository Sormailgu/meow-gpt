# Basic image, built based on golang's alpine image -- compilation phase
FROM golang:alpine AS builder

# Author
MAINTAINER sormailgu

# Global working directory
WORKDIR /go/app

# Copy all files in the current directory where the Dockerfile is run to the target directory
COPY . /go/app

# Environment variables
ENV TELEGRAM_BOT_TOKEN=''
ENV OPENAI_API_KEY=''
ENV OPENAI_MODEL='gpt-3.5-turbo'
ENV OPENAI_TEMPERATURE=1
ENV OPENAI_MAX_TOKENS=2048
ENV OPENAI_TOP_P=1
ENV OPENAI_FREQUENCY_PENALTY=0
ENV OPENAI_PRESENCE_PENALTY=0

# Used as a proxy to download the packages that the go project depends on
#ENV GOPROXY https://goproxy.cn,direct

# Compile, close CGO, prevent the compiled files from having dynamic links, and some c libraries in the alpine image do not have any errors, so there is no file error
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build main.go

# Use the lightweight image of alpine as the base image -- run phase
FROM alpine AS runner

# Global working directory
WORKDIR /go/app

# Copy the running file compiled during the compilation phase to the target directory
COPY --from=builder /go/app/main .

# Copy the config folder in the compilation phase to the target directory
COPY --from=builder /go/app/config ./config

# Ports to be exposed
#EXPOSE 8888

# Directory that can be plugged in
#VOLUME ["/go/app/config","/go/app/log"]

# The real command triggered by the docker run command (equivalent to directly running the compiled executable file)
ENTRYPOINT ["./main"]