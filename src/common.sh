#!/usr/bin/env bash

current_dir="${current_dir-"$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"}"
project_dir="$current_dir/.."
export CHILD_MAX=5

check_tools(){
    local tools=("jq" "curl" "git" "nkf")
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


# shellcheck disable=SC2034
{
    readonly ascii_black="\033[30m"
    readonly ascii_red="\033[31m"
    readonly ascii_green="\033[32m"
    readonly ascii_yellow="\033[33m"
    readonly ascii_blue="\033[34m"
    readonly ascii_magenta="\033[35m"
    readonly ascii_cyan="\033[36m"
    readonly ascii_white="\033[37m"
    readonly ascii_bgblack="\033[40m"
    readonly ascii_bgred="\033[41m"
    readonly ascii_bggreen="\033[42m"
    readonly ascii_bgyellow="\033[43m"
    readonly ascii_bgblue="\033[44m"
    readonly ascii_bgmagenta="\033[45m"
    readonly ascii_bgcyan="\033[46m"
    readonly ascii_bgwhite="\033[47m"

    readonly ascii_bold="\033[1m"
    readonly ascii_dim="\033[2m"
    readonly ascii_italic="\033[3m"
    readonly ascii_underline="\033[4m"
    readonly ascii_blink="\033[5m"
    readonly ascii_reverse="\033[7m"
    readonly ascii_concealed="\033[8m"
    readonly ascii_strike="\033[9m"

    readonly ascii_reset="\033[0m"
}

msg_info(){
    echo -e "${ascii_green}[INFO]${ascii_reset} $*" >&2
}

msg_warn(){
    echo -e "${ascii_yellow}[WARN]${ascii_reset} $*" >&2
}

msg_error(){
    echo -e "${ascii_red}[ERROR]${ascii_reset} ${*}" >&2
}

# make_workdir "$0"
make_workdir(){
    local baseworkdir="${project_dir}/build"
    local script_name="${1-""}"
    [[ -n "$script_name" ]] || {
        msg_error "make_workdir: Missing argument"
        return 1
    }
    script_name="$(basename "$script_name")"
    local workdir="${baseworkdir}/${script_name}"
    mkdir -p "${workdir}"
    echo "${workdir}"
    return 0
}

# download_file <work> <url>
download_file(){
    local output work_dir="${1-""}"  url="${2-""}"
    { [[ -n "${url}" ]] && [[ -n "${work_dir}" ]]; } || {
        msg_error "download_file: No url"
        return 1
    }
    output="${work_dir}/$(basename "$url")"
    if [[ -e "${output}" ]]; then
        echo "${output}"
        return 0
    fi
    msg_info "Downloading ${url} to ${output} ..."
    curl -sfL -o "${output}" "$url" || {
        msg_error "download_file: Failed to download ${url}"
        return 1
    }
    echo "${output}"
    return 0
}

hira2kata(){
    nkf -w --katakana <<< "$*"
}

kata2hira(){
    nkf -w --hiragana <<< "$*"
}

get_iddef(){
    cat "$(download_file "$(make_workdir "$0")" "https://raw.githubusercontent.com/google/mozc/master/src/data/dictionary_oss/id.def")"
}

write_string(){
    local lockfile="$1" targetfile="$1" text="$2"
    echo "$text" >> "$targetfile"
}
