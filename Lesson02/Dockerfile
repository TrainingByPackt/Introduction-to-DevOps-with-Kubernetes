FROM ubuntu:18.10

RUN apt-get update
RUN apt-get install -y nodejs npm
RUN npm install -g http-server

WORKDIR /usr/apps/hello-world/

CMD ["http-server", "-p", "8080"]