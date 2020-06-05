# sudo docker build -t go-website .
# sudo docker run -dit -e GH_WEB_TOKEN=foo -v /var/log/go-website:/go/src/app/log -p 8000:80 --name go-website go-website
FROM golang:1.13

WORKDIR /go/src/app
COPY . .

RUN go install .

RUN mkdir /mudcord
WORKDIR /mudcord

CMD ["app"]
