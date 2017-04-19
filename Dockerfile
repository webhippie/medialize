FROM webhippie/alpine:latest
MAINTAINER Thomas Boerger <thomas@webhippie.de>

RUN apk update && \
  apk add \
    ca-certificates \
    bash && \
  rm -rf \
    /var/cache/apk/* && \
  addgroup \
    -g 1000 \
    medialize && \
  adduser -D \
    -h /var/lib/medialize \
    -s /bin/bash \
    -G medialize \
    -u 1000 \
    medialize

COPY medialize /usr/bin/

USER medialize
ENTRYPOINT ["/usr/bin/medialize"]

# ARG VERSION
# ARG BUILD_DATE
# ARG VCS_REF

# LABEL org.label-schema.version=$VERSION
# LABEL org.label-schema.build-date=$BUILD_DATE
# LABEL org.label-schema.vcs-ref=$VCS_REF
LABEL org.label-schema.vcs-url="https://github.com/webhippie/medialize.git"
LABEL org.label-schema.name="Templater"
LABEL org.label-schema.vendor="Thomas Boerger"
LABEL org.label-schema.schema-version="1.0"
