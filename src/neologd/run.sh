#!/usr/bin/env bash

set -Eeuo pipefail

project_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")/../../" && pwd)"

# shellcheck source=../common.sh
source "${project_dir}/src/common.sh"

work_dir="$(make_workdir "${BASH_SOURCE[0]}")"
msg_info "Build directory is $work_dir"

golang_binary="${work_dir}/convert_dic"

get_dict(){
    

    # ファイル名を取得
    {
        local filelist_json=() filename url
        msg_info "Getting file name from GitHub API..."
        readarray -t filelist_json < <(
            curl -H "Accept: application/vnd.github+json" \
                -H "X-GitHub-Api-Version: 2022-11-28" \
                "https://api.github.com/repos/neologd/mecab-ipadic-neologd/git/trees/master?recursive=true"
        )

        filename="$(
            printarg "${filelist_json[@]}" | jq -r ".tree[].path" | grep "seed/mecab-user-dict-seed"
        )"
        url="https://github.com/neologd/mecab-ipadic-neologd/raw/master/${filename}"
        unset filelist_json filename
    }

    # ファイルをダウンロード
    local downloaded_file
    downloaded_file="$(download_file "${work_dir}" "${url}")"

    # ファイルを解凍
    local dic_csv
    msg_info "Extracting downloaded file..."
    7z x -aos "${downloaded_file}" -o"${work_dir}" >&2
    dic_csv="$(find "${work_dir}" -type f -name "*.csv")"
    msg_info "${dic_csv} is extracted."

    unset downloaded_file

    # test file
    if [[ -z "${dic_csv-""}" ]] || [[ ! -f "${dic_csv}" ]]; then
        msg_error "File not found: ${dic_csv}"
        exit 1
    fi
    echo "${dic_csv}"
    unset dic_csv

}

convert_dic(){
    local csv_path="$1"
    msg_info "Converting dictionary..."
    rm -rf "${work_dir}/dic.txt"
    #"$golang_binary" "$(get_iddef_path)" "${csv_path}" | tee "${work_dir}/dic.txt"
    {
        "$golang_binary" "$(get_iddef_path)" "${csv_path}" > "${work_dir}/dic.txt"
    } &
    
    
    while true; do
        num_children="$(jobs | wc -l)"
        if (( num_children == 0 )); then
            break
        else
            echo -ne "\033[2K"
            wc -l "${work_dir}/dic.txt" 2>/dev/null || continue
            printf "\033[%dA" "1"
        fi
        sleep 1
    done 
}


get_id(){
    local target_rawcsv="$1" target_csv
    target_csv="$(cut -d "," -f 5,6,7,8  <<< "$target_rawcsv")"

    get_iddef | grep -E "[0-9]* ${target_csv}.*" | cut -d " " -f 1 | head -n 1
}

make_gobinary(){
    rm -rf "${golang_binary}"
    go build -o "$golang_binary" "${project_dir}/src/neologd/convert_dic/"*".go"
}

main(){
    make_gobinary
    convert_dic "$(get_dict)"
}

main
