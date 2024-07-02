#!/bin/sh
set -e

/varsoloki/apply_env.sh /etc/loki/minio-config.yaml

/usr/bin/loki $@
