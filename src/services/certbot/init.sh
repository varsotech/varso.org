#!/bin/sh

function keep_running() {
    # Keep running in background to let our cron cert renewal job to run
    echo "Running in the background"
    tail -f /dev/null
}

if [ "${USE_TEST_CERT}" = "true" ]; then 
    echo "Using test certificate"
    keep_running
fi

domain_args="" 
for domain in ${CERTBOT_DOMAINS}; do
    domain_args="$domain_args -d $domain"
done

if [ -e "/etc/letsencrypt/live/.realcert" ]; then
    echo "/etc/letsencrypt/live/.realcert exists, not requesting certificate"
    keep_running
fi

echo "Waiting for NGINX to create .fakecert.."
while [ ! -f "/etc/letsencrypt/live/.fakecert" ]; do
    printf "."
    sleep 5  # Adjust the sleep duration if needed
done

for domain in ${CERTBOT_DOMAINS}; do
    rm -Rf /etc/letsencrypt/live/$domain
    rm -Rf /etc/letsencrypt/archive/$domain
    rm -Rf /etc/letsencrypt/renewal/$domain.conf
done

rm /etc/letsencrypt/live/.fakecert
mkdir -p /var/www/certbot

certbot certonly --webroot -w /var/www/certbot --email ${CERTBOT_EMAIL} $domain_args --rsa-key-size 4096 --agree-tos --keep-until-expiring --non-interactive || exit 1

touch /etc/letsencrypt/live/.realcert
keep_running