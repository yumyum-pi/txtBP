#!/bin/bash

clear
#build directory
buildDir="./build"

#project name
projectName="txtBP"


#project output filepath
buildPath=$buildDir/$projectName

#bin dir
binDir="/usr/local/bin"

mvCMD=$(sudo mv $buildPath $binDir)

if $mvCMD; then
    #make the executable
    echo '> move:successful\n'
else
    echo '> move:failed\n'
    exit 1
fi

