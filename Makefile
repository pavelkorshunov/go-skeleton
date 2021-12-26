# Типа выбор компилятора
GO=go

install: build

build:
	$(GO) build -o app -ldflags="-X 'github.com/pavelkorshunov/go-skeleton/internal/app.Mode=production'" cmd/app/main.go

build_dev:
	$(GO) build -o app cmd/app/main.go

run:
	$(GO) run cmd/app/main.go