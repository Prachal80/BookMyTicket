FROM golang:latest

RUN mkdir MovieApi
ADD Makefile MovieApi/Makefile
ADD src MovieApi/src
EXPOSE 27017
EXPOSE 3000
RUN cd MovieApi && make go-get && make set-path
CMD cd MovieApi && make go-build
CMD cd MovieApi && make go-run
