#! /bin/bash

purpose=$'Скрипт для автоматизации запуска докера.'
docker_output_path=$'docker_output.out.txt'

exitfn () {
  trap SIGINT
  sudo systemctl start postgresql.service # возвращаем системный постгрес на место
  exit 0
}

trap "exitfn" INT

# готовим базу данных
sudo systemctl stop postgresql.service # останавливаем системный постгрес, чтобы он не занимал порт 5432
cd database
/bin/bash cleanup.bash   # запускаем скрипт для генерации скрипта для заполнения бд

# запускаем докер-контейнер
cd ..
if [[ $* == *--build* ]]; then
  sudo docker-compose up --build &> "$docker_output_path"
else
  sudo docker-compose up &> "$docker_output_path"
fi
sudo docker-compose down &> "$docker_output_path"
sudo systemctl start postgresql.service # возвращаем системный постгрес на место
