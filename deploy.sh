#!/bin/zsh

echo "输入你改动的内容："
read msg
echo "格式化已有目录\n"
tree -N -o tree.txt
echo "---开始构建---\n"
git add .
git commit -m "$msg"
git push origin master
echo "\n"
echo "---构建完成---"