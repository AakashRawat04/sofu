#!/bin/sh
#
# Use this script to run your program LOCALLY.
#

set -e # Exit early if any commands fail
echo "Running your program locally..."

# Create a local tmp directory if it doesn't exist
SCRIPT_DIR="$(dirname "$0")"
LOCAL_TMP="${SCRIPT_DIR}/tmp"
mkdir -p "${LOCAL_TMP}"

(
  cd "${SCRIPT_DIR}" # Ensure compile steps are run within the repository directory
  go build -o "${LOCAL_TMP}/sofuserver" ./cmd/server
)


exec "${LOCAL_TMP}/sofuserver" "$@"
