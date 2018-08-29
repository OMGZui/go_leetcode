#!/bin/zsh

echo -n "输入你改动的内容："
read msg

tree -N -o tree.txt
echo '###防止老年痴呆，go语言描述LeetCode' > README.md
echo '```' >> README.md
cat tree.txt >> README.md
echo '```' >> README.md

echo "---开始构建---\n"
git add .
git commit -m "$msg"
git push origin master
echo "\n---构建完成---"