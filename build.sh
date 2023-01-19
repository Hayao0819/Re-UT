#!/usr/bin/env bash

set -Eeuo pipefail

project_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")/" && pwd)"

# shellcheck source=./src/common.sh
source "${project_dir}/src/common.sh"
work_dir="$project_dir/build/"
mkdir -p "$work_dir"

## 現在ビルド可能な辞書の一覧
open_dictlist=("skkdic" "neologd")
all_dictlist=("${open_dictlist[@]}" "jawiki")

## ビルド対象の辞書の一覧
target_dictlist=()

usage_text(){
    echo "Usage: $0 [options]"
    echo "Useful options:"
    echo "    --all     すべての辞書をビルドします"
    echo "    --open    再配布可能な辞書のみをビルドします"
    echo
    echo "Manual build options:"
    echo "    --skk     skk辞書をビルド"
    #echo "    --jawiki  jawiki辞書をビルド"
    echo "    --neologd neologd辞書をビルド"
    echo
    echo "Other options:"
    #echo "    --mozc    Mozcのビルドも行います"
    #echo "    --clean   一時ファイルを削除"
    echo "    --help    ヘルプを表示"
}


while (( "$#" != 0 )); do
    case "$1" in
        "--skk")
            target_dictlist+=("skkdic")
            ;;
        "--neologd")
            target_dictlist+=("neologd")
            ;;
        "--all")
            target_dictlist=("${all_dictlist[@]}")
            ;;
        "--open")
            target_dictlist=("${open_dictlist[@]}")
            ;;
        "--help")
            usage_text
            exit 0
            ;;
        *)
            msg_error "Unknown option: $1"
            usage_text
            exit 1
            ;;
    esac
done


if (( "${#target_dictlist[@]}" == 0 )); then
    ## あとで辞書選択メニューを実装する
    msg_error "No dictionary selected."
    usage_text
    exit 1
fi

## 辞書のビルド
dictxt_list=()
for dict in "${target_dictlist[@]}"; do
    msg_info "Building $dict..."
    case "$dict" in
        "skkdic")
            "${project_dir}/src/skkdic/run.sh"
            dictxt_list+=("${work_dir}/skkdic/dict.txt")
            ;;
        "neologd")
            "${project_dir}/src/neologd/run.sh"
            dictxt_list+=("${work_dir}/neologd/dict.txt")
            ;;
    esac

    cat "${dictxt_list[@]}" > "${work_dir}/dict.txt"
done
