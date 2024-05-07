FROM golang:latest
 
 
ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/EDDYCJY/go-gin-example
COPY . $GOPATH/src/github.com/EDDYCJY/go-gin-example
RUN go build .
 
 
EXPOSE 8000
ENTRYPOINT ["./go-gin-example"]

#
# Build
# docker build -t go-gin-example .
#
# Run
# docker run -p 8000:8000 go-gin-example
