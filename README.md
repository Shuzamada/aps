Онлайн агентство управляет заявками на услуги по организации мероприятий: свадьбы, корпоративы, дни рождения...

Вариант **23**: 10. ИБ ИЗ1 ПЗ2 Д10З2 Д10О5 Д2П1 Д2Б5 ОР1 ОД2

- ИБ — бесконечный источник
- ИЗ1 — пуассоновский источник для бесконечных, экспоненциальная задержка для конечных
- ПЗ2 — равномерный прибор
- Д1ОЗ2 — дисциплины буферизации в порядке поступления
- Д1ОО5 — дисциплины отказа - вновь пришедшая
- Д2П1 — приоритет дисциплин постановки по номеру прибора
- Д2Б5 — приоритет дисциплин постановки  по номеру источника, заявки в пакете
- ОД2 — формализованная схема модели, текущее состояние
- ОР1 — сводная таблица результатов


Блок-схема

![](https://www.plantuml.com/plantuml/svg/ZLB1aX8n5BpdAvwxbzrV60ukebPP9OXmy9mbFH1zaup9or2KvR-R6QQMWTgeEG0XdKxtwtwAeeEaXjKREQ01jQojK_67TQqEkY6s3d9gGJx8mIiru5MrO0idjErtCwKoLgnOx-XkthPiTthORDDpOsFCF-_ebwAyR4qFKyJ69q6WBNk9DE-qPVtCUFj8Se2ZXnFYew9nLOi0BQ0FNV_qX6LsWGmlld88Zg07vozizm6csk0xTWKkEAF6RrkhtrSctwMwHg1ry3jaVsu_zJSUVapC1ICRDBu3VM63M78oDuPMEBv-Uk-3TdjKAo-sFj4vZOg-wQpFxUcpZLBotHDFIKaydHNFYdQNskbz7nCOExc9ynp72FbdXcUAt0mhFOgkRXUp6ktsqiM_sFoQa70PIc4Q_vVeLIdHMHnBbLq24ZFcPx813Pp4WjLyAVWEfQQA3ZBQNxDsxafgjGx8wa3l8NcWC7tNMXyrJqgm3u2xSJqoBhohRMXWYYCzWTaVgzx7K49ZXabefa-kzg6vrRk1f31fItdkAy3Gqh-8Ztz9eF_ZAN-aXly0)


**Диаграмма классов**

- **Client (Клиент)**:
    - Представляет физических или юридических лиц
    - Содержит базовую информацию о клиенте
    - Метод `SubmitApplication()` для подачи заявки
- **Application (Заявка)**:
    - Хранит детали мероприятия: тип, дата, количество гостей
    - Содержит специфические требования
    - Имеет метод валидации `Validate()`
- **Buffer (Буфер)**:
    - Реализует очередь заявок по принципу FIFO
    - Ограничен максимальным размером
    - Методы добавления, удаления и очистки старых заявок
- **ApplicationDispatcher (Диспетчер постановки)**:
    - Управляет буфером заявок
    - Распределяет заявки между менеджерами
    - Обрабатывает переполнение буфера
- **ManagerSelector (Диспетчер выборки)**:
    - Выбирает менеджера для обработки заявки
    - Рассчитывает приоритет менеджера на основе загрузки и квалификации
- **EventManager (Менеджер по организации мероприятий)**:
    - Имеет специализацию и уровень квалификации
    - Ограничен максимальной нагрузкой
    - Методы проверки возможности принять заявку и обработать её

![](./sequence.svg)



Сиквенс-диаграмма

- Пользователь подает заявку
- Диспетчер проверяет буфер
- Если место есть – находит преподавателя
- Если нет места – добавляет в очередь
- При переполнении буфера удаляются старые заявки

![](./classes.svg)


BPMN диаграмма

![](https://www.plantuml.com/plantuml/svg/TPBFajCm3CRlVWeTkw-G7jYom0mpB9JsmDbDbAr0iOijTwPC7XwbTJ9sAJCzf9Rq_JxzUOvYWwI1pVZ2X4v6StqnpJuT1XBuUDP7s9nEJAqLycvjcctmP-emnpwWMEAuDcOiqc3das0OJRF35kcCjGRri7gV-Xu3V4hCxvv0ZkXWTS7uP02qQF1QyuqxZ7TuT1trfdavOliBiiRch1RiWPZaCXdP1l91ltDyB7bVlJncRhgk987uNBxMgFwgzXR27XbRySfGYTiVs2SIU-1vGg_MsRT5PhAINsRVZFTcb5RLJkgJHMVvyWShcyly0QxwDqAPyOee4NRuEr70GLNZlpY0HzlBeKP3mT_jCCUMkycfEbxgBqrkVSlMEN9li9-kAfj4ZWWhvsTolQu2kqD9mJrUol_uKvlNc-awcYeimsVN-p0KhcydL1jvR2EVXL9_M9AfaUy94wvDZIuRdnE_K1JDox0IjTGLLTNmyIUsWjrywr7yoJphTnhuBm00)
