# Payload binary output filename
PAYLOAD_NAME ?= payload_gen

Payload: payload.go
	@echo -e "\x1b[34m[\u001b[0m\x1b[31m*\x1b[34m\x1b[34m]\u001b[0m Building payload binary ${PAYLOAD_NAME}..."
	@echo -e "\x1b[34m[\u001b[0m\x1b[31m#\x1b[34m\x1b[34m]\u001b[0m go build -o ${PAYLOAD_NAME} -compiler gccgo payload.go"
	@go build -x -v -o ${PAYLOAD_NAME} -compiler gccgo payload.go