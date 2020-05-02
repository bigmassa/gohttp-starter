FROM golang:alpine as builder

WORKDIR /go/src/app
COPY . .

RUN build/package/build.sh

FROM alpine

WORKDIR /go/src/app
COPY --from=builder /go/src/app /go/src/app

ENV PATH="/go/src/app:${PATH}"

ENTRYPOINT serve
