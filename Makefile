run:
	go run .

tidy:
	go mod tidy -v

utest:
	go test ./test -v

swaggen:
	swag init