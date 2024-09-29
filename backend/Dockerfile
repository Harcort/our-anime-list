FROM golang:1.21 as builder
WORKDIR /build
COPY . ./
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o app .

FROM alpine:latest
RUN mkdir /.config
COPY --from=builder /usr/local/go/ /usr/local/go/
COPY --from=builder /build/config/.dev.env /env
COPY --from=builder /build/entrypoint.sh /entrypoint.sh
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/app ./
ENV PATH="/usr/local/go/bin:${PATH}"

ENTRYPOINT ["/entrypoint.sh"]
