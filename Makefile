MAKEFLAGS += --no-print-directory

.PHONY: esbuild dev app
.SILENT:   

clean: 
	rm -rf app/web/static/js
	rm -rf app/js/node_modules

esbuild:
	npm install --prefix app/js
	npm run dev	--prefix app/js

app:
	wgo run -file app/web/js/index.js app/cmd/main.go 

dev:
	make -j2 esbuild app

