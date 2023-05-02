BIN    := bin
PREFIX := /usr/local

CMD := subanner
PRJ := $(CMD)
OUT := $(BIN)/$(CMD)

FONT     := font_8x8/font_8x8.go
FONT_SRC := font_8x8/font_8x8_src.go

.PHONY: all build run install uninstall clean

all: build

build: $(OUT)

$(FONT): $(FONT_SRC)
	go run $(FONT_SRC) > $(FONT)

go.mod:
	go mod init $(PRJ)

$(OUT): $(FONT) go.mod
	mkdir -p $(BIN)
	go build -o $(BIN) $(PRJ)/cmd/$(CMD)/

run:
	go run cmd/$(CMD)/$(CMD).go

install: $(OUT)
	cp $(OUT) $(PREFFIX)/$(BIN)/

uninstall:
	rm -f $(PREFFIX)/$(BIN)/$(CMD)

clean:
	rm -f $(FONT)
	rm -f $(OUT)
	rm -rf $(BIN)

