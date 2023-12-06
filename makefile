BUILD_OUT=pyg
SHARE_VER=0.0.0
SHARE_OUT=libpyg.$(SHARE_VER).go.so

.PHONEY: all
all: clean build shared

clean:
	rm $(BUILD_OUT) $(SHARE_OUT)

build:
	go build -o $(BUILD_OUT) .

shared:
	go build -o $(SHARE_OUT) -buildmode=c-shared -linkshared .
