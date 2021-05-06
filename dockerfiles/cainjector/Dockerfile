# Copyright (C) 2020, 2021, Oracle and/or its affiliates.
# Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

# builder image
FROM container-registry.oracle.com/os/oraclelinux:7.9@sha256:5aa7df08f9ab8cd6237223b0b6c5fd605f140164235b462a01e8b9d56fb03daf as cert_base

RUN yum update -y \
    && yum-config-manager --save --setopt=ol7_ociyum_config.skip_if_unavailable=true \
    && yum install -y oracle-golang-release-el7 \
    && yum-config-manager --enable ol7_developer_golang113 \
    && yum-config-manager --add-repo http://yum.oracle.com/repo/OracleLinux/OL7/developer/golang114/x86_64 \
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

WORKDIR /root/src/cainjector

RUN mkdir -p /root/src/cainjector
COPY . /root/src/cainjector/.

RUN make cainjector

FROM container-registry.oracle.com/os/oraclelinux:7-slim@sha256:fcc6f54bb01fc83319990bf5fa1b79f1dec93cbb87db3c5a8884a5a44148e7bb

# update the packages which have been flagged as having vulnerabilities in the image scan
RUN yum update -y \
    && yum install -y dbus-libs bind-license dbus expat bind-license bash libxml2 shared-mime-info python libxml2-python python-libs libcurl curl file-libs unzip nspr nss-util glib2 nss-sysinit nss-softokn-freebl openldap nss-tools cpio libcom_err libssh2 nss systemd-libs glibc glibc-common systemd \
    && yum clean all \
    && rm -rf /var/cache/yum

# the hash in this path is the hash of the output path so hard coding it here is ok as long as we don't change the output path
COPY --from=cert_base /root/src/certmanager/bazel-out/k8-fastbuild-ST-5e46445d989a02f948782b024db718465cb931c7b73759c1823f80cfbb4c6af3/bin/cmd/cainjector/cainjector_/cainjector /usr/bin/cainjector

RUN ls -la /usr/bin/cainjector

USER nobody

ENTRYPOINT ["/usr/bin/cainjector"]
