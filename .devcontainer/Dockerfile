FROM golang:1-trixie

ARG USER=dev

RUN useradd --create-home --shell /bin/bash ${USER}

RUN apt-get install -y curl && \
    curl -fsSL https://deb.nodesource.com/setup_26.x | bash - && \
    apt-get update && \
    apt-get install -y nodejs

USER ${USER}