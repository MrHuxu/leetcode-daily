FROM golang:1.16.2-alpine as builder

RUN apk add make

WORKDIR /work

ENV CGO_ENABLED 0

COPY . /work/

ENV GOPROXY https://goproxy.io,direct
RUN make build

FROM scratch

ENV GIN_MODE release

WORKDIR /work

COPY ./questions /work/questions
COPY ./website/templates /work/website/templates

COPY --from=builder /work/output /work/output

CMD [ "./output/website" ]