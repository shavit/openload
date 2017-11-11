build:
	docker build -t itstommy/ol-inspector .

start_dev:
	docker run --rm \
		-v ${PWD}:/go/src/github.com/shavit/openload \
		-ti itstommy/ol-inspector
