FROM golang:latest

RUN mkdir userapi
ADD Makefile userapi/Makefile
ADD src userapi/src
EXPOSE 27017
EXPOSE 3000
RUN cd userapi && make go-get && make set-path
CMD cd userapi && make build
CMD cd userapi && make run