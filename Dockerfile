FROM  node:14.15-alpine as assets-builder

WORKDIR /assets

COPY ./package.json ./package-lock.json ./
COPY ./conf/webpack.*js ./conf/
COPY ./assets/. ./assets/

RUN npm install && npm run build

FROM golang:1.15-buster

ENV GO111MOD=on \
    PORT=80

RUN go get github.com/beego/beego && \
  go get github.com/beego/bee

RUN apt-get update && apt-get install -y --no-install-recommends nginx

WORKDIR /app

COPY . .

COPY conf/nginx/app.conf.template /etc/nginx/conf.d/default.conf

RUN go mod download

COPY --from=assets-builder /assets/static/. ./static/

EXPOSE $PORT

ENTRYPOINT ["make", "production"]
