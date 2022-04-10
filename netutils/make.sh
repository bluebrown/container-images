#!/usr/bin/env sh

set -eu

: "${CONTAINER_CLI:=docker}"
: "${REGISTRY:=docker.io}"
: "${NAMESPACE:=bluebrown}"

image="$REGISTRY/$NAMESPACE/netutils"

build() {
    context="$(dirname "$(realpath "$0")")"
    "$CONTAINER_CLI" build -t "$image:latest" "$context"
}

push() {
    # set the version
    version="0.1.0"

    # tag the image
    "$CONTAINER_CLI" tag "$image:latest" "$image:$version"

    # push the image
    "$CONTAINER_CLI" push "$image:latest"
    "$CONTAINER_CLI" push "$image:$version"
}

case "$1" in
build)
    build
    ;;
push)
    push
    ;;
*)
    echo "Usage: $0 build|push"
    exit 1
    ;;
esac

exit 0
