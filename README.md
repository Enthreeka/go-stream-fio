# Сервис для обогащения данных


## Используемые библиотеки

[pgx/v5](https://github.com/jackc/pgx) - драйвер и инструмент для PostgreSQL

[fiber](https://github.com/gofiber/fiber) - http движок

[segmentio/kafka-go](https://github.com/segmentio/kafka-go) - инструмент для работы с kafka

[zap](https://github.com/uber-go/zap) - логгер

[go-redis/v9](https://github.com/redis/go-redis) - Redis клиент

[gofakeit/v6](https://github.com/brianvoe/gofakeit) - генерация случайнных данных

[godotenv](https://github.com/joho/godotenv) - загрузка данных с .env

## Инструкция по запуску

Развертывание dev среды:

`make docker-up`

Закрытие dev среды:

`make docker-down`

Запуск главного сервера:

`make server`

Запуск producer сервера:

`make producer`

Создание миграций:

`make migrate-up`

Дроп таблиц:

`make migrate-down`



## API


### Producer
`port:8081`

#### GET - отправить fio в kafka

`/`

Запрос:
````
{
"name": "Nikita",
"surname":"Novikov"
}
````

Ответ:
```
{
	"FIO": {
		"name": "Nikita",
		"surname": "Novikov"
	},
	"message": "sends"
}
```

### Главный сервер
`port:8080`

#### GET - получить пользователей по имени

`/users`

Запрос:

`localhost:8080/users?name=Arkady`

Ответ:

```
{
	"users": [
		{
			"id": "0afc2656-c03f-4e9e-b76c-4c273a1ec2bb",
			"firstname": "Arkady",
			"lastname": "Novikov",
			"birthday": [
				{
					"age": 100,
					"probability": 0.2
				}
			],
			"gender": [
				{
					"gender": "male",
					"probability": 0.8
				}
			],
			"address": [
				{
					"county_code": "MD",
					"probability": "0.200"
				},
				{
					"county_code": "BY",
					"probability": "0.200"
				},
				{
					"county_code": "AM",
					"probability": "0.200"
				},
				{
					"county_code": "TM",
					"probability": "0.200"
				},
				{
					"county_code": "UZ",
					"probability": "0.200"
				}
			]
		}
	]
}
```

#### DELETE - удаление пользователя

`/:id`

Запрос: 

```
{
	"id":"272ed579-630e-46fa-8da1-560a331b101c"
}
```

Ответ:

```
{
	"deleted id": {
		"id": "272ed579-630e-46fa-8da1-560a331b101c"
	},
	"message": "completed successfully"
}
```

#### POST
Запрос:

`/user`

```
{
	"name":"Shurik",
	"surname": "Kochnev"
}
```

Ответ:

```
{
	"created person": {
		"name": "Shurik",
		"surname": "Kochnev"
	},
	"message": "completed successfully"
}
```

Необходимо реализовать:

- [x] API с рандомными данными;
- [x] Реализовать Kafka
- [x] Реализовать Postgres
- [x] Реализовать Redis
- [x] Написать методы http
- [ ] Методы по GraphQL
- [ ] Unit тесты для usecase
- [x] Покрыть логи

Предложения по улучшению проекта:
- Добавить в docker compose migrate, а также создать dockerfile и также добавить в docker compose;
- Создать больше кастомных ошибок и сделать больше обработок;
- Чтобы привести к чистой архитектуре, работать с каждой сущностью из базы данных отдельно;
- Добавление Swagger для улучшения документации;
- Дописание некоторых методов, улучшение фильтрации, создание пагинации;
- Юнит тестирование;
- Забыл реализовать `FIO_FAILED` в случае ошибки.



Подробности по заданию:

v1:
~~В мною найденном API для получения случайных пользователей стоит ограничение в 1000 элементов. 
Из-за этого получается маленькая выборка с необходимыми `dto.fio`. В случае, если у всех найденных
пользователей одинаковое соотношение по возрасту, то берется возраст самого первого пользователя - 
`user.Age[0].Age`, который находится в `Age    []Age    `json:"birthday"``~~

v2: Создан собственный файл `fake_data.json` с фейковыми данными. Использовался пакет [gofakeit](https://github.com/brianvoe/gofakeit).
В данном файле 20000 сгенерированных пользователей, каждый четный пользователей генерировался 
пакетом `gofakeit`, каждый нечетный пользователь генерировал имя,фамилию и код случайным образом 
из уже определнного слайса. Для тестирования список всех имен и фамилий можно найти в 
`go-stream-fio/pkg/generatefakeusers.go`.

Пример генерации:

```go
countryCodes := []string{"AZ", "AM", "BY", "KZ", "KG", "MD", "RU", "TJ", "TM", "UZ"}
randomCountryCode := gofakeit.RandomString(countryCodes)
```
