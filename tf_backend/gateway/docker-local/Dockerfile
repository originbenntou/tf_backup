FROM golang:1.14.0

ARG GITHUB_ACCESS_TOKEN

ENV WORK_DIR /go/src/github.com/TrendFindProject/tf_backend/gateway

WORKDIR $WORK_DIR

RUN go env -w GOPRIVATE=github.com/TrendFindProject/*
RUN git config --global url."https://${GITHUB_ACCESS_TOKEN}@github.com".insteadOf "https://github.com"

ENV GO111MODULE on
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . $WORK_DIR

RUN mkdir -p /var/log/ap && touch /var/log/ap/ap.log

# HotReload for LOCAL
RUN go get github.com/pilu/fresh
CMD ["fresh", "-c", "docker-local/fresh/runner.conf"]
