FROM debian:9.9-slim

EXPOSE 9080

COPY dist/lvm-exporter_linux_amd64 /app/lvm-exporter
RUN chmod 755 /app/*

CMD [ "/app/lvm-exporter" ]
