call env.bat

java -jar  %ALIENSBOT_HOME%/bin/datatool.jar -d go -i %PROJECT_PATH%/copy/data -o %PROJECT_PATH%/src/aliens/testserver/constant/tableconstant.go -t %PROJECT_PATH%/copy/templates/data/go_constant.template
java -jar  %ALIENSBOT_HOME%/bin/datatool.jar -d go -i %PROJECT_PATH%/copy/data -o %PROJECT_PATH%/src/aliens/testserver/data/tabledata.go -t %PROJECT_PATH%/copy/templates/data/go_model.template