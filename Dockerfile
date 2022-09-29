FROM golang:1.18-alpine

# vars
WORKDIR /app

# files
COPY . /home/

# dependencies & compile
RUN cd /home && go mod download && go build -o /app ./...

# Clean up
RUN rm -rf /home/*

EXPOSE 8080
CMD [ "/app/inventory" ]
