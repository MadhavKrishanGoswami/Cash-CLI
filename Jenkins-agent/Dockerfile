FROM jenkins/inbound-agent:4.13-1

# Install necessary packages including Docker and Go
USER root
RUN apt-get update && apt-get install -y curl bash docker.io

# Install the latest Go version
RUN curl -L https://go.dev/dl/go1.23.1.linux-amd64.tar.gz -o go.tar.gz && \
  tar -C /usr/local -xzf go.tar.gz && \
  rm go.tar.gz && \
  ln -s /usr/local/go/bin/go /usr/bin/go && \
  ln -s /usr/local/go/bin/gofmt /usr/bin/gofmt

# Set up Go environment variables
ENV GOPATH=/go
ENV PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

# Add the jenkins user to the docker group for permission
RUN usermod -aG docker jenkins

# Switch back to the jenkins user
USER jenkins

