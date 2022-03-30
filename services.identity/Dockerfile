FROM golang:alpine
WORKDIR /build
COPY services.identity services.identity
COPY shared shared
WORKDIR /build/services.identity
RUN go mod download
RUN go build -o main ./src/main.go
WORKDIR /dist
RUN cp /build/services.identity/main .
RUN cp /build/services.identity/src/app.env .
CMD ["/dist/main"]

# WORKDIR /usr/local/go/src/wptest
# COPY services.identity services.identity
# COPY shared shared
# WORKDIR /usr/local/go/src/wptest/services.identity
# # RUN go mod download

# CMD ["./src"]