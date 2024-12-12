# Start from golang base image
FROM golang:1.23.1-alpine3.20 AS build-stage

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add bash && apk add build-base

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

RUN --mount=type=cache,target="/go/pkg/mod" \
    CGO_ENABLED=0 go build -tags viper_bind_struct -o /build/app ./cmd/app.go

FROM gcr.io/distroless/base-debian11 AS build-release-stage

COPY --from=build-stage /build/app /build/app

EXPOSE 8078

# Run the executable
CMD ["/build/app"]
