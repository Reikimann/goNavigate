
# This is not a part of the program. It is a test and will be deleted!
function goNavigates() {
  local dir
  dir=$(goNavigate "$@")
  if [[ -d "$dir" ]]; then
    cd "$dir"
  else
    echo "Nope"
  fi
}


