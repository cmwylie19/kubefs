# Frontend UI running in NGINX
FROM arm64v8/node AS builder
LABEL description="KubeFS-Web" \
      maintainer="Casey Wylie casewylie@gmail.com"
# Set working directory
WORKDIR /app
# Copy all files from current directory to working dir in image
COPY . .
# install node modules and build assets
RUN npm i && npm run build 

# nginx state for serving content
# FROM arm64v8/nginx
FROM arm64v8/nginx
# Set working directory to nginx asset directory
WORKDIR /usr/share/nginx/html
# Remove default nginx static assets
RUN rm -rf ./*
COPY ./nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=builder /app/build .
# Containers run nginx with global directives and daemon off
ENTRYPOINT ["nginx", "-g", "daemon off;"]