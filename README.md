# seabotserver

This is tcp server for seabattle game. Players are AI bots


# tcp protocol
```
<- client to server
-> server to client

[0 0 0 20]{"dkfhsjdfh"}

```
## auth
```json
<- { "auth" : "12334yger5348fhf8d7tdg8s76g" }
-> { "auth" : 
		{ 
			"ok": false, 
			"error": "some error",
			"id" : 123  
		}
	}
```
```
<- { "exit" : true }
-> disconnect
```

```
// bot versus bot
// сервер сам расставляет корабли
<- { "bvb" : { "place": 0 } } 
// игрок расставляет корабли на поле
// сервер должен проверить и допустить или не допустить расстановку
<- { "bvb" : { "place":  1, "ships" : 
		{
			"ship4" :  [ [0, 0],[0, 1],[0, 2],[0, 3] ], 
			"ship31" : [ [0, 0],[0, 1],[0, 2] ], 
			"ship32" : [ [0, 0],[0, 1],[0, 2] ], 
			"ship21" : [ [0, 0],[0, 1] ], 
			"ship22" : [ [0, 0],[0, 1] ], 
			"ship23" : [ [0, 0],[0, 1] ], 
			"ship11" : [ [0, 0] ], 
			"ship12" : [ [0, 0] ], 
			"ship13" : [ [0, 0] ], 
			"ship14" : [ [0, 0] ]
		}
	}
}
// no bots yet, wait
-> { "bvb" : { "wait": 1 } }
// start battle, opponent bot info
-> { "bvb" : { "id": 321, "name": "dopinfo", "ships": [0,0...] } }
0 0 0 0 0 0 0 0 0 1
0 4 0 0 0 0 0 0 0 0
0 4 0 0 0 0 3 0 0 2
0 4 0 0 0 0 3 0 0 2
0 4 0 0 0 0 3 0 0 0
0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0

```
##### LOOP: GAME
```
// сервер предлагает игроку 123 сделать ход
-> { "turn" : { "id" : 123 } }

// Бот сделал ход, выстрел пришелся в точку (А2) - 0,0
// АБВГДЕ....
// 0123456789
// первая цифра это номер ряда, вторая цифра номер колонки
<- { "turn" : { "shot": [0, 1] } }

// результат выстрела, -1 - мимо, 1 - попал, 2 - убил
-> { "turn" : { "result": 0 } }

-> { "turn" : {"opponent": { "shot": [0, 1], "result": 1 } } }<!>
```
##### GOTO GAME
```
// 10 second timeout
// after timeout -> lose
```
## Battle end
```
-> { "end": { "winner": 123, "opponent": {...ships...} } }
```
# DB structure draft
```
User
	ID
	email
	pass
	name
	
Bot
	ID айди бота
	User
	AuthKey
	
Tournament
	ID
	Type - тип турнира
		SandBox - все боты имеют доступ в песочницу
		Tournament - после отборочного тура админ переносит в турнир
		Quality - игрок подает заявку в диапазоне дат, пишем запись в таблицу TourBot
	Name - Имя турнира
	RegStart
	RegUntil


TourBot
	Bot
	Tour
	State: Access, Deny
	RegisteredDate
	Played - сыграно боев
	Win - побед
	Lose - поражений
	Disconnect - дисконектов


```