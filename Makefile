MAKEFLAGS += --no-print-directory

.PHONY: esbuild dev app
.SILENT:   

ifeq ($(prod),up)
all: prod-up
else ifeq ($(prod),down)
all: prod-down
else ifeq ($(dev),up)
all: dev-up
else ifeq ($(dev),down)
all: dev-down
else ifeq ($(dev),stop)
all: dev-stop
endif  

dev-up:
	npm install --prefix web/app
#docker compose -f deployments/dev/compose.yml up -d     

dev-down:
	rm -rf web/static/js
	rm -rf web/app/node_modules

clean: 
	rm -rf web/static/js
	rm -rf web/app/node_modules

esbuild:
	npm run dev	--prefix web/app

app:
	wgo run -file .html cmd/main.go

dev:
	make -j2 esbuild app

