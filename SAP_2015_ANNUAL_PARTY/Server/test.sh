function random()
{
    min=$1;
    max=$2-$1;
    num=$(date +%s+%N);
    ((retnum=num%max+min));
    echo $retnum;
}

for ((i = 0; i < 2000; i++)); do
#for (( j = 0; j < 3; j++ )); do
userID=$(random 1 2000);
vote1=$(random 0 7);
vote2=$(random 0 2);
        curl --data "tag=EV0&uid=$i" http://139.196.195.185:8080/sap
#        curl --data "tag=DV0&uid=$userID&vid=$vote1" http://139.196.195.185:8080/sap
#        curl --data "tag=VV0&uid=$userID&vid=$vote2" http://139.196.195.185:8080/sap
#done
done                      