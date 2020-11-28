FROM node:12.10.0-alpine

WORKDIR /workbench
ARG APIURL_ENV
ARG COOKIE_DOMAIN

COPY ./frontend .
RUN set -x && \
	apk update --no-cache && \
	npm install && \
	npm run build

USER node

ENV NODE_ENV="production"
ENV COOKIE_DOMAIN=${COOKIE_DOMAIN}
ENV TZ='Asia/Tokyo'

CMD ["node", "server.js"]

EXPOSE 3000
