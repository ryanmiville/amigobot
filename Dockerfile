FROM scratch
EXPOSE 8080
ENTRYPOINT ["/amigobot"]
COPY ./bin/ /