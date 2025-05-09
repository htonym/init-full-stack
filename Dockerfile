
FROM ubuntu:latest
WORKDIR /app

COPY ./bin/linux-amd64/app /app/
RUN chmod +x /app/app

COPY web /app/web

EXPOSE 8000

CMD ["/app/app"]