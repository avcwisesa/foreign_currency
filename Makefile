test:
	@./scripts/coverage.sh
	go tool cover -html=coverage.out -o coverage.html