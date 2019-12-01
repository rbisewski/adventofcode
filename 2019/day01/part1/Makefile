# Project name
PROJECT_NAME=day01

# Version
VERSION = `date +%y.%m`

# If unable to grab the version, default to N/A
ifndef VERSION
    VERSION = "n/a"
endif

# Flags
CFLAGS = -std=c++17 \
         -O2 \
         -fpic \
         -Wall \
         -Wextra \
         -Wpedantic \
         -Wno-missing-braces \
         -Wformat=2 \
         -Wformat-signedness \
         -Wnull-dereference \
         -Winit-self \
         -Wmissing-include-dirs \
         -Wshift-overflow=2 \
         -Wswitch-default \
         -Wswitch-enum \
         -Wunused-const-variable=2 \
         -Wuninitialized \
         -Wunknown-pragmas \
         -Wstrict-overflow=5 \
         -Warray-bounds=2 \
         -Wduplicated-cond \
         -Wfloat-equal \
         -Wundef \
         -Wshadow \
         -Wcast-qual \
         -Wcast-align \
         -Wwrite-strings \
         -Wconversion \
         -Wlogical-op \
         -Waggregate-return \
         -Wcast-align \
         -Wmissing-declarations \
         -Wpacked \
         -Wredundant-decls \
         -Winline \
         -Winvalid-pch \
         -Wstack-protector

#
# Makefile options
#


# State the "phony" targets
.PHONY: all clean build


all: clean build

build:
	@echo 'Building...'
	@g++ main.cpp ${CFLAGS} -o ${PROJECT_NAME}

clean:
	@echo 'Cleaning...'
	@rm -f ${PROJECT_NAME}
