#!/bin/sh

set -e

if [ ! -f "build/env.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

# Create fake Go workspace if it doesn't exist yet.
workspace="$PWD/build/_workspace"
root="$PWD"
spdir="$workspace/src/github.com/SmartPool"
if [ ! -L "$spdir/smartpool-client" ]; then
    mkdir -p "$spdir"
    cd "$spdir"
    ln -s ../../../../../. smartpool-client
    cd "$root"
fi

# Set up the environment to use the workspace.
GOPATH="$workspace"
export GOPATH

# Run the command inside the workspace.
cd "$spdir/smartpool-client"
PWD="$spdir/smartpool-client"

# Launch the arguments with the configured environment.
exec "$@"
