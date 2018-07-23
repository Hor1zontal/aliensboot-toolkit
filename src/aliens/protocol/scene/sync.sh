for f in `find . -type f`;
do
        if [ -s $f ] && [ "${f##*.}"x = "proto"x ];then
              echo copy $f
              cp -rf $f /Users/hejialin/git/demo_mmorpg/tools/
        fi
done

cd /Users/hejialin/git/demo_mmorpg/tools/
python protoToJs.py