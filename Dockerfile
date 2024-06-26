FROM alpine:latest
COPY main .
EXPOSE 8080
CMD ["/main"]
