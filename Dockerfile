FROM golang:1.8 as builder

WORKDIR /go/src/git.heroku.com/dodosoft-api
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build
CMD ./dodosoft-api

FROM alpine:latest  
WORKDIR /root/
COPY --from=builder /go/src/git.heroku.com/dodosoft-api/dodosoft-api .
CMD ["./dodosoft-api"]