ARG GO_VERSION=1.13.6
FROM golang:${GO_VERSION}-alpine AS dev

# Install Git
RUN apk add --update git

# Set Go build Env
ENV GO111MODULE="on" \
    CGO_ENABLED=0 \
    GOOS=linux

# Set Application path
ENV APP_PATH="/converse"

# Move to that directory
WORKDIR ${APP_PATH}

# Copy and cache Go Modules
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy rest of the code
COPY . .

# Now run all the tests
CMD go test ./...