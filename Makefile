.PHONY: build

IMAGE := tensorflow-serving-grpc

build: build-image build-api

clean:
	-docker rmi -f $(IMAGE)

build-image:
	docker build -t $(IMAGE) .

build-api: 
	CONTAINER_ID=$$(docker run -d $(IMAGE)); \
	docker wait $${CONTAINER_ID} ; \
	docker cp $${CONTAINER_ID}:/usr/src/vendor ./; \
	docker rm -f $${CONTAINER_ID};

interactive: build-image 
	docker run -it --rm $(IMAGE) bash

master: build-api
	rm -rf go && mv vendor/github.com/netbrain/tf-grpc/go .
