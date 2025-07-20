# TeamFinder API - Postman Collection

Полная коллекция для тестирования API TeamFinder с автоматизированными тестами и переменными.

## Установка

1. **Скачайте коллекцию:**
   - Импортируйте файл `TeamFinder_Postman_Collection.json` в Postman

2. **Настройте переменные:**
   - Коллекция содержит предустановленные переменные
   - При необходимости измените `base_url` (по умолчанию: `http://localhost:8090`)

## Структура коллекции

### System
- **Health Check** - проверка статуса сервера

### Authentication
- **Register** - регистрация нового пользователя
- **Login** - вход в систему
- **Send Email Verification Code** - отправка кода верификации
- **Verify Email Code** - подтверждение email с кодом
- **Get Telegram Auth URL** - получение ссылки для Telegram авторизации
- **Verify Telegram Auth** - подтверждение Telegram авторизации
- **Check Token** - проверка валидности токена
- **Refresh Token** - обновление токена
- **Logout** - выход из системы
- **Delete Account** - удаление аккаунта

### Hackathons
- **Create Hackathon** - создание хакатона
- **Get All Hackathons** - получение всех хакатонов
- **Get Hackathon by ID** - получение хакатона по ID
- **Update Hackathon** - обновление хакатона
- **Delete Hackathon** - удаление хакатона

### Profiles
- **Create Profile** - создание профиля участника
- **Get Profiles by Hackathon ID** - получение профилей для хакатона
- **Get Profile by ID** - получение профиля по ID
- **Update Profile** - обновление профиля
- **Delete Profile** - удаление профиля

## Автоматизация

### Автоматическое извлечение токенов
Коллекция автоматически извлекает и сохраняет токены после авторизации:
- `access_token` - JWT токен доступа
- `refresh_token` - токен обновления

### Автоматическое извлечение ID
Коллекция автоматически сохраняет ID созданных сущностей:
- `hackathon_id` - ID созданного хакатона
- `profile_id` - ID созданного профиля

## Быстрое тестирование

### Полный workflow:
1. **Health Check** - проверьте, что сервер запущен
2. **Register** - зарегистрируйте пользователя (токены сохранятся автоматически)
3. **Create Hackathon** - создайте хакатон (ID сохранится автоматически)
4. **Create Profile** - создайте профиль для хакатона
5. **Get All Hackathons** - проверьте созданные хакатоны
6. **Get Profiles by Hackathon ID** - проверьте профили

### Тестирование Email верификации:
1. **Send Email Verification Code**
2. Получите код из email (или логов сервера)
3. Обновите переменную `verification_code`
4. **Verify Email Code**

## Переменные коллекции

### Основные:
- `base_url` - адрес сервера (http://localhost:8090)
- `access_token` - JWT токен (автоматически заполняется)
- `refresh_token` - токен обновления (автоматически заполняется)

### Для тестирования:
- `test_username` - имя пользователя для тестов
- `test_email` - email для тестов
- `test_password` - пароль для тестов
- `verification_code` - код верификации email

### Данные хакатона:
- `hackathon_name` - название хакатона
- `hackathon_description` - описание хакатона
- `hackathon_start_date` - дата начала
- `hackathon_end_date` - дата окончания

### Данные профиля:
- `profile_name` - имя участника
- `profile_surname` - фамилия участника
- `profile_desired_role` - желаемая роль
- `profile_skills` - навыки (JSON массив)
- `profile_achievements` - достижения (JSON массив)

## Примеры ответов

### Health Check:
```json
{
  "status": "healthy",
  "timestamp": "2025-07-20T14:36:59Z",
  "service": "teamfinder-backend",
  "version": "1.0.0"
}
```

### Успешная регистрация:
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Создание хакатона:
```json
{
  "id": 11,
  "message": "Hackathon created successfully"
}
```

### Создание профиля:
```json
{
  "id": 25,
  "message": "Profile created successfully"
}
```

## Устранение неполадок

### Сервер недоступен:
- Проверьте, что Docker контейнеры запущены: `docker-compose ps`
- Проверьте Health Check endpoint

### 401 Unauthorized:
- Убедитесь, что токен получен через Register/Login
- Проверьте переменную `access_token`

### 404 Not Found:
- Проверьте правильность URL
- Убедитесь, что используете правильные ID

### Ошибки создания профиля:
- Проверьте, что пользователь существует (user_id)
- Убедитесь, что хакатон существует (hackathon_id)
- Помните: один пользователь = один профиль на хакатон

## Дополнительная информация

- API использует JWT для аутентификации
- Все даты в формате ISO 8601 UTC
- Навыки и достижения передаются как JSON массивы
- Профили привязаны к конкретным хакатонам и пользователям 