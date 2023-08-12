FROM golang:1.20 as build
ENV GO111MODULE=on
RUN apt-get update && apt-get install -y build-essential
WORKDIR /usr/src/app
COPY . .
RUN go mod download
RUN make build

FROM alpine
MAINTAINER jake <cr.jake92@gmail.com>

ENV REPORT_LOG_PATH /home/catchreview/cr-api/logs
VOLUME ["/home/catchreview/cr-api/logs"]

RUN mkdir -p /home/catchreview/cr-api/logs
WORKDIR /home/catchreview/cr-api
COPY --from=build /usr/src/app/bin/cr-api .
CMD ["./cr-api"]
EXPOSE 8088