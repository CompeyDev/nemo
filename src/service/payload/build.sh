echo "[*] Building payload binary $1..."
gccgo -c payload.go
gccgo -o $1 payload.o