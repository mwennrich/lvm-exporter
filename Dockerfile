
FROM golang:1.23-alpine as builder
RUN apk add make binutils
COPY / /work
WORKDIR /work
RUN make lvm-exporter

FROM alpine:3.20
RUN apk add lvm2
COPY --from=builder /work/bin/lvm-exporter /lvm-exporter
USER root
ENTRYPOINT ["/lvm-exporter"]

EXPOSE 9080
