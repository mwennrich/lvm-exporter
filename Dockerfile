FROM debian:buster-slim

EXPOSE 9080

RUN apt-get update && \
    apt-get install lvm2

COPY dist/lvm-exporter_linux_amd64 /app/lvm-exporter
RUN chmod 755 /app/*

CMD [ "/app/lvm-exporter" ]
