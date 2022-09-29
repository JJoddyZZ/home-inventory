# local deploy
build:
	@go build -o bin/inventory/ ./...
run-local: build
	@./bin/inventory/inventory

# testing
test:
	@go test -v -cover -coverprofile=cover.out -race
display-cover:
	@go tool cover -html=cover.out

# container deployment
docker-build:
	docker build --tag inventory .
# --progress=plain logs output to stdout
docker-build-verbose:
	docker build --progress=plain --no-cache --tag inventory .
# -d flag will ignore STDOUT prints
run-detached-container: docker-build
	@docker run -d -p 8180:8080 --name inventory inventory
run-container: docker-build
	@docker run -p 8180:8080 --name inventory inventory
clean:
	-@docker stop inventory
	-@docker rm inventory
	-@docker rmi inventory
