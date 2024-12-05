# Start from golang base image
FROM golang:1.23.1-alpine3.20 AS build-stage

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add git && apk add bash && apk add build-base

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

ENV GOCACHE=/root/.cache/go-build

RUN --mount=type=cache,target="/root/.cache/go-build" \
    CGO_ENABLED=0 go build -o /build/app ./cmd/app.go

FROM gcr.io/distroless/base-debian11 AS build-release-stage

COPY --from=build-stage /build/app /build/app

EXPOSE 8078 8080

# Run the executable
CMD ["/build/app"]
