.PHONY: build run watch build-image deploy

username = ghost023
image = jcli
version = latest

build:
	@go build -o ./build/jcli

run: build
	@go run main.go

watch:
	@CompileDaemon -build="make build -B" -command="./build/jcli" -color=true

build-image:
	docker build -t $(username)/$(image):$(version) .

deploy:
	docker push $(username)/$(image):$(version)
