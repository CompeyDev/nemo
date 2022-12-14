Daemon: src/lib.rs
	@echo -e "\x1b[34m[\u001b[0m\x1b[31m*\x1b[34m\x1b[34m]\u001b[0m Building daemon shared libraries..."
	@echo -e "\x1b[34m[\u001b[0m\x1b[31m#\x1b[34m\x1b[34m]\u001b[0m cargo build --release"
	@cargo build --release
	@echo -e "\x1b[34m[\u001b[0m\x1b[31m#\x1b[34m\x1b[34m]\u001b[0m mv target/release/libshared.so ./libshared.so"
	@mv target/release/libshared.so ./libshared.so