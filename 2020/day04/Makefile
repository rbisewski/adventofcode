PROJECT_NAME = 'day04'

# Version
VERSION = `date +%y.%m`

# If unable to grab the version, default to N/A
ifndef VERSION
    VERSION = "n/a"
endif

#
# Makefile options
#


# State the "phony" targets
.PHONY: all debug build run clean

all: build

debug:
	@echo 'Building debug binary of ${PROJECT_NAME}...'
	@cargo build

build:
	@echo 'Building release for ${PROJECT_NAME}...'
	@cargo build --release

static:
	@echo 'Building static binaries for ${PROJECT_NAME}...'
	@rustup target add x86_64-unknown-linux-musl
	@cargo build --release --target x86_64-unknown-linux-musl

run:
	@RUST_LOG=trace cargo run

clean:
	@echo 'Cleaning...'
	@cargo clean
