FROM certbot/certbot:latest

COPY ./src/services/certbot/init.sh /etc/certbot/init.sh

RUN SLEEPTIME=$(awk 'BEGIN{srand(); print int(rand()*(3600+1))}'); echo "0 0,12 * * * root sleep $SLEEPTIME && certbot renew -q" | tee -a /etc/crontab > /dev/null

ENTRYPOINT ["/bin/sh", "/etc/certbot/init.sh"]