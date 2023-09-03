# コンパイル用のターゲット

.PHONY: build test clean mtime-sync

build: mtime-sync

mtime-sync: cmd/mtime-sync/main.go
	go build -o mtime-sync cmd/mtime-sync/main.go

fmt:
	go fmt cmd/mtime-sync/main.go

# テスト用のターゲット
test:
	go test ./...

# cleanターゲット: 生成されたバイナリを削除
clean:
	rm -f mtime-sync

# デフォルトのターゲットを指定（単に 'make' と実行した場合の挙動）
.DEFAULT_GOAL := build
