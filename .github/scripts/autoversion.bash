#!/bin/bash
# shellcheck disable=SC2046
latest_tag=$(git describe --tags $(git rev-list --tags --max-count=1))
echo "$latest_tag" > version.txt
echo "version.txt updated with latest tag: $latest_tag"
