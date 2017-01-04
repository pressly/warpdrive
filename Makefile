LDFLAGS+=-X github.com/pressly/warpdrive.VERSION=$$(git describe --tags --abbrev=0 --always)
LDFLAGS+=-X github.com/pressly/warpdrive.LONGVERSION=$$(git describe --tags --long --always --dirty)

OS=darwin linux windows
ARCH=386 amd64

##
## Tools
##

tools:
	go get -u github.com/pressly/fresh
	go get -u github.com/pressly/goose
	go get -u github.com/pressly/sup
	go get -u github.com/kardianos/govendor
	go get -u github.com/jstemmer/gotags
	go get -u golang.org/x/tools/cmd/goimports
	go get -u golang.org/x/mobile/cmd/gomobile

##
## Dependency mgmt
## only use this if you want to resync new vendor package
## if you don't need to build for appengine, go to vendor/vendor.json
## and add `appengine` into `ignore` field. for example: `"ignore": "test appengine",` 
##

vendor-clean:
	@rm -rf vendor
	@govendor init

vendor-rebuild:
	@govendor add +e
	@govendor update +external

##
## Building Server
##

clean-server:
	@rm -rf ./bin/warpdrive
	@mkdir -p ./bin/warpdrive

build-server: clean-server
	GOGC=off go build -i -gcflags="-e" -ldflags "$(LDFLAGS)" -o ./bin/warpdrive ./cmd/warpdrive

build-all-servers: clean-server
	for GOOS in $(OS); do 																											\
		for GOARCH in $(ARCH); do 																								\
			echo "building $$GOOS $$GOARCH ..."; 																		\
			export GOGC=off;																												\
			export GOOS=$$GOOS; 																										\
			export GOARCH=$$GOARCH; 																								\
			go build -ldflags "$(LDFLAGS)" 																					\
			-o ./bin/warpdrive/warpdrive-$$GOOS-$$GOARCH ./cmd/warpdrive;       		\
		done 																																			\
	done

##
## Building cli
##
clean-cli:
	@rm -rf ./bin/warp
	@mkdir -p ./bin

build-cli: clean-cli
	GOGC=off go build -i -gcflags="-e" -ldflags "$(LDFLAGS)" -o ./bin/warp ./cmd/warp

##
## Building WarpDrive's Clients, ios and android
##

clean-ios:
	@rm -rf ./client/ios/Warpify.framework

build-ios: clean-ios
	@cd ./cmd/warpify && gomobile bind -target=ios -ldflags="-s -w" . && mv -f Warpify.framework ../../client/ios

clean-android:
	@rm -rf client/android/lib/warpify.aar

build-android: clean-android
	@cd ./cmd/warpify && gomobile bind -target=android -ldflags="-s -w" . && mv -f warpify.aar ../../client/android/lib

build-clients: build-ios build-android

##
## Building everything
##

build: build-server build-cli build-clients

##
## Database
##

db-create:
	@goose up

db-destroy:
	@goose down

db-reset: db-destroy db-create

##
## Development
##

print-api:
	@go run ./cmd/doc/main.go

build-dev-folder:	
	@mkdir -p ./bin/tmp/warpdrive

run: build-dev-folder
	fresh -c ./etc/fresh-runner.conf -p ./cmd/warpdrive -r '-config=./etc/warpdrive.conf' -o ./bin/warpdrive

kill:
	@lsof -t -i:8221 | xargs kill

clean-ios-example:
	@rm -rf ./client/examples/Sample1/node_modules/react-native-warpdrive/ios/Warpify.framework

build-ios-example: clean-ios-example
	@cd ./cmd/warpify && gomobile bind -target=ios . && mv -f Warpify.framework ../../client/examples/Sample1/node_modules/react-native-warpdrive/ios

clean-android-example:
	@rm -rf ./client/examples/Sample1/android/warpify/warpify.aar

build-android-example: clean-android-example
	@cd ./cmd/warpify && gomobile bind -target=android . && mv -f warpify.aar ../../client/examples/Sample1/android/warpify


clean-cli-dev:
	@rm -rf ./client/examples/Sample1/warp

build-cli-dev: clean-cli-dev
	GOGC=off go build -i -gcflags="-e" -ldflags "$(LDFLAGS)" -o ./client/examples/Sample1/warp ./cmd/warp
