FROM  node:14.15-alpine as assets-builder

WORKDIR /assets

COPY ./package.json ./package-lock.json ./
COPY ./conf/webpack.*js ./conf/
COPY ./assets/. ./assets/

RUN npm install && npm run build


FROM golang:1.15-alpine as go-builder

ENV GO111MOD=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

WORKDIR /app

COPY . .

RUN go mod download

RUN go build


FROM alpine

COPY . .
COPY --from=assets-builder /assets/static/. ./static/
COPY --from=go-builder /app/go-scraper ./

EXPOSE 8080

ENTRYPOINT ["./go-scraper"]
