Сервис предназначен для поиска в БД названия книги по ее автору и авторов
по названию книги.

# Настройка

Скопируйте .env.dist в .env и поправьте:
```bash
cp .env.dist .env
```

Запустите контейнер с mysql:
```bash
docker-compose up -d
```

При разворачивании приложения подключитесь к контейнеру и выполните миграции
```bash
docker-compose exec db /bin/bash
mysql -u user -ppass --default-character-set=utf8 db < /schema.sql
exit
```

# Запуск 
Удостоверьтесь, что контейнер с mysql запущен:
```bash
docker-compose up -d
```
Запустите приложение:
```bash
make run
```

# Запуск тестов
Удостоверьтесь, что контейнер с mysql запущен:
```bash
docker-compose up -d
```
Запустите тесты:
```bash
make test
```

Комментарии от разработчика:
- в ТЗ не было указанно, так что я сделал поиск по точному вхождению 