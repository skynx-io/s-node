FROM alpine
RUN apk update && apk add --update --no-cache ca-certificates bash
COPY skynx-node /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/skynx-node"]
