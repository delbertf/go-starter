# Load environment variables from .env file
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

MAIN_PATH = tmp/bin/main

init:
	@echo "download tailwindcss standalone binary & daisy UI"
	curl -sLo tailwindcss https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64
	curl -sLo ./internal/assets/daisyui.js https://github.com/saadeghi/daisyui/releases/latest/download/daisyui.js
	curl -sLo ./internal/assets/daisyui-theme.js https://github.com/saadeghi/daisyui/releases/latest/download/daisyui-theme.js
	chmod +x tailwindcss
	@echo "install go-templ"
	go install github.com/a-h/templ/cmd/templ@latest
	go mod tidy
	

tailwind:
	@echo "build and watch the assets"
	@./tailwindcss -i ./internal/assets/input.css -o ./internal/assets/dist/styles.css --watch

templ:
	@echo "build and watch the templ files"
	@templ generate --watch --proxy="http://localhost$(HTTP_LISTEN_ADDR)" --open-browser=false

server:
	@air \
	--build.cmd "go build --tags dev -o ${MAIN_PATH} ./cmd/main.go" --build.bin "${MAIN_PATH}" --build.delay "100" \
	--build.exclude_dir "node_modules" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit true \
	--screen.clear_on_rebuild true \
	--log.main_only true	

# start the application in development
dev:
	@make -j3 tailwind templ server 