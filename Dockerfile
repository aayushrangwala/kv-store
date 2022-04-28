FROM golang:1.18-bullseye AS build
WORKDIR /go/src/kv-store
COPY ./ ./
ARG GIT_COMMIT
ARG GOPROXY

RUN go build -o /go/bin/kv-store .

FROM gcr.io/distroless/base-debian11
COPY --from=build /go/bin/kv-store /usr/local/bin/kv-store
ENTRYPOINT [ "/usr/local/bin/kv-store" ]
