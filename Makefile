.PHONY: build
build:
	@rm -rf build && mkdir build && go build -o build/imageGenerator -v ./cmd
.PHONY: run
run:
	@go run cmd/main.go

.PHONEY: create_service
create_service:
	cp images.service /etc/systemd/system/images.service
	systemctl daemon-reload
	systemctl enable images
	systemctl start images

.PHONEY: stop_service
stop_service:
	systemctl stop images

.PHONEY: restart_service
restart_service:
	systemctl restart images
	
DEFAULT_GOAL := build