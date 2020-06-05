# sudo docker build -t go-website .
# sudo docker run -dit -e WEB_ADDR=foo -e GH_WEB_TOKEN=bar -v /var/log/go-website:/go/src/app/log -p 8000:80 --name go-website go-website
FROM golang:1.13

WORKDIR /go/src/app
COPY . .
RUN mkdir /go/src/app/log

RUN go install .

CMD ["go-website"]
