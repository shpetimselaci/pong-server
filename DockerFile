FROM golang:1.18.0

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN apt-get update
RUN apt-get install --yes  libgl1-mesa-dev xorg-dev
RUN go build -v -o /usr/local/bin/app ./main.go

CMD ["app"]

EXPOSE 8080