
#use official golang image
FROM golang:1.23.4

#set working directory
WORKDIR /app

#copy the source code to working directory
COPY . .

#build the source code
RUN go build -o main .

#expose the port
EXPOSE 8080

#run the executable
CMD ["./main"]