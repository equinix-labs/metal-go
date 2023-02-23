#!/bin/sh

set -eu -o pipefail

function find_path_prefix () {
  local parent_path=$(realpath `dirname $1`)
  local child_path=$(realpath `dirname $2`)
  echo $parent_path
  echo $child_path
  path_prefix=${child_path#$parent_path}
}

function download_file () {
  local spec_base_url=$1
  local output_dir=$2
  local path=$3
  local full_path=$output_dir$path
  if [ ! -f "$full_path"  ]
  then
    echo "Fetching $spec_base_url$path"
    mkdir -p $(dirname $full_path)
    wget -O $full_path $spec_base_url$path

    ref_targets="$(yq e '.. | select(has("$ref")) | .$ref' $full_path)"

    for path in $ref_targets
    do
      find_path_prefix "$output_dir/." $full_path
      download_file $spec_base_url $output_dir $path_prefix/$path
    done
  fi
}

SPEC_BASE_URL=$1
OUTPUT_BASE_DIR=$2
ROOT_FILE=$3

rm -rf $OUTPUT_BASE_DIR

download_file $SPEC_BASE_URL $OUTPUT_BASE_DIR /$ROOT_FILE
