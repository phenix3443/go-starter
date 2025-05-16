ARG VARIANT="1-1.23-bookworm"
FROM mcr.microsoft.com/devcontainers/go:${VARIANT}


# This Dockerfile adds a non-root user with sudo access. Update the “remoteUser” property in
# devcontainer.json to use it. More info: https://aka.ms/vscode-remote/containers/non-root-user.
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

RUN apt-get update

# Install common packages and setup non-root user. Use a separate RUN statement to add your own dependencies.
ARG INSTALL_ZSH="true"
ARG UPGRADE_PACKAGES="true"
ARG COMMON_SCRIPT_SOURCE="https://raw.githubusercontent.com/microsoft/vscode-dev-containers/main/script-library/common-debian.sh"

RUN export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends curl ca-certificates build-essential iputils-ping bash-completion shfmt 2>&1 \
    && curl -sSL ${COMMON_SCRIPT_SOURCE} -o /tmp/common-setup.sh \
    && /bin/bash /tmp/common-setup.sh "${INSTALL_ZSH}" "${USERNAME}" "${USER_UID}" "${USER_GID}" "${UPGRADE_PACKAGES}"

# Install Node.js for use with web applications
ARG NODE_VERSION="20"
ARG NODE_SCRIPT_SOURCE="https://raw.githubusercontent.com/microsoft/vscode-dev-containers/main/script-library/node-debian.sh"
ENV NVM_DIR=/usr/local/share/nvm
ENV NVM_SYMLINK_CURRENT=true
ENV PATH=${NVM_DIR}/current/bin:${PATH}
RUN curl -sSL ${NODE_SCRIPT_SOURCE} -o /tmp/node-setup.sh \
    && /bin/bash /tmp/node-setup.sh "${NVM_DIR}" "${NODE_VERSION}" "${USERNAME}" \
    && npm install -g cspell@latest

# Install database tools
RUN apt-get install -y postgresql-client mysql-client redis-tools

# Install go dependencies
RUN go install github.com/goreleaser/goreleaser/v2@latest \
    && go install golang.org/x/tools/cmd/goimports@latest \
    && go install golang.org/x/tools/gopls@latest \
    && go install github.com/spf13/cobra-cli@latest \
    && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Install scripts
COPY script /usr/local/bin/
RUN  bash /usr/local/bin/git-completion.sh

# Clean up
RUN apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/* /tmp/common-setup.sh /tmp/node-setup.sh

USER ${USERNAME}
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 CMD [ "bash", "-c", "echo ok" ]