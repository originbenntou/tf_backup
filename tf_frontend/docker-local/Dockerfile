
####################################
# ビルド用
####################################
FROM node:12.10.0-alpine AS build

WORKDIR /workbench
ARG APIURL_ENV="http://localhost:30880"

COPY ./frontend .
RUN set -x && \
	apk update --no-cache && \
	npm install && \
	npm run build

####################################
# 実行環境用
####################################
FROM node:12.10.0-alpine

WORKDIR /frontend
RUN set -x apk update --no-cache
ENV TZ='Asia/Tokyo'

COPY --from=build /workbench/next.config.js .
COPY --from=build /workbench/.next/ .next/
COPY --from=build /workbench/node_modules/ node_modules/

USER node

CMD ["node", "node_modules/next/dist/bin/next", "start", "-p", "3000"]

EXPOSE 3000