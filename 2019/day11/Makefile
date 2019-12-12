# Project name
PROJECT_NAME=day11

#
# Makefile options
#

# State the "phony" targets
.PHONY: all clean build


all: build

build:
	@echo 'Building...'
	@cargo build --release
	@cp ./target/release/${PROJECT_NAME} ./${PROJECT_NAME}
	@cargo clean

clean:
	@echo 'Cleaning...'
	@cargo clean
	@rm -f ./${PROJECT_NAME}
