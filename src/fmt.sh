golist=`find . -name "*.go"`
for gofile in $golist
do
	if [[ ! ${gofile} =~ ".pb.go" ]]
	then
		cp ${gofile} ${gofile}.bak
		gofmt ${gofile}.bak > ${gofile}
		rm -f ${gofile}.bak
		echo "gofmt ${gofile}"
	fi
done
