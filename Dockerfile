FROM golang:latest

# hot reload & debug
RUN go install github.com/air-verse/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY . /app
WORKDIR /app

# make init.sh executable
RUN ["chmod", "+x", "init.sh"]

# Expose the port that the application will listen on
EXPOSE 8080
EXPOSE 2345

# Set the entry point for the container
CMD ["sh", "init.sh"]
