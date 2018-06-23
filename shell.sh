#!/usr/bin/env bash

dids[0]="EUgnN1N1Eu"
dids[1]="MRl06xqlYR"
dids[2]="R6E2YeGHi5"
dids[3]="G9ceDX6fTT"
dids[4]="Zwkpbm93vj"
dids[5]="gFGKPNdTiO"
dids[6]="Py5qQAcf0i"
dids[7]="GlwuEiwrHO"
dids[8]="9WMPR3n0J9"
dids[9]="QkypngR4p4"

ll=($(python -c "import random;for i in range(1,10):print('%0.6f'%random.uniform(1, 100))"))
echo "${ll[0]}"

lat[0]=$(python -c "import random;print('%0.6f'%random.uniform(1, 100))")
lat[1]=$(python -c "import random;print('%0.6f'%random.uniform(1, 100))")
lat[2]=$(python -c "import random;print('%0.6f'%random.uniform(1, 100))")
lat[3]=$(python -c "import random;print('%0.6f'%random.uniform(1, 100))")
lat[4]=$(python -c "import random;print('%0.6f'%random.uniform(1, 100))")
lat[5]=$(python -c "import random;print('%0.6f'%random.uniform(1, 100))")
lat[6]=$(python -c "import random;print('%0.6f'%random.uniform(1, 100))")
lat[7]=$(python -c "import random;print('%0.6f'%random.uniform(1, 100))")
lat[8]=$(python -c "import random;print('%0.6f'%random.uniform(1, 100))")
lat[9]=$(python -c "import random;print('%0.6f'%random.uniform(1, 100))")


genString() {
    int=$(shuf -i 0-9 -n 1)
	echo -ne "GTPL $"$(shuf -i 0-9 -n 1),${dids[$int]},"A,"
	date +%d%m%y,%I%M%S | tr -d '\n'
	weight=$RANDOM
	echo -ne ,"\n"
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
# GTPL $1,867322035135813,A,290518,062804,18.709738,N,80.068397,E,0,406,309,11,0,14,1,0,26.4470#