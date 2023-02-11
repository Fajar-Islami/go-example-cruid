# Start by building the application.
FROM golang:1.19.2 as build
LABEL stage=dockerbuilder
WORKDIR /app/example
COPY . .

# Build the binary
RUN make build

# Now copy it into our base image.
FROM alpine:3.9

# Copy configs files
# COPY resources/config.yaml.prod /resources/config.yaml

# Copy bin file
COPY --from=build /app/example/dist/example /app/example
COPY .env /.env
# VOLUME ["/logs"]
# ARG APP_ENV

# ENV APP_ENV ${APP_ENV}
EXPOSE 8000
ENTRYPOINT ["/app/example"]