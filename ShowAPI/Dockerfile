FROM golang:latest

RUN mkdir showapi
ADD Makefile showapi/Makefile
ADD src showapi/src
EXPOSE 27017
EXPOSE 3000
RUN cd showapi && make go-get && make set-path
#RUN make go-get
#RUN make set-path
CMD cd showapi && make build
CMD cd showapi && make run