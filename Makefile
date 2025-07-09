VERSION := 0.1.2

build:
	docker build -t sreng1neer/norrdns:$(VERSION) .

push: build
	docker push sreng1neer/norrdns:$(VERSION)