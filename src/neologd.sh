#!/usr/bin/env bash

set -Eeuo pipefail

current_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# shellcheck source=./common.sh
source "${current_dir}/common.sh"

work_dir="$(make_workdir "$0")"
original_csv=""

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

original_csv="$(get_dict)"


