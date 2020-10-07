#########################################################################
# File Name: run.sh
# Author: bjzhang03
# mail: bjzhang03@foxmail.com
# Created Time: 2020年08月14日 星期五 14时52分43秒
#########################################################################
#!/bin/bash

set -e
go list github.com/bjzhang03/go-leetcode/... | xargs -t -i go test -cover -race {}

exit 0
