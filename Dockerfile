FROM alpine

ADD build/ ./
ENV PORT 8080
EXPOSE 8080
ENTRYPOINT ["/govolutto.zip"]
