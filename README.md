### Запуск приложения

cd deploy  
docker-compose up --build

## Задания

### Сервис отчетов
1. Взаимодействует с базой данных в read-only формате  
2. Должен предоставлять данные о том как студенты прошли тесты

### Система для e2e тестирования
1. Нужно поднимать пустую базу данных (либо наполенную какими-то тестовыми данными) 
2. После чего запускать тесты) 
3. Можно сделать на каком-нибудь питоновском фрейсворке

### Сервис админка для создания лабораторных работ
1. Создать сервис для создания лабораторных работ
2. Записывает данные в БД
3. Нужно подумать над проверкой ответа на тест

### Мониторинг сервиса (это мой диплом)
1. Записывать ошибки в хранилище
2. Настроить алерты
3. Настроить графики