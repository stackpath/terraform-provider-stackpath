#!/bin/bash

envsubst < "/etc/traefik/traefik.toml.tmpl" > /etc/traefik/traefik.toml
traefik -C /etc/traefik/traefik.toml -l debug
