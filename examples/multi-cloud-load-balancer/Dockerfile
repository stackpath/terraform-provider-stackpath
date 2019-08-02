FROM traefik:v1.7-alpine
RUN apk add gettext
COPY ./traefik.toml /etc/traefik/traefik.toml.tmpl
COPY ./entrypoint.sh /entrypoint.sh
ENTRYPOINT [ "sh", "/entrypoint.sh" ]
