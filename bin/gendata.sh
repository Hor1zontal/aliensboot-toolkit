#!/bin/sh
java -jar  ${ALIENSBOT_HOME}/bin/datatool.jar -d go -i ../data -o ../src/aliens/testserver/constant/tableconstant.go -t ../copy/templates/data/go_constant.template
java -jar  ${ALIENSBOT_HOME}/bin/datatool.jar -d go -i ../data -o ../src/aliens/testserver/data/tabledata.go -t ../copy/templates/data/go_model.template