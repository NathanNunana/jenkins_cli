build:
	@go build -o ./build/jenkins_cli

run:build
	@go run main.go

watch:
	@CompileDaemon -build="make build -B" -command="./build/orto" -color=true