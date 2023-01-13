#!/usr/bin/env bash

check_tools(){
    local tools=("jq" "curl" "git")
    for tool in "${tools[@]}"; do
        if ! command -v "${tool}" >/dev/null 2>&1; then
            echo "Error: ${tool} is not installed." >&2
            exit 1
        fi
    done
}

printarg(){
    printf "%s\n" "$@"
}

printary(){
    eval "printarg \"\${${1}[@]}\""
}

