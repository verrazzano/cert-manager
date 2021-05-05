# Copyright (C) 2020, 2021, Oracle and/or its affiliates.
# Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

# builder image
FROM container-registry.oracle.com/os/oraclelinux:8@sha256:816d99f0bbe8c6cca3d2935865e23e754483d2cface2476d21376a7c656d920f as cert_base

RUN yum update -y \
    && yum-config-manager --save --setopt=ol7_ociyum_config.skip_if_unavailable=true \
    && yum install -y oracle-golang-release-el7 \
    && yum-config-manager --enable ol7_developer_golang113 \
    && yum-config-manager --add-repo http://yum.oracle.com/repo/OracleLinux/OL7/developer/golang113/x86_64 \
    && yum install -y git gcc make golang-1.13.3-1.el7.x86_64 \
    && yum -y install yum-utils unzip \
    && yum -y install which patch \
    && yum -y install java-11-openjdk-devel \
    && yum -y install gcc-c++ \
    && yum clean all \
    && go version

# Compile to /usr/bin
ENV GOBIN=/usr/bin

# Set go path
ENV GOPATH=/go

# Bazel
RUN curl -LO https://github.com/bazelbuild/bazel/releases/download/3.3.0/bazel-3.3.0-installer-linux-x86_64.sh \
    && chmod +x bazel-3.3.0-installer-linux-x86_64.sh
RUN ./bazel-3.3.0-installer-linux-x86_64.sh

WORKDIR /root/src/certmanager

RUN mkdir -p /root/src/certmanager
COPY . /root/src/certmanager/.

RUN make controller

FROM container-registry.oracle.com/os/oraclelinux:8-slim@sha256:83644dd0a3c3cdfeac738cff70300a65cfb7577eb86fac431a572126a3a65133

COPY --from=cert_base /root/src/certmanager/bazel-out/k8-fastbuild-ST-5e46445d989a02f948782b024db718465cb931c7b73759c1823f80cfbb4c6af3/bin/cmd/controller/controller_/controller /usr/bin/cert-manager

# Add License and Readme files to the image
COPY LICENSE LICENSES README.md CONTRIBUTING.md THIRD_PARTY_LICENSES.txt  /license/

RUN ls -la /usr/bin/cert-manager

USER nobody

ENTRYPOINT ["/usr/bin/cert-manager"]
