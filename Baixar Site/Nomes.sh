#menu de ajuda
if [ $# -eq 0 ];
then
  echo "Parametro 1: url da pagina do hackerrank"
  echo "Parametro 2: extensao do arquivo"
  echo "Parametro 3: numero da ultima atividade baixada"
else
  javac *.java ; java Download $1 | grep '</a>' | java Tratar $2 $3
fi
exit
