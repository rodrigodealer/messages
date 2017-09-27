FROM alpine:edge

COPY messages /opt/messages

RUN apk add --update curl && \
    rm -rf /var/cache/apk/*

HEALTHCHECK CMD curl --fail http://localhost:8080/healthcheck || exit 1

EXPOSE 8080
CMD ["/opt/messages"]
