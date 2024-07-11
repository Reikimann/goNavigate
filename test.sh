
function goNavigates() {
  local dir
  dir=$(goNavigate "$@")
  if [[ -d "$dir" ]]; then
    cd "$dir"
  else
    echo "Nope"
  fi
}


