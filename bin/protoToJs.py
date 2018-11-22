# -*- coding: UTF-8 -*- 需要protobuf 5这个版本  6脚本命令改了  巨坑
import os
import shutil
import platform
import sys

#生成的js文件名
scriptName='protocol'
#协议包名
packageName='protocol'
#项目下的路径


#destDir = '/Users/hejialin/git/demo_mmorpg/assets/Script/module/'+scriptName+'.js'

srcDir = sys.argv[1]
destFile = sys.argv[2]

if __name__ == '__main__':
    system = platform.system()
    cmd = ''
    param = srcDir + '/*.proto -t js -e '+packageName+' -i populateAccessors -m true -o ./'+scriptName+'.js'
    if (system == "Darwin") :
        cmd = ' ./protobufjsForMac/bin/pbjs ' + param
    else:
        cmd = '.\protobufjsForWin\pbjs  ' + param
    print(cmd)
    os.system(cmd)

    file_object = open('./'+scriptName+'.js')
    file_context = file_object.read()
    file_context = file_context.replace('proto2','proto3')
    file_object.close()

    file_object = open('./'+scriptName+'.js','w')
    file_object.write(file_context)
    file_object.close()

    # 打开一个文件
    fo = open('./'+scriptName+'.js', "a+")
    fo.write( "\nmodule.exports="+packageName+";");
    # 关闭打开的文件
    fo.close()
    shutil.move('./'+scriptName+'.js',destFile)
