# TimerTracker

Бекенд для TimerTracker, позволяет менеджерам контролировать работу своей команды и получать информацию о том, сколько времени было потрачено на ту или иную задачу.

## Технологии, которые были использованы
1. Язык программирование: Golang
2. Фреймворк: Fiber
3. База данных: PostgreSQL
4. testify - для юнит тестов
5. env - для конфигурации
6. GORM - в качестве авто миграции и как ORM
7. Swagger - для документации запросов
8. Makefile - для сокращение команд

## Документация
Чтобы получить документацию по запросам, необходимо развернуть данный сервис и зайти по ссылки: http://localhost:8083(ВАШ_URL)/swagger


## Запуск \ деплой
Для деплоя, необходимо получить текущий репозиторий использовать Docker (в данном проекте нету, Dockerfile), для запуска достаточно ввести любую из команд
```makefile
git clone <ссылка>
make run
go run .
```
## Чего не сделано
Данный код не является идеальным, однако старался разделись: сервис логику, обработчики, бизнес-логику и репозитории.
```
- Доделать юнит тесты и добавить mock тесты
- Добавить сбор метрик Grafana \ Prometheus
- По метрикам определять, если база данных будет испытывать нагрузку, нужно кэшировать данные используя Redis или Memcache
- Добавить Dockerfile
- В качестве масштабирование если продукт расширется и будут больше функциональностей и монолит вырастит, рассмотреть 
  гибридный подход к архитектуре или микросервисный
- До делать дебаг логи
```
