FROM  node:14.15-alpine as assets-builder

WORKDIR /assets

COPY ./package.json ./package-lock.json ./
COPY ./conf/webpack.*js ./conf/
COPY ./assets/. ./assets/

RUN npm install && npm run build

FROM golang:1.15-buster

ENV GO111MOD=on

RUN go get github.com/astaxie/beego && \
  go get github.com/beego/bee

WORKDIR /app

COPY . .

RUN go mod download

COPY --from=assets-builder /assets/static/. ./static/

EXPOSE 8080

ENTRYPOINT ["make", "production"]
