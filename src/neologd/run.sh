#!/usr/bin/env bash

set -Eeuo pipefail

project_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")/../../" && pwd)"

# shellcheck source=../common.sh
source "${project_dir}/src/common.sh"

work_dir="$(make_workdir "${BASH_SOURCE[0]}")"
msg_info "Build directory is $work_dir"

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
    local csv_path="$1" current_process=0
    while read -r raw_csv; do
        if [[ "${current_process}" -ge 8192 ]]; then
            wait
            current_process=0
        fi
        current_process="$((current_process + 1))"
        convert_oneline_dic "${raw_csv}" >> "${work_dir}/dic.txt"
    done < "${csv_path}"
}

# めちゃくちゃ遅いのでGoで書き直す
convert_oneline_dic_old(){
    local yomi tango raw_csv="$1" id cost
    local parsed_csv=()
    readarray -t parsed_csv < <(tr "," "\n" <<< "$raw_csv")
    yomi="$(kata2hira "${parsed_csv[11]}")"
    tango="$(sed "s|^#||g" <<< "${parsed_csv[10]}")"
    cost="${parsed_csv[2]}"
    id="$(get_id "${raw_csv}")"
    write_string "${work_dir}/dic.txt" "$(echo -e "${yomi}\t${id}\t${id}\t${cost}\t${tango}")"
}

convert_oneline_dic(){
    local execFile="${project_dir}/src/neologd/convert_oneline_dic/convert_oneline_dic"
    [[ -e "$execFile" ]] || go build -o "$execFile" "${project_dir}/src/neologd/convert_oneline_dic/"*".go"
    local run=("$execFile" "$(get_iddef_path)" "$1")
    #msg_info "Running ${run[*]}"
    "${run[@]}"
}

get_id(){
    local target_rawcsv="$1" target_csv
    target_csv="$(cut -d "," -f 5,6,7,8  <<< "$target_rawcsv")"

    get_iddef | grep -E "[0-9]* ${target_csv}.*" | cut -d " " -f 1 | head -n 1
}

remove_old_binary(){
    rm -rf "${project_dir}/src/neologd/convert_oneline_dic/convert_oneline_dic"
}

main(){
    remove_old_binary
    convert_dic "$(get_dict)"
}

main
