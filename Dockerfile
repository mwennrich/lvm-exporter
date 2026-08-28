
FROM golang:1.27-alpine AS builder
RUN apk add make binutils
COPY / /work
WORKDIR /work
RUN make lvm-exporter

FROM alpine:3.24
RUN apk add lvm2
COPY --from=builder /work/bin/lvm-exporter /lvm-exporter
USER root
ENTRYPOINT ["/lvm-exporter"]

EXPOSE 9080
