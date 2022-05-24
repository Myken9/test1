Сервис предназначен для поиска в БД названия книги по ее автору и авторов
по названию книги.


Запустите контейнер с mysql:
```bash
docker-compose up -d
```

При разворачивании приложения подключитесь к контейнеру и выполните миграции
```bash
docker-compose exec db /bin/bash
mysql -u user -ppass db < /schema.sql
exit
```

