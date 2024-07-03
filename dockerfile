FROM golang:1.22

LABEL maintainer="hbudalam, fhamza, zabulla, moadwan, aaffoune"
LABEL description="A Containerized Forum"
LABEL version="1.0"

WORKDIR /app

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD [ "./main" ]
