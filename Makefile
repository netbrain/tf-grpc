.PHONY: build

IMAGE := tensorflow-serving-grpc

go: build-go-image copy-vendor
java: build-java-image copy-vendor

clean:
	rm -rf vendor && \
	docker rmi -f $(IMAGE)

build-go-image:
	docker build -t $(IMAGE) -f dockerfiles/Dockerfile.go .

build-java-image:
	docker build -t $(IMAGE) -f dockerfiles/Dockerfile.java .

copy-vendor: 
	CONTAINER_ID=$$(docker run -d $(IMAGE)); \
	docker wait $${CONTAINER_ID} ; \
	docker cp $${CONTAINER_ID}:/usr/src/vendor ./; \
	docker rm -f $${CONTAINER_ID};

interactive:
	docker run -it --rm $(IMAGE) bash
