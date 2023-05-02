BIN    := bin
PREFIX := /usr/local

CMD := subanner
PRJ := $(CMD)
OUT := $(BIN)/$(CMD)

FONT    := pkg/font8/font8.go
FONTSRC := pkg/font8src/font8src.go pkg/font8src/font8make.go
MKFONT  := mkfont/mkfont.go

.PHONY: all mod font sysv build run clean install uninstall

all: build

mod: go.mod

font: $(FONT)

sysv: pkg/font8sysv/font8sysv.go pkg/font8sysv/font8make.go
	go run mkfont/mkfont-sysv.go > $(FONT)

pkg/font8sysv/font8sysv.go:
	bash pkg/font8sysv/grab_banner.sh > pkg/font8sysv/font8sysv.go

build: $(OUT)

run: $(FONT)
	go run cmd/$(CMD)/$(CMD).go "Hello!"

clean:
	rm -f $(FONT)
	rm -f $(OUT)
	@#rm -rf $(BIN)

install: $(OUT)
	cp $(OUT) $(PREFIX)/$(BIN)/

uninstall:
	rm -f $(PREFIX)/$(BIN)/$(CMD)

go.mod:
	go mod init $(PRJ)

$(FONT): $(FONTSRC)
	go run $(MKFONT) > $(FONT)

$(OUT): $(FONT) pkg/font8/font8print.go go.mod
	mkdir -p $(BIN)
	go build -o $(BIN) $(PRJ)/cmd/$(CMD)/


