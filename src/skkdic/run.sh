#!/usr/bin/env bash

set -Eeuo pipefail

project_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")/../../" && pwd)"

# shellcheck source=../common.sh
source "${project_dir}/src/common.sh"
work_dir="$(make_workdir "${BASH_SOURCE[0]}")"

golang_binary="${work_dir}/convert_dic"

get_dict(){
    local downloaded_file 
    local url="https://skk-dev.github.io/dict/SKK-JISYO.L.gz"
    msg_info "Downloading $url"
    downloaded_file=$(download_file "$work_dir" "$url")
    unset url

    # extract file
    msg_info "Extracting downloaded file..."
    7z x -aos "$downloaded_file" -o"$work_dir" >&2

    local original_file="$work_dir/SKK-JISYO.L"
    if [[ ! -f "$original_file" ]]; then
        msg_error "File not found: $original_file"
        return 1
    fi

    # convert to utf-8
    local converted_file="$work_dir/SKK-JISYO.L.utf8"
    msg_info "Converting to utf-8..."
    nkf -w "$original_file" | grep -v "^;;" > "$converted_file"

    echo "$converted_file"
}

convert_dic(){
    local basedic_path="$1"

    # 英字を含む行を削除
    local filtered_file="${basedic_path}.filtered"
    grep -vE "[a-zA-Z]" "$basedic_path" > "$filtered_file"

    msg_info "Converting to mozc format..."
    "${golang_binary}" "$(get_iddef_path)" "$filtered_file" > "${work_dir}/dict.txt"
}

make_gobinary(){
    build_go_tool "${project_dir}/src/skkdic/convert_dic" "$golang_binary"
}

main(){
    make_gobinary
    convert_dic "$(get_dict)"
}

main
