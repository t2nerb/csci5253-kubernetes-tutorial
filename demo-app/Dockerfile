FROM golang:alpine as build
WORKDIR /go/src/app

COPY . /go/src
RUN go get && go build

FROM alpine:latest
WORKDIR /app
COPY --from=build /go/src/app .

EXPOSE 8080
CMD ["/app/app"]
