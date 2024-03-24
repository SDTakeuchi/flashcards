# 1st stage build for prod env
FROM golang:1.22.1-alpine as builder

WORKDIR /go/flashcards/

RUN apk update \
&&  apk add --no-cache bash gcc musl-dev \
&&  rm -rf /var/lib/apt/lists/* \
&&  export GO111MODULE=on

COPY ./src/ .

RUN go mod tidy \
&&  go build -o /go_app .

# prod env; runs built image
FROM alpine as prod

COPY --from=builder /go_app .
COPY ./app/.env .

EXPOSE 8000

ENTRYPOINT [ "/go_app" ]


# dev env; live reload enabled
FROM golang:1.22.1-alpine as dev

WORKDIR /go/flashcards/

RUN apk update \
&&  apk add --no-cache bash gcc musl-dev \
&&  rm -rf /var/lib/apt/lists/* \
&&  export GO111MODULE=on

COPY ./src/ .

RUN go mod tidy

EXPOSE 8000
