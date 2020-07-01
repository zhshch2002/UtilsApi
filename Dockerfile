FROM node:12
WORKDIR /app
LABEL maintainer="zhshch<zhshch@athorx.com>"

COPY package*.json ./
RUN npm config set sharp_binary_host "https://npm.taobao.org/mirrors/sharp" && \
npm config set sharp_libvips_binary_host "https://npm.taobao.org/mirrors/sharp-libvips" && \
npm install && npm install typescript -g
COPY . .
RUN npm run build

EXPOSE 4000
CMD ["npm", "run", "server"]