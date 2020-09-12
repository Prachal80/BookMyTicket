FROM golang:latest

EXPOSE 3000 27017
RUN mkdir TheaterAPI
RUN cd TheaterAPI
ADD . .
RUN make set-path
RUN make go-get
RUN make go-build
CMD [ "/go/./theaterapi" ]
