FROM golang:alpine AS build

RUN apk add --update git
WORKDIR /go/src/GourseAPI
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/gourse-api cmd/api/main.go

# Building image with the binary
FROM scratch
COPY --from=build /go/bin/gourse-api /go/bin/gourse-api
ENTRYPOINT ["/go/bin/gourse-api"]