Общие сведения, которым необходимо следовать:

`[info]`      Информация

`[strictly]`  Строгие ограничения

`[!]`         Организация слоёв или описание

`[link]`      Ссылки

`[architecture]` Организация архитектуры

<br>

### [architecture] <br>
    `/cmd`
- `main.go`

<br>

`/internal`
- `/app`
- `/handler`
- `/entity`
- `/infrastructure`
    - `/repository`
- `/usecase`

<br>

`/pkg`
- `/logging`
- `/utils`

### [!architecture]

<br>

`[info]`  Модули верхних уровней не зависят от нижних уровней

Внутренние слои:

`[!]` Entity

`[!]` UseCase

Не зависят от слоев:

`[!]` Controller

`[!]` Infrastructure - это слой с абстракциями

`[!]` UseCase - логика приложения

`[info]` Entity не зависят от вн.слоев
    (Infrastructure,UseCase)

`[info]` Entity - это бизнес логика, которая редко меняется
	(структура с методами)

`[info]` Вместо dao/dto использовать структуры Entity

`[info]` /internal 

`[info]` в /pkg всё что может быть переиспользовано

`[strictly]` Не должно быть импортов из других пакетов проекта, нужно через интерфейсы
 
`[strictly]` http framework не должен выходить за слой controller

`[strictly]` Не использовать ORM

`[strictly]` Не забывать про duck typing, интерфейсы хранить отдельным пакетом, 
                либо на стороне consumer интерфейса.

`[link]`  Использовал Чистую архитектуру под микросервис
        (https://www.youtube.com/watch?v=V6lQG6d5LgU), 
        также придерживался мыслям по чистой архитектуре от 
        (https://www.youtube.com/watch?v=dyvYXidvc8g&t=1064s) и
        (https://www.youtube.com/watch?v=mesl2Si6saw).