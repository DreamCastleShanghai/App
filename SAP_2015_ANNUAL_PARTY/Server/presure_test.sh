for (( i = 0; i < 2000; i++ )); do
	#curl http://localhost:8080/sap?tag=EV0&uid=
	#echo $i
	#content = "tag=djstatus&v=1"
	#$i
	#echo &content
	echo $i
	curl --data "tag=EV0&uid=1"$i http://localhost:8080/sap
	#curl --data "tag=EV0&uid=1"$i http://139.196.195.185:8080/sap
done
