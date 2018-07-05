FROM webhippie/alpine:latest

LABEL maintainer="Thomas Boerger <thomas@webhippie.de>" \
  org.label-schema.name="Medialize" \
  org.label-schema.vendor="Thomas Boerger" \
  org.label-schema.schema-version="1.0"

RUN apk add --no-cache ca-certificates mailcap bash && \
  addgroup -g 1000 medialize && \
  adduser -D -h /var/lib/medialize -s /bin/bash -G medialize -u 1000 medialize

USER medialize
ENTRYPOINT ["/usr/bin/medialize"]
CMD ["help"]

COPY dist/binaries/medialize-*-linux-amd64 /usr/bin/
