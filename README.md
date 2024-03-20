# FilmAPI

FilmAPI - это REST API сервис для управления базой данных фильмов и актеров, представляющий бэкенд приложения "Фильмотека".

## Описание

FilmAPI позволяет пользователям добавлять, получать, изменять и удалять информацию о фильмах и актерах. В приложении реализована авторизация с использованием JWT и разделением ролей пользователей. Администратору доступны все действия, а пользователю - только получение информации о фильмах и актерах.

Язык реализации приложения - Go, используется база данных PostgreSQL.


<!-- ![Go](images/go.png){width=200px} -->
<img src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png" border="2% solid red" width="8.1%"/><img src="https://upload.wikimedia.org/wikipedia/commons/2/29/Postgresql_elephant.svg" border="2% solid red" width="6.9%"/><img src="https://upload.wikimedia.org/wikipedia/commons/e/ea/Docker_%28container_engine%29_logo_%28cropped%29.png" border="2% solid red" width="12.45%"/><img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRhS7TNcQhzz7g5vb5AGQpKM42slfGfZC56yM5y47-ecw&s" border="2% solid red" width="8%"/><img src="https://git-scm.com/images/logos/downloads/Git-Logo-2Color.png" border="2% solid red" width="17%"/>



## Запуск

1. Клонируйте репозиторий:
```sh
git clone https://github.com/WtoJglg4/filmAPI.git && cd filmAPI
```

2. Убедитесь, что у вас включен Docker Engine.

3. Запустите скрипт:

```sh
./run.sh
```

## Документация API

Получите подробную документацию по API, перейдя по ссылке http://localhost:3000/swagger/ после запуска приложения.


## Автор

Glazov Vadim, WtoJglg4