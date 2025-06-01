It's a demonstration of gRPC in golang.

### Steps
- Clone the repository.
- Go to repository.
- Run ```go run main.go```

### Note
- Proto file is in shop directory.
- By default, this will run the server and client only once. If you want the server to keep on running, just change the value of wg.Add(1) from 1 to 2 in main.go file.