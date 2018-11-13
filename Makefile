VERSION=`git rev-parse HEAD`

LDFLAGS=-ldflags "-X github.com/rroe/Novus.Version=${VERSION}"

clean:
	rm novus || true

build: clean
	go build ${LDFLAGS}