# Go API Dockerfile

FROM golang:1.20-alpine

RUN apk update && apk add --no-cache git

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 80

CMD ["app"]
