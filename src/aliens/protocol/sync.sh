module=$1
for f in `find ./${module}/ -type f`;
do
        if [ -s $f ] && [ "${f##*.}"x = "proto"x ];then
              echo copy $f
              cp -rf $f /Users/hejialin/git/demo_mmorpg/tools/protocol/${module}/
        fi
done

cd /Users/hejialin/git/demo_mmorpg/tools/
python protoToJs.py ${module}