FROM golang:1.14.0 as builder

ARG GITHUB_ACCESS_TOKEN

ENV WORK_DIR /go/src/github.com/TrendFindProject/tf_backend/account

WORKDIR $WORK_DIR

RUN go env -w GOPRIVATE=github.com/TrendFindProject/*
RUN git config --global url."https://${GITHUB_ACCESS_TOKEN}@github.com".insteadOf "https://github.com"

ENV GO111MODULE on
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . $WORK_DIR

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go install $WORK_DIR

FROM alpine:latest
ENV TZ Asia/Tokyo
RUN mkdir -p /var/log/ap && touch /var/log/ap/ap.log
COPY --from=builder /go/bin/account /go/bin/account

ARG JWT_SECRET_KEY
ARG SEND_GRID_API_KEY

ENV JWT_SECRET_KEY ${JWT_SECRET_KEY}
ENV SEND_GRID_API_KEY ${SEND_GRID_API_KEY}

ENTRYPOINT ["/go/bin/account"]
