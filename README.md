# L0 WB task
В БД:

1.Развернуть локально postgresql  
2.Создать свою бд    
3.Настроить своего пользователя.     
4.Создать таблицы для хранения полученных данных. 

В сервисе:
1. Подключение и подписка на канал в nats-streaming 
2. Полученные данные писать в Postgres  
3. Так же полученные данные сохранить in memory в сервисе (Кеш) 
4. В случае падения сервиса восстанавливать Кеш из Postgres  
5. Поднять http сервер и выдавать данные по id из кеша 
6. Сделать простейший интерфейс отображения полученных данных, для 
их запроса по id


Для запуска:
1. make run - запускает все нужные контейнеры (nats-streaming и postgres), если таблицы в бд не созданы - создает их  и запускает сервер
2. make publish - запускает паблишера для отправки рандомно сгенерированных данных в nats-streaming  
localhost:8080/api/orders - для просмотра простого интерфейса для получения информации о заказах  