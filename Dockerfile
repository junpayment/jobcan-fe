FROM golang:1.11.5 AS build-env

ADD . /work
WORKDIR /work

RUN CGO_ENABLED=0 go build -o app

FROM selenium/standalone-chrome

COPY --from=build-env /work/app ./app
ADD entry.sh /entry.sh

ENTRYPOINT ["/entry.sh"]
