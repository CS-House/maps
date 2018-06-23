#!/usr/bin/env bash

genString() {
	echo -ne '$865691034264978,'
	date +%y/%m/%d,%I:%M:%S | tr -d '\n'
	weight=$RANDOM
	echo -ne +22,12.824345,80.214066,000088716E97,2,${weight:0:1}.${weight:1}
	# echo -n $RANDOM
	echo "#"
}

genString

# (
# 	while true; do
# 		genString | nc 139.59.70.218 10331
# 		sleep 5
# 	done
# )&
# DUMMYPID="$!"
# echo $DUMMYPID
# export DUMMYPID

# #*ZJ,2030295125,V1,073614,A,3127.7080,N,7701.8360,E,0.00,0.00,040618,00000000#