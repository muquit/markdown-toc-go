#!/bin/bash

########################################################################
# push.sh
# Script to push the current Git branch to its origin remote.
# Sets upstream if not already set.
# muquit@muquit.com Jul-04-2025 
########################################################################

BRANCH=$(git rev-parse --abbrev-ref HEAD)
echo "Branch: ${BRANCH}"
exit

if [[ -z "${BRANCH}" ]]; then
    echo "Error: Could not determine current branch. Are you in a Git repository?" >&2
    exit 1
fi

echo "Pushing branch '${BRANCH}' to origin (and setting upstream if not already set)..."
git push -u origin "${BRANCH}"

if [[ $? -eq 0 ]]; then
    echo "Successfully pushed branch '${BRANCH}'."
else
    echo "Error pushing branch '${BRANCH}'." >&2
    exit 1
fi
