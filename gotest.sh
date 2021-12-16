rootDir=$(cd $(dirname $0); pwd)

go test routing-algorithms/random
go test routing-algorithms/network