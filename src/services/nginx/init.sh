#!/bin/sh

certdir="/etc/letsencrypt/live"
realcert="${certdir}/.realcert"
fakecert="${certdir}/.fakecert"

# Only create dummy certificate if .realcert doesn't exist
if [ ! -f "$realcert" ]; then
    for domain in ${CERTBOT_DOMAINS}; do
        certPath="${certdir}/${domain}"
        mkdir -p "$certPath"
        openssl req -x509 -nodes -newkey rsa:4096 -days 1 -keyout "$certPath/privkey.pem" -out "$certPath/fullchain.pem" -subj "/CN=localhost"
    done

    mkdir -p "${certdir}"
    touch "$fakecert"
fi

# If we are not in DEV mode
if [ "$APP_ENV" != "DEV" ]; then
    # Wait for certbot to create a real certificate
    (
        while [ ! -f "$realcert" ]; do
            echo "Waiting for certbot to create .realcert.."
            sleep 5  # Adjust the sleep duration if needed
        done

        nginx -s reload
    ) &
fi
