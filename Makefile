build: 
	@go build -o bin/flow-scraper-bridge

run: build
	@./bin/flow-scraper-bridge
