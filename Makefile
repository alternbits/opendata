# Compiler binary output
COMPILER_DIR := compiler
COMPILER_BIN := $(COMPILER_DIR)/bin
COMPILER_EXE := $(COMPILER_BIN)/compile

.PHONY: build run clean validate

build:
	cd $(COMPILER_DIR) && go build -o bin/compile .

run: build
	./$(COMPILER_EXE)

clean:
	rm -rf $(COMPILER_BIN)

# Validate YAML and regenerate readme.md (alias for run)
validate: run
