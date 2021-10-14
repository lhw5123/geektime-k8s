FROM ubuntu

RUN apt-get update && apt-get install
ADD bin/amd64/httpserver /httpserver
EXPOSE 8080
ENTRYPOINT /httpserver
