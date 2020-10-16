# To listen on local port 8000 (change "8000" in the -p option to change the port):
#  docker build -t go-website .
#  docker run -dit -e WEB_ADDR=0.0.0.0:80 -v /var/log/go-website:/go/src/go-website/log -p 127.0.0.1:8000:80/tcp --name go-website go-website

# Remove "127.0.0.1:" from the -p option to fully expose the port (don't do if you're using a reverse-proxy).
# Change the path before the color in the -v option to change the log file location.

FROM golang:1.15

WORKDIR /go/src/go-website
COPY . .
RUN mkdir /go/src/go-website/log

RUN go install .

CMD ["go-website"]
