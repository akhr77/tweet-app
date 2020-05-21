FROM node:10.19.0-alpine

COPY . /app

RUN apk update && \
    npm install -g npm && \
    npm install -g @vue/cli \
    npm run build
