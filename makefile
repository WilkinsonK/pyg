BUILD_OUT=pygmy
SHARE_VER=0.0.0
SHARE_OUT=libpygmy.$(SHARE_VER).go.so

.PHONEY: all
all: clean build shared

clean:
	rm $(BUILD_DIR) $(SHARE_OUT)

build:
	go build -o $(BUILD_OUT) .

shared:
	go build -o $(SHARE_OUT) -buildmode=c-shared -linkshared .
