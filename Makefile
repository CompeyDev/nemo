# Rules to build all binaries & services.

all:
	@cd src/service/payload/ && make Payload PAYLOAD_NAME="payload_test"
	@cd src/ && make Server
	@cd src/service/client && make Client
	@cd src/service/daemon && make Daemon

clean:
	@echo -e "\x1b[34m[\u001b[0m\x1b[31m*\x1b[34m\x1b[34m]\u001b[0m Removing generated binaries..."
	@rm -rf server* payload* src/service/daemon/daemon*