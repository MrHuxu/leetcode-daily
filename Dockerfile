FROM golang:1.16.2-alpine as builder

WORKDIR /work

ENV CGO_ENABLED 0

COPY ./go.mod /work/
COPY ./go.sum /work/
COPY ./website/main.go /work/

RUN go build -o website main.go

FROM scratch

ENV GIN_MODE release

WORKDIR /work

COPY ./questions /work/questions
COPY ./website/templates /work/website/templates

COPY --from=builder /work/website /work/output/

CMD [ "./output/website" ]