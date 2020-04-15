#!/bin/bash

#build directory
buildDir="./build"

#project name
projectName="txtBP"


#project output filepath
output=$buildDir/$projectName


buildCMD=$(go build -o=$output)

if $buildCMD; then
    #make the executable
    chmod +x $output
    echo '> build:successful & executable'
else
    echo '> build:failed\n'
    exit 1
fi

#bin dir
binDir="/usr/local/bin"

mvCMD=$(sudo mv $output $binDir)

if $mvCMD; then
    #make the executable
    echo '> move:successful\n'
else
    echo '> move:failed\n'
    exit 1
fi