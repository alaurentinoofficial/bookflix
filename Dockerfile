############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR /go/src/ng-book-service/
COPY . .

# Fetch dependencies.
# Using go get.
RUN go get -d -v

# Build the binary.
RUN CGO_ENABLED=0 go build -o /run


############################
# STEP 2 deploy final image
############################
FROM scratch

# Copy our static executable.
COPY --from=builder /run /run
COPY --from=builder /go/src/ng-book-service/.env /.env
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Run the hello binary.
CMD ["/run"]