FROM postgres:latest

# Install the Certbot software
# RUN apt-get update && apt-get install -y \
#     certbot \
#     && rm -rf /var/lib/apt/lists/*

# # Generate a Let's Encrypt SSL certificate
# RUN certbot certonly --standalone --non-interactive --agree-tos \
#     --email thenewl369@gmail.com \
#     --domains yourdomain.com

# # Copy the SSL certificate and private key to the appropriate location
# COPY /etc/letsencrypt/live/yourdomain.com/fullchain.pem /etc/ssl/certs/server.crt
# COPY /etc/letsencrypt/live/yourdomain.com/privkey.pem /etc/ssl/private/server.key

# # Modify the PostgreSQL configuration to use the SSL certificate and private key
# RUN echo "ssl_cert_file = '/etc/ssl/certs/server.crt'" >> /var/lib/postgresql/data/postgresql.conf \
#     && echo "ssl_key_file = '/etc/ssl/private/server.key'" >> /var/lib/postgresql/data/postgresql.conf
