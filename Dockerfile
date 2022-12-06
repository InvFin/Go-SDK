# Build Stage
FROM lacion/alpine-golang-buildimage:1.13 AS build-stage

LABEL app="build-InvFinSDK"
LABEL REPO="https://github.com/InvFin/Go-SDK"

ENV PROJPATH=/go/src/github.com/InvFin/Go-SDK

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/InvFin/Go-SDK
WORKDIR /go/src/github.com/InvFin/Go-SDK

RUN make build-alpine

# Final Stage
FROM lacion/alpine-base-image:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/InvFin/Go-SDK"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/InvFinSDK/bin

WORKDIR /opt/InvFinSDK/bin

COPY --from=build-stage /go/src/github.com/InvFin/Go-SDK/bin/InvFinSDK /opt/InvFinSDK/bin/
RUN chmod +x /opt/InvFinSDK/bin/InvFinSDK

# Create appuser
RUN adduser -D -g '' InvFinSDK
USER InvFinSDK

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/InvFinSDK/bin/InvFinSDK"]
