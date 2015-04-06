
rm stop.txt
touch stop.txt
while [ $(wc -l stop.txt | cut -f1 -d " ") -eq 0 ];
do echo test;
   ls *.go | shuf | head -2 | tee test.txt | xargs ~/install/bin/gccgo > out.txt 2>err.txt;
   echo $?;
   grep report out.txt err.txt > stop.txt;
done
