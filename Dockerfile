FROM golang:1.15-alpine

WORKDIR /go/src/github.com/Studentersamfundet/er-det-torsdag
COPY main.go .
RUN go build -o app

FROM golang:1.15-alpine

WORKDIR /
COPY --from=0 /go/src/github.com/Studentersamfundet/er-det-torsdag/app .
COPY imgs ./imgs/
COPY templates ./templates/
EXPOSE 80
ENTRYPOINT ["/app"]