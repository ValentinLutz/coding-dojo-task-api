FROM mcr.microsoft.com/vscode/devcontainers/base:ubuntu-22.04

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# Install flyway
USER vscode
RUN wget -qO- https://repo1.maven.org/maven2/org/flywaydb/flyway-commandline/9.8.1/flyway-commandline-9.8.1-linux-x64.tar.gz \
    | tar xvz -C /home/vscode
USER root
RUN ln -s /home/vscode/flyway-9.8.1/flyway /usr/local/bin 

# Install oapi-codegen
RUN go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.12.4
