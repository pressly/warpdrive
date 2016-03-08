#!/bin/bash
set -e

mkdir -p $WORKDIR
cd $WORKDIR

# Clone the repo for the first time.
if [[ ! -d $WORKDIR || ! -d $WORKDIR/.git ]]; then
	git clone -b $BRANCH $REPO ./
fi

# Clean up the repository.
git reset --hard && git clean -f -d

# Switch to master and update.
git checkout master && git pull

# Remove all branches except for master.
# (This prevents error when someone rebases his branch.)
for branch in $(git branch | grep -v master || :); do
	git branch -D $branch
done

# If we're deploying a custom branch, we need to pull it first.
if [ "$BRANCH" != "master" ]; then
	git checkout -b $BRANCH origin/$BRANCH
fi

# Cleanup.
sudo rm -rf bin

# Builder image.
sudo docker run --rm \
	-v $(pwd):/go/src/github.com/$IMAGE \
	-w /go/src/github.com/$IMAGE \
	--name ${NAME}_builder \
	$BUILDER_IMAGE make build

# Build the resulting image. Tag it with version, then retag the latest.
VERSION=$(./scripts/version.sh --long)
sudo docker build --rm --no-cache -t $IMAGE:$VERSION .
sudo docker push $IMAGE:$VERSION
sudo docker tag -f $IMAGE:$VERSION $IMAGE:latest
sudo docker push $IMAGE:latest
