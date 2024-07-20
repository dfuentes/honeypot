FROM alpine
LABEL org.opencontainers.image.authors="daniel.richard.fuentes@gmail.com"

COPY ./honeypot /
ENV PORT=8080
EXPOSE 8080/tcp
ENTRYPOINT ["/honeypot"]
