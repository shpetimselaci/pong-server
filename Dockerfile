FROM golang:1.18.0

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN apt-get update && apt-get install --yes libgl1-mesa-dev mesa-utils xorg-dev xvfb
RUN go build -v -o /usr/local/bin/app ./main.go

EXPOSE 8080
ENV DISPLAY :99

CMD ["xvfb-run", "-a", "app"]


