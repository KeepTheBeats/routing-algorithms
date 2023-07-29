rootDir=$(cd $(dirname $0); pwd)

go test ${rootDir}/random -count=1
go test ${rootDir}/network -count=1
go test ${rootDir}/mymath -count=1