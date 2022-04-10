#!/usr/bin/env sh

set -eu

: "${CONTAINER_CLI:=docker}"
: "${REGISTRY:=docker.io}"
: "${NAMESPACE:=bluebrown}"

image="$REGISTRY/$NAMESPACE/azcopy"
context="$(dirname "$(realpath "$0")")"

build() {
    "$CONTAINER_CLI" build -t "$image:latest" "$context"
}

push() {
    # get the version
    version="$("$CONTAINER_CLI" run --rm "$image:latest" -v | cut -d' ' -f3)"
    major="$(echo "$version" | cut -d'.' -f1)"
    minor="$(echo "$version" | cut -d'.' -f2)"
    major_minor="$major.$minor"

    # tag the image
    "$CONTAINER_CLI" tag "$image:latest" "$image:$major"
    "$CONTAINER_CLI" tag "$image:latest" "$image:$major_minor"
    "$CONTAINER_CLI" tag "$image:latest" "$image:$version"

    # push the image
    "$CONTAINER_CLI" push "$image:latest"
    "$CONTAINER_CLI" push "$image:$major"
    "$CONTAINER_CLI" push "$image:$major_minor"
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
