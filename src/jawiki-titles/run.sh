#!/usr/bin/env bash

set -Eeuo pipefail

project_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")/../../" && pwd)"

# shellcheck source=../common.sh
source "${project_dir}/src/common.sh"
work_dir="$(make_workdir "${BASH_SOURCE[0]}")"

get_dict(){
    local downloaded_file
    msg_info "Getting dictionary file from Wikipedia..."
    downloaded_file=$(download_file "$work_dir" "https://dumps.wikimedia.org/jawiki/latest/jawiki-latest-all-titles-in-ns0.gz")
    

    # extract file
    msg_info "Extracting downloaded file..."
    7z x -aos "$downloaded_file" -o"$work_dir" >&2

    echo "$work_dir/jawiki-latest-all-titles-in-ns0"
}

main(){
    check_tools
    get_dict
}

main
