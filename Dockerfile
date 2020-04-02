FROM golang:1.13-buster

COPY . /hello
WORKDIR /hello
RUN go build -o hello

FROM debian:buster-slim
COPY --from=0 /hello/hello /app/

## Cannot use the --chown option of COPY because it's not supported by
## Docker Hub's automated builds :/
WORKDIR /app
RUN chown -R www-data:www-data *

ENTRYPOINT ["/app/hello"]
EXPOSE 8000
USER www-data
