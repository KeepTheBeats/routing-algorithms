#go run "/d/code/go/src/routing-algorithms/experiments/r2t-dsdn-config/main.go"


#!/bin/bash
i=1
while ((i<=20))
do
    go run "/d/code/go/src/routing-algorithms/experiments/r2t-dsdn-config/main.go"
    if ((i<10))
    then
        cp -rf "/d/code/go/src/routing-algorithms/experiments/r2t-dsdn-config/jsonnetworks" "/d/code/go/src/routing-algorithms/experiments/r2t-dsdn-config/results/result0${i}"
    else
        cp -rf "/d/code/go/src/routing-algorithms/experiments/r2t-dsdn-config/jsonnetworks" "/d/code/go/src/routing-algorithms/experiments/r2t-dsdn-config/results/result${i}"
    fi
    ((i=i+1))
done

