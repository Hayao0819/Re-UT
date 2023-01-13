#!/usr/bin/env bash

set -Eeuo pipefail

current_dir=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)

# shellcheck source=./common.sh
source "${current_dir}/common.sh"

get_download_dict(){
    local filelist_json=() filename url

    readarray -t filelist_json < <(
        curl -H "Accept: application/vnd.github+json" \
            -H "X-GitHub-Api-Version: 2022-11-28" \
            "https://api.github.com/repos/neologd/mecab-ipadic-neologd/git/trees/master?recursive=true"
    )

    filename="$(
        printarg "${filelist_json[@]}" | jq ".tree[].path" | grep "seed/mecab-user-dict-seed"
    )"
    

    url="https://github.com/neologd/mecab-ipadic-neologd/raw/master/${filename}"
    
}
