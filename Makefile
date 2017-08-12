
BUILD_CMD=go
BUILD_TAG=release

OUT=phantomize
APK=$(OUT).apk

build:
	$(BUILD_CMD) build -tags="$(BUILD_TAG)"

mobile:
	make build BUILD_CMD=gomobile

clean:
	rm -f $(OUT)
	rm -f $(APK)
