rootDir=$(cd $(dirname $0); pwd)

go test ${rootDir}/random
go test ${rootDir}/network