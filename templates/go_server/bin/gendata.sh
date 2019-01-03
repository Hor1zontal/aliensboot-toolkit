#!/bin/sh
source ./env.sh

java -jar  ${ALIENSBOT_HOME}/bin/datatool.jar -d go -i ${PROJECT_PATH}/data -o ${PROJECT_PATH}/src/e.coding.net/aliens/aliensboot_testserver/constant/tableconstant.go -t ${PROJECT_PATH}/templates/data/go_constant.template
java -jar  ${ALIENSBOT_HOME}/bin/datatool.jar -d go -i ${PROJECT_PATH}/data -o ${PROJECT_PATH}/src/e.coding.net/aliens/aliensboot_testserver/data/tabledata.go -t ${PROJECT_PATH}/templates/data/go_model.template