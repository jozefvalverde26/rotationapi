FROM golang:1.12 
RUN mkdir /rotationapi 
ADD . /rotationapi/ 
WORKDIR /rotationapi
RUN export PATH="$GOPATH/bin:$PATH"
RUN go build cmd/rotationapi/main.go 
EXPOSE 5000
CMD ["./main"]
