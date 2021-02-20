#!/usr/bin/env bash

set -euo pipefail

function main() {
    # Get the last argument, i.e. PROJECT
    for project; do true; done
    if [[ -z "${project}" ]]; then
        usage
        exit 1;
    fi

    # =======================================
    # GENERATING PROCESS
    # =======================================
    pushd "$project"
        # Generate proto models.
        echo "Generating proto models..."
        go generate -x ./proto

        # Generate contracts.
        # echo "Generating contracts..."
        # go generate -x ./contracts

        # Generate rest-api-svc Swagger models
        # echo "Generating Swagger models..."
        # go generate -x ./services/rest-api-svc/swaggergen
    popd
}

function usage() {
    echo "Usage:"
    echo "    ./scripts/generate <PROJECT>"
}

main "$@"
