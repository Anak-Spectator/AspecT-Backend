FROM golang:alpine as dev

RUN apk add git

# copy the local package files to the container workspace
ADD . /apps

# Setting up working directory
WORKDIR /apps
RUN pwd
# copy module file
COPY go.mod go.sum ./

# Install dependencies
ENV GO111MODULE on
RUN go mod download
RUN go mod tidy

# build binary
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o ./dist/backend_apps .main.go

# Store and run ass Production 
FROM alpine:3.9.2 as prod

WORKDIR /root/

COPY --from=dev /apps/dist .

ENTRYPOINT ["./backend_apps"]