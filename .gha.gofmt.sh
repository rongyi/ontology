#!/bin/bash
# code from https://github.com/Seklfreak/Robyul2
which goimports || cd /tmp/; go get -u -v golang.org/x/tools/cmd/goimports; cd -

unset dirs files
dirs=$(go list -f {{.Dir}} ./... | grep -v /vendor/)

for d in $dirs
do
    for f in $d/*.go
    do
    files="${files} $f"
    done
done

diff <(goimports -d $files) <(echo -n)
