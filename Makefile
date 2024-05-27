build:
	@go build -o ./build/jcli

run:build
	@go run main.go

watch:
	@CompileDaemon -build="make build -B" -command="./build/jcli" -color=true

build:image
	docker build -t jcli .

deploy:
	docker push jcli
