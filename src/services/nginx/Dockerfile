FROM node:21-alpine3.18 as builder

WORKDIR /src/services/www

COPY ./src/services/www/package.json ./
RUN npm i

COPY ./src/services/www/ .
COPY ./src/packages/varsoapi/ ../../packages/varsoapi/

COPY ./scripts/apply_env.sh .

ARG SYSTEM_EXTERNAL_URL
ENV SYSTEM_EXTERNAL_URL=${SYSTEM_EXTERNAL_URL}

RUN chmod +x ./apply_env.sh && ./apply_env.sh ./src/config/config.ts
RUN npm run build


FROM nginx:latest

RUN mkdir -p /etc/letsencrypt
RUN curl -s https://raw.githubusercontent.com/certbot/certbot/master/certbot-nginx/certbot_nginx/_internal/tls_configs/options-ssl-nginx.conf > "/etc/nginx/options-ssl-nginx.conf"
RUN curl -s https://raw.githubusercontent.com/certbot/certbot/master/certbot/certbot/ssl-dhparams.pem > "/etc/nginx/ssl-dhparams.pem"

# Copy the init.sh script to the container
COPY ./src/services/nginx/init.sh /docker-entrypoint.d/

# Set execute permissions for the script
RUN chmod +x /docker-entrypoint.d/init.sh

# Copy the nginx configuration
COPY ./src/services/nginx/nginx.conf /etc/nginx/nginx.conf
COPY ./src/services/nginx/cors.conf /etc/nginx/cors.conf

# Copy the built frontend code
COPY --from=builder /src/services/www/build /var/www/varso.org/html

EXPOSE 443
EXPOSE 80


CMD ["nginx", "-g", "daemon off;"]