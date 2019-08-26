FROM debian

ARG BRANCH=r1.13
ARG PROTOC_VERSION=3.8.0
ARG PROTOC_JAVA_VERSION=1.23.0

WORKDIR /usr/src

RUN apt-get -y update && apt-get -y install git curl unzip && \
	git clone --depth 1 -b $BRANCH https://github.com/tensorflow/serving.git && \
	git clone --depth 1 -b $BRANCH https://github.com/tensorflow/tensorflow.git && \
	curl -o /tmp/protoc.zip -SL https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOC_VERSION/protoc-$PROTOC_VERSION-linux-x86_64.zip && \
	unzip -d /usr/local /tmp/protoc.zip && \
	curl -o /tmp/protoc-gen-grpc-java -SL https://repo1.maven.org/maven2/io/grpc/protoc-gen-grpc-java/$PROTOC_JAVA_VERSION/protoc-gen-grpc-java-$PROTOC_JAVA_VERSION-linux-x86_64.exe && \
	chmod +x /tmp/protoc-gen-grpc-java && \
	mkdir vendor

ENV PROTOC_OPTS="-I tensorflow -I serving --plugin=protoc-gen-grpc-java=/tmp/protoc-gen-grpc-java --grpc-java_out=vendor --java_out=vendor"

# Compile proto files
CMD find serving/ tensorflow/ -type f | \
	grep \\.proto$ | \
	grep -v tensorflow/compiler | \
	grep -v tensorflow/contrib | \
	xargs -I@ sh -c "protoc $PROTOC_OPTS @"
