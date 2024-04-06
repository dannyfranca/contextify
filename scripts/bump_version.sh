#!/bin/bash

# Function to get current version from version.go
get_current_version() {
  current_version=$(grep -oP 'version = "\K[^"]+' cmd/version.go)
  if [[ -z $current_version ]]; then
    echo "Could not find current version in version.go"
    exit 1
  fi
  echo "$current_version"
}

# Function to bump version based on version type
bump_version() {
  local current_version=$1
  local version_type=$2

  IFS='.' read -ra version_parts <<< "$current_version"
  major=${version_parts[0]}
  minor=${version_parts[1]}
  patch=${version_parts[2]}

  case $version_type in
    patch)
      patch=$((patch + 1))
      ;;
    minor)
      minor=$((minor + 1))
      patch=0
      ;;
    major)
      major=$((major + 1))
      minor=0
      patch=0
      ;;
    *)
      echo "Invalid version type. Please specify either patch, minor, or major."
      exit 1
      ;;
  esac

  bumped_version="$major.$minor.$patch"
  echo "$bumped_version"
}

# Prompt for version type
read -p "Enter the version type to bump (patch, minor, major): " version_type

# Get current version
current_version=$(get_current_version)
echo "Current version: $current_version"

# Bump version
bumped_version=$(bump_version "$current_version" "$version_type")
echo "Bumped version: $bumped_version"

# Get current commit hash
commit=$(git rev-parse HEAD)

# Get current date
date=$(date +%F)

# Update version in version.go
sed -i -E "s/(version\s*=\s*\")[^\"]+/\1$bumped_version/" cmd/version.go
sed -i -E "s/(commit\s*=\s*\")[^\"]+/\1$commit/" cmd/version.go
sed -i -E "s/(date\s*=\s*\")[^\"]+/\1$date/" cmd/version.go

# Commit changes
git add cmd/version.go
git commit -m "Bump version to $bumped_version"

# Create Git tag
git tag -a "v$bumped_version" -m "Version $bumped_version"

echo "Version bumped to $bumped_version, commit $commit, date $date"
