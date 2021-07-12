FROM ubuntu:latest

ENV TZ=Europe/Warsaw

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone && \
    apt update && apt install -y \
    golang ca-certificates

EXPOSE 8086

WORKDIR /home/server

COPY . .

RUN go get github.com/gorilla/mux

CMD ["go", "run", "main.go"]