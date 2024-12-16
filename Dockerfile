# ============================================== BUILD STAGE: Building the binary of the App
FROM golang:1.22 AS build

# Build folder
WORKDIR /go/src

# Copy all the Code and stuff to compile everything
COPY . .

# Downloads all the dependencies in advance (could be left out, but it's more clear this way)
RUN go mod download


#=== FINAL STAGE: Builds the application as a static linked one, to allow it to run on alpine
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app ./main.go

# Moving the binary to the 'final Image' to make it smaller
FROM alpine:latest as release

WORKDIR /app

# `boilerplate` should be replaced here as well
COPY --from=build /go/src/app /usr/bin/app

# Add packages
RUN apk -U upgrade \
    && apk add --no-cache ca-certificates \
    && chmod +x /usr/bin/app

ENTRYPOINT ["/usr/bin/app"]