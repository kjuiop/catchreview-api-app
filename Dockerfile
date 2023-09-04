FROM golang:1.20 as build
ENV GO111MODULE=on
RUN apt-get update && apt-get install -y build-essential
WORKDIR /usr/src/app
COPY . .
RUN go mod download
RUN make build

FROM alpine
MAINTAINER jake <arneg0shua@gmail.com>

WORKDIR /home/gig/cr-api
COPY --from=build /usr/src/app/bin/cr-api .
CMD ["./cr-api"]
EXPOSE 8080
