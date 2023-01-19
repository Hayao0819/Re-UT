#!/usr/bin/env bash

set -Eeuo pipefail

project_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")/../../" && pwd)"

# shellcheck source=../common.sh
source "${project_dir}/src/common.sh"
work_dir="$(make_workdir "${BASH_SOURCE[0]}")"

golang_binary="${work_dir}/convert_dic"

get_dict(){
    local downloaded_file
    local url="http://ftp.edrdg.org/pub/Nihongo/edict2.gz"
    downloaded_file=$(download_file "$work_dir" "$url")
    unset url

    # extract file
    msg_info "Extracting downloaded file..."
    7z x -aos "$downloaded_file" -o"$work_dir" >&2

    local original_file="$work_dir/edict2"
    if [[ ! -f "$original_file" ]]; then
        msg_error "File not found: $original_file"
        return 1
    fi

    # convert to utf-8
    local converted_file="$work_dir/edict2.utf8"
    msg_info "Converting to utf-8..."
    nkf -w "$original_file" | grep -v "^　" | grep " /(n" | cut -d "/" -f 1 > "$converted_file"

    echo "$converted_file"
}

convert_dic(){
    local basedic_path="$1"

    msg_info "Converting to mozc format..."
    "${golang_binary}" "$(get_iddef_path)" "$basedic_path" > "${work_dir}/dict.txt"
}

make_gobinary(){
    build_go_tool "${project_dir}/src/edict2/convert_dic" "$golang_binary"
}

main(){
    make_gobinary
    convert_dic "$(get_dict)"
}

main
