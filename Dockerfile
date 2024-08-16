# Start from the official GoLang image
FROM golang:latest

RUN go install github.com/air-verse/air@latest

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

RUN ["chmod", "+x", "./init.sh"]

# Expose the port that the application will listen on
EXPOSE 8080

# Set the entry point for the container
CMD ["sh", "init.sh"]
# CMD ["./main"]
