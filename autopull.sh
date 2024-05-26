#!/bin/bash

BRANCH=$1
REMOTE_BRANCH=origin/$BRANCH
DELAY=10

while true; do
	git fetch -q origin $BRANCH
	own_commit=$(git rev-parse $BRANCH)
	remote_commit=$(git rev-parse $REMOTE_BRANCH)
	commits_count=$(git rev-list --count $BRANCH..$REMOTE_BRANCH)

	if [ "$own_commit" != "$remote_commit" ]; then
		git reset -q --hard "$REMOTE_BRANCH"
		if [ "$commits_count" != "0" ]; then
			echo "Received update with $commits_count new commits"
		else
			echo "Received force update"
		fi
		echo "Rebuilding project..."
		docker compose up --build -d
	fi

	sleep $DELAY
done
