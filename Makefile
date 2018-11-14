VERSION=`date -u +%Y.%m.%d.%H%M%S`

LDFLAGS=-ldflags ""-X main.Version=${VERSION}"

clean:
	rm novus || true

build: clean
	go build ${LDFLAGS}