FROM node:12
WORKDIR /app
LABEL maintainer="zhshch<zhshch@athorx.com>"

COPY package*.json ./
RUN npm install && npm install tsc -g
COPY . .
RUN npm run build

EXPOSE 4000
CMD ["npm", "run", "server"]