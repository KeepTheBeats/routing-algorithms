#!/bin/bash

i=1

while ((i<=50))
do
	if ((i<10))
	then
		itm net0${i}.script
		sgb2alt net0${i}.script-0.gb net0${i}.alt
	else
		itm net${i}.script
		sgb2alt net${i}.script-0.gb net${i}.alt
	fi
	((i=i+1))
done
