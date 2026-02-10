FROM docker.io/library/golang:1-trixie AS server
FROM docker.io/oven/bun:1 AS client

FROM server 

ARG USERNAME=dev
ARG USER_UID=1000
ARG USER_GID=$USER_UID

RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -s /usr/bin/bash -m $USERNAME \
    && apt-get update \
    && apt-get install -y ca-certificates sudo \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME

# Install curl
RUN apt install -y curl

# Install Bun
COPY --from=client /usr/local/bin/bun /usr/local/bin/bunx /usr/local/bin/

USER $USERNAME
