FROM golang:latest

RUN mkdir mailapi
ADD Makefile mailapi/Makefile
ADD src mailapi/src
EXPOSE 27017
EXPOSE 3000
RUN cd mailapi && make go-get && make set-path
CMD cd mailapi && make build
CMD cd mailapi && make run