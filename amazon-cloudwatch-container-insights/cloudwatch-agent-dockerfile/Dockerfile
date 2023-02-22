FROM ubuntu:latest as build

# NOTE: This arg will be populated by docker buildx
# https://docs.docker.com/engine/reference/builder/#automatic-platform-args-in-the-global-scope
ARG TARGETARCH

RUN apt-get update &&  \
    apt-get install -y ca-certificates curl && \
    rm -rf /var/lib/apt/lists/*

RUN curl -O https://s3.amazonaws.com/amazoncloudwatch-agent/ubuntu/amd64/latest/amazon-cloudwatch-agent.deb && \
    dpkg -i -E amazon-cloudwatch-agent.deb

COPY config.json/ /opt/aws/amazon-cloudwatch-agent/bin/
COPY config.json/ /etc/cwagentconfig

ENTRYPOINT ["/opt/aws/amazon-cloudwatch-agent/bin/amazon-cloudwatch-agent-ctl -a fetch-config -m ec2 -s -c file:/opt/aws/amazon-cloudwatch-agent/bin/config.json"]
