#!/bin/bash
set -e
DIR="$(dirname $(realpath $0))"
DIST=${DIR}/../dist

if [ $1 == "all" ]; then
    oss=(linux)
    archs=(amd64)
else
    eval $(go tool dist env)
    oss=($GOOS)
    archs=($GOARCH)
fi

cmds=(gitlab-upgrade-artifact)

set -x

# See https://docs.gitlab.com/ee/ci/variables/predefined_variables.html
# CI_COMMIT_TIMESTAMP ???
BUILT=$(date +"%Y-%m-%dT%H:%M:%S")
BRANCH="$CI_COMMIT_BRANCH"
if [ -z "$BRANCH" ]; then
    BRANCH=$(git rev-parse --abbrev-ref HEAD)
fi
COMMIT="$CI_COMMIT_SHA"
if [ -z "$COMMIT" ]; then
  COMMIT=$(git rev-parse HEAD)
fi


mkdir -p ${DIST}
echo "$BUILT $HEAD" > ${DIST}/version.txt

mkdir -p internal/httpd/admin_webroot
mkdir -p internal/httpd/dl_webroot

for os in ${oss[@]}
do
    for arch in ${archs[@]}
    do
        output_dir=${DIST}/${os}/${arch}
        mkdir -p ${output_dir}
        for cmd in ${cmds[@]}
        do
            cd ${DIR}/../cmd/${cmd}
            env GOOS=${os} GOARCH=${arch} \
                go build \
                    -ldflags "
                        -X github.com/nagylzs/grimdam-upgrade-artifact/internal/version.Built=${BUILT}
                        -X github.com/nagylzs/grimdam-upgrade-artifact/internal/version.Commit=${COMMIT}
                        -X github.com/nagylzs/grimdam-upgrade-artifact/internal/version.Branch=${BRANCH}" \
                    -o ${output_dir}/ ${cmd}.go
        done
    done
done
