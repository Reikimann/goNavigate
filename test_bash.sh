#!/bin/bash

function goNavigatess() {
  # Run goNavigate directly without capturing its output
  # goNavigate "$@"
  local dir
  dir=$(goNavigate "$@")
  echo "$dir"

  # Capture the exit status
  local status=$?

  # If goNavigate was successful, it should have printed a directory path as its last line
  if [[ $status -eq 0 ]]; then
    echo "successful"
    local dir
    dir=$(goNavigate "$@" | tail -n 1)
    if [[ -d "$dir" ]]; then
      cd "$dir"
    else
      echo "Invalid directory: $dir"
    fi
  else
    echo "Navigation cancelled or failed"
  fi
}

# If this script is sourced, make the function available to the parent shell
if [[ "${BASH_SOURCE[0]}" != "${0}" ]]; then
  export -f goNavigatess
fi
