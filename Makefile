.PHONY: binary
binary: 
	@go build

.PHONY: clean
clean:
	@rm -rf seal *.log

