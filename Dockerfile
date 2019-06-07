FROM golang

ARG BRANCH=r1.13
ARG PROTOC_VERSION=3.8.0
ARG PROTOC_GO_VERSION=v1.2.0
WORKDIR /usr/src

RUN apt-get -y update && apt-get -y install git curl unzip && \
	git clone --depth 1 -b $BRANCH https://github.com/tensorflow/serving.git && \
	git clone --depth 1 -b $BRANCH https://github.com/tensorflow/tensorflow.git && \
	curl -o /tmp/protoc.zip -SL https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOC_VERSION/protoc-$PROTOC_VERSION-linux-x86_64.zip && \
	unzip -d /usr/local /tmp/protoc.zip && \
	go get -d -u github.com/golang/protobuf/protoc-gen-go && \
	git -C "$(go env GOPATH)"/src/github.com/golang/protobuf checkout $PROTOC_GO_VERSION && \
	go install github.com/golang/protobuf/protoc-gen-go && \
	mkdir vendor

ENV PROTOC_OPTS="-I tensorflow -I serving --go_out=plugins=grpc:vendor"

# Insert missing go_package
RUN find serving/tensorflow_serving | \
 	grep \\.proto$ | \
 	xargs -L1 -I@ sh -c 'echo "option go_package = \"github.com/tensorflow/serving/tensorflow_serving\";" >> @'

# Insert missing go_package but dont include compiler/contrib
RUN find ./tensorflow | grep \\.proto$ | \
	grep -v tensorflow/compiler | \
	grep -v tensorflow/contrib | \
	xargs grep -L "option go_package" | \
	xargs -L1 -I@ sh -c 'dirname @ | \
	sed "s/.\/tensorflow\/tensorflow/github.com\/tensorflow\/tensorflow\/tensorflow\/go/g" | \
	xargs -I% sh -c "echo option go_package = \\\"%\\\"\\;" >> @'

# Rename packages
RUN find serving/ tensorflow/ -type f | \
	grep \\.proto$ | \
	xargs -L1 sed -i "s/option go_package = \"github.com\/tensorflow/option go_package = \"github.com\/netbrain\/tf-grpc\/go/g"

# Compile proto files
CMD find serving/ tensorflow/ -type f | \
	grep \\.proto$ | \
	xargs grep -l "option go_package" | \
	xargs -I@ sh -c "protoc $PROTOC_OPTS @"
