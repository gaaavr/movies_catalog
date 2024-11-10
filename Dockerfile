FROM golang:latest

ENV GOPATH=/

COPY ./ ./

# build go app
RUN go build -o movies_catalog ./cmd/movies_catalog/main.go

CMD ["./movies_catalog"]
