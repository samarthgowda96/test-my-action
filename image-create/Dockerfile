# Container image that runs your code
FROM golang:1.21.4

# Copies your code file from your action repository to the filesystem path `/` of the container
WORKDIR /app
COPY . .
RUN go install samarthgowda96/test-my-action/image-create/cmd

# Code file to execute when the docker container starts up (`/go/bin/cmd`)
ENTRYPOINT ["/go/bin/cmd"]