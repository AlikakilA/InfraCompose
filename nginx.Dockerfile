FROM nginx:alpine

# Copier le fichier de configuration nginx dans le conteneur
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Copier les certificats dans le conteneur
COPY certs/b1.com.crt /etc/ssl/certs/b1.com.crt
COPY certs/b1.com.key /etc/ssl/private/b1.com.key
COPY certs/blog.b1.com.crt /etc/ssl/certs/blog.b1.com.crt
COPY certs/blog.b1.com.key /etc/ssl/private/blog.b1.com.key
