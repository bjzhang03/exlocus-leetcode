#########################################################################
# File Name: build.sh
# Author: bjzhang03
# mail: bjzhang03@foxmail.com
# Created Time: 2020年08月14日 星期五 15时33分21秒
#########################################################################
#!/bin/bash

set -e
go list github.com/bjzhang03/go-leetcode/... | xargs -t -i go build -v {}
