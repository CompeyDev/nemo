# Rules to build all binaries & services.

all:
	@cd src/service/payload/ && make Payload PAYLOAD_NAME="payload_test"
	@cd src/ && make Server
	@cd src/service/client && make Client
	@cd src/service/daemon && make Daemon