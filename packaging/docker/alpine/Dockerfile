FROM golang:alpine AS build
ENV CGO_ENABLED=0
RUN apk add --no-cache --update git
ADD . /go/src/app
WORKDIR /go/src/app
RUN go build \
    -o miniflux \
    -ldflags="-s -w -X 'miniflux.app/v2/version.Version=`git describe --tags --abbrev=0`' -X 'miniflux.app/v2/version.Commit=`git rev-parse --short HEAD`' -X 'miniflux.app/v2/version.BuildDate=`date +%FT%T%z`'" \
    main.go

FROM alpine:latest

LABEL org.opencontainers.image.title=Miniflux
LABEL org.opencontainers.image.description="Miniflux is a minimalist and opinionated feed reader"
LABEL org.opencontainers.image.vendor="Frédéric Guillot"
LABEL org.opencontainers.image.licenses=Apache-2.0
LABEL org.opencontainers.image.url=https://miniflux.app
LABEL org.opencontainers.image.source=https://github.com/miniflux/v2
LABEL org.opencontainers.image.documentation=https://miniflux.app/docs/

EXPOSE 8080
ENV LISTEN_ADDR 0.0.0.0:8080
RUN apk --no-cache add ca-certificates tzdata
COPY --from=build /go/src/app/miniflux /usr/bin/miniflux
USER nobody
CMD ["/usr/bin/miniflux"]
