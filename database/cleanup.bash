#! /bin/bash

purpose="Скрипт для автоматизации заполнения бд"

rm -f load_all.sql; # затираем предыдущий скрипт заполнения
# пишем слова, добавляем sql-скрипт для удаления предыдущих таблиц
cat > load_all.sql <<< "-- $purpose"

cd scripts; # файлы с заполнением должны лежать в ./scripts
for FILENAME in $(ls . | sort -g); # перебираем все файлы с заполнением
do
  new_name="$(echo $FILENAME | tr ' ' _ | tr - _)";
  mv "$FILENAME" "$new_name" -n; # заменяем в имени файла все пробелы и '-' на '_'
  if [[ $new_name = *.sql ]]; then # если файл имеет расширение .sql
    echo "\ir scripts/$new_name" >> ../load_all.sql; # добавляем его в скрипт заполнения
  fi;
done;
