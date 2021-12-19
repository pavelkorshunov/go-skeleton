# Типа выбор компилятора
GO=go

install: build

build:
	$(GO) build -o application cmd/app/main.go