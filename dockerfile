

# download base image 
FROM golang:1.18.1-alpine3.15

#CGO_ENABLED
ENV CGO_ENABLED=0

# REQUEST FOR PORT
ENV PORT = 3000

# CREATE FOLDER INSIDE THE CONTAINER

RUN mkdir -p fictional-engine

# change directory to the new directory
ADD . /fictional-engine

WORKDIR /fictional-engine

# copy project configuration files
COPY go.mod go.sum ./

RUN go mod tidy

# download golang dependenices 
RUN go mod download

# build application
RUN go build main.go

# expose port
EXPOSE 3000


# User Information 
LABEL Info = Webui listen @ 127.0.0.1:3000

# start up 
CMD [ "/fictional-engine/main" ]