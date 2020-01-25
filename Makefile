
BUILD_CMD=go
BUILD_TAG=release

OUT=phantomize
APK=$(OUT).apk

build:
	$(BUILD_CMD) build -tags="$(BUILD_TAG)"

mobile:
	make build BUILD_CMD=gomobile

mobile-install: mobile
	adb install -r $(APK)

clean:
	rm -f $(OUT)
	rm -f $(APK)

build-on-docker:
	docker run \
		-it \
		-v $(CURDIR):/app \
		-v phantomize-go-module-tmp:/go \
		-w /app \
		pankona/gomo-simra \
		make mobile
