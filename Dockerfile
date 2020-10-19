FROM golang:alpine as builder
WORKDIR $GOPATH/app
ENV RUNTIME_USER=go
ENV RUNTIME_UID=10001

# Create runtime user
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/hohome" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid 10001 \
    "go"

# Build app
COPY . .
RUN go get -d -v
RUN CGO_ENABLED=0 go build -o /go/bin/app
RUN ls -l /go/bin

############
## Runner ##
############
FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /go/bin/app /go/bin/app

USER go:go
EXPOSE 8000
ENTRYPOINT ["/go/bin/app"]