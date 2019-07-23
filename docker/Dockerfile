FROM node:11.3.0-stretch-slim as node_image

FROM python:3.7.2-alpine3.8 as python_image
RUN pip install mkdocs

FROM golang:1.12.7-stretch

ARG PROJECT
ARG USER_ID
#ARG SSH_PRIVATE_KEY

ENV BUILD_DIR /builder/src/gitlab.picus.io/picussecurity/$PROJECT
ENV BUILD_SCRIPT "$BUILD_DIR"/docker/build.sh
ENV DEP_RELEASE_TAG=v0.5.0

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh - \
    && apt-get -y update && apt-get -y install libc6 jq

# https://smartystreets.com/blog/2018/09/private-dependencies-in-docker-and-go
# alternative way
#RUN mkdir -p /builder/.ssh && umask 0077 && echo "${SSH_PRIVATE_KEY}" > /builder/.ssh/id_rsa \
#	&& git config --global url."git@gitlab.picus.io:".insteadOf https://gitlab.picus.io/ \
#	&& ssh-keyscan gitlab.picus.io >> /builder/.ssh/known_hosts

WORKDIR "$BUILD_DIR"
COPY . "$BUILD_DIR"

# takes node and yarn etc. from intermediate image
COPY --from=node_image / /
COPY --from=python_image / /

VOLUME "$BUILD_DIR"

RUN adduser -h /builder/ -D -u "$USER_ID" jenkins \
    && chown jenkins -R /builder \
    && chmod +x "$BUILD_SCRIPT"

USER jenkins

ENTRYPOINT "$BUILD_SCRIPT"
