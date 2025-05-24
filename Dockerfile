
FROM ubuntu:latest
WORKDIR /app

# Needed to make external request:
RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates

COPY ./bin/linux-amd64/app /app/
RUN chmod +x /app/app

COPY web /app/web

EXPOSE 8000

CMD ["/app/app"]