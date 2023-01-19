#!/usr/bin/env bash

set -Eeuo pipefail

project_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")/../../" && pwd)"

# shellcheck source=../common.sh
source "${project_dir}/src/common.sh"
work_dir="$(make_workdir "${BASH_SOURCE[0]}")"

golang_binary="${work_dir}/convert_dic"

get_dict(){
    local downloaded_file

    # ここに辞書取得のコードを書く

    echo "$downloaded_file"
}

convert_dic(){
    local basedic_path="$1"

    msg_info "Converting to mozc format..."
    "${golang_binary}" "$(get_iddef_path)" "$basedic_path" > "${work_dir}/dict.txt"
}

make_gobinary(){
    build_go_tool "${project_dir}/src/dicname/convert_dic" "$golang_binary"
}

main(){
    make_gobinary
    convert_dic "$(get_dict)"
}

main
