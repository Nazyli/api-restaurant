FROM golang:1.13 as builder

WORKDIR /app


COPY go.mod .
COPY go.sum .
RUN go mod download
# Copy the code into the container
COPY . .
# Statically compile our app for use in a distroless container
# RUN CGO_ENABLED=0 go build -ldflags="-w -s" -v -o app/cmd/ .
# Build the Go app
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -mod=readonly -a -installsuffix cgo -o app ./cmd/


# A distroless container image with some basics like SSL certificates
# https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/static

COPY --from=builder /app/app /app
COPY .env.sample .env
# COPY file file

ENTRYPOINT ["/app"]