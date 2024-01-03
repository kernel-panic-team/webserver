FROM nginx:latest

COPY nginx.conf /etc/nginx/nginx.conf
COPY ../ssl/domain.cert.pem /etc/ssl/certs/domain.cert.pem
COPY ../ssl/private.key.pem /etc/ssl/private/private.key.pem