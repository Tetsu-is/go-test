# Run go run main.go
run :
	go run main.go

# run test
test :

compile :
	goimports -w -l . 
	go vet 
	go build -o main main.go

review :
	golangci-lint run