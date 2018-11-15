VERSION=`date -u +%Y.%m.%d.%H%M%S`

LDFLAGS=-ldflags "-X main.Version=${VERSION}"

clean:
	rm novus || true

rebuildstd:
	rm novus_std.go || true
	~/go/bin/go-bindata -o novus_std.go novus_std

build: clean rebuildstd
	~/go/bin/go-bindata -o novus_std.go novus_std
	go build ${LDFLAGS}

run: rebuildstd
	go run ${LDFLAGS} *.go

test: rebuildstd
	go run ${LDFLAGS} *.go test.novus