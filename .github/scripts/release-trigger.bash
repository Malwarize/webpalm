#get the latest tag name from remote
git fetch --tags
# shellcheck disable=SC2046
# shellcheck disable=SC2006
latest_tag=$(git describe --tags `git rev-list --tags --max-count=1`)
echo "latest tag is $latest_tag"

# then increment it
new_tag=$(echo "$latest_tag" | awk -F. -v OFS=. '{$NF++;print}')
echo "new tag is $new_tag"

# store it in version.txt
echo $new_tag > version.txt
# commit and push the change
git add version.txt
git commit -m "bump version to $new_tag"
git push origin master

# create a tag
git tag "$new_tag"
git push origin "$new_tag"