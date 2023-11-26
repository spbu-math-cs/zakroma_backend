#! /bin/bash

purpose="Скрипт для автоматизации заполнения бд"
warning="ВНИМАНИЕ! Запуск load_all.sql УДАЛИТ ВСЕ ТАБЛИЦЫ в целевой базе" # как грится, use with caution

rm -f load_all.sql; # затираем предыдущий скрипт заполнения
# пишем слова, добавляем sql-скрипт для удаления предыдущих таблиц
cat > load_all.sql <<< """-- $purpose
-- !!!
-- !!! $warning
-- !!!

DO \$$ DECLARE
    r RECORD;
BEGIN
    FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = current_schema()) LOOP
            EXECUTE 'DROP TABLE IF EXISTS ' || quote_ident(r.tablename) || ' CASCADE';
        END LOOP;
END \$$;
"""

cd scripts; # файлы с заполнением должны лежать в ./scripts
for FILENAME in `ls . | sort -g`; # перебираем все файлы с заполнением
do
  new_name=`echo $FILENAME | tr ' ' _ | tr - _`; # заменяем в имени файла все пробелы и - на _
  mv "$FILENAME" $new_name -n; # переименовываем
  if [[ $new_name = *.sql ]]; then # если файл имеет расширение .sql
    echo "\ir scripts/$new_name" >> ../load_all.sql; # добавляем его в скрипт заполнения
  fi;
done;
