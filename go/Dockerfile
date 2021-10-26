#
FROM golang AS builder

WORKDIR /app

COPY . .

RUN go get -d
RUN CGO_ENABLED=0 go build server.go


#
FROM alpine

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080
CMD ["/app/server"]