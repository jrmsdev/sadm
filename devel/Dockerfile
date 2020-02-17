FROM golang:1.13-alpine

LABEL maintainer="Jeremías Casteglione <jrmsdev@gmail.com>"
LABEL version="2020.02.16"

USER root:root
WORKDIR /root

ARG SADM_UID
ARG SADM_GID
ARG SADM_UMASK

RUN addgroup -g ${SADM_GID} sadm && \
	adduser -D -g sadm -h /home/sadm -G sadm -s /bin/ash -u ${SADM_UID} sadm && \
	chmod -v 0750 /home/sadm

ENV GOPATH /go
ENV SADM_SRC ${GOPATH}/src/sadm

RUN rm -rf ${GOPATH} && mkdir -vp ${SADM_SRC} && \
	chown -vR sadm:sadm ${GOPATH} && chmod -vR 0750 ${GOPATH}

USER sadm:sadm
WORKDIR /home/sadm

ENV USER sadm
ENV HOME /home/sadm

RUN echo "umask ${SADM_UMASK}" >>/home/sadm/.profile
RUN echo "export PATH=${PATH}" >>/home/sadm/.profile
RUN echo "export PS1='\[\e]0;\h \w\a\]\w\$ '" >>/home/sadm/.profile

RUN (go env | sort) && (echo ""; go version; echo "")

WORKDIR ${SADM_SRC}
ENTRYPOINT /bin/ash -l -i
