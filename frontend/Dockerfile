# Frontend UI running in NGINX
FROM arm64v8/node
LABEL description="Kubefs-Web" \
      maintainer="Casey Wylie casewylie@gmail.com"
# Set working directory
WORKDIR /app
# Copy all files from current directory to working dir in image
COPY . .
# install node modules and build assets
RUN npm i
ENTRYPOINT ["npm", "start"]