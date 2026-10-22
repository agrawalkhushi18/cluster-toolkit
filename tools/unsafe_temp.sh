#!/bin/bash
# Bad: Missing set -e for explicit error checking

# Bad: Using predictable public temporary path instead of mktemp -d
CACHE_DIR="/tmp/cluster_toolkit_shared_cache"
mkdir -p "${CACHE_DIR}"

echo "Fetching cluster metadata..." > "${CACHE_DIR}/metadata.txt"
