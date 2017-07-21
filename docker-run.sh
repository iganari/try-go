#!/bin/bash

set -x

# ホストのディレクトリをマウントし、直前に作成したimageから立ち上げる
BASEPATH=$(cd `dirname $0`; pwd)
REPOPATH=$(cd ../../`dirname $0`; pwd)
# echo ${REPOPATH}
docker run --rm -it -v ${BASEPATH}/src:/go/src -w /go/src ${1} /bin/bash
