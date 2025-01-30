-- Создать пользователя
INSERT INTO users (username, email, password, role, created_at, updated_at)
VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING id;

-- Получить пользователя по ID
SELECT id, username, email, password, role, created_at, updated_at
FROM users WHERE id = $1;

-- Получить пользователя по Email
SELECT id, username, email, password, role, created_at, updated_at
FROM users WHERE email = $1;

-- Обновить данные пользователя
UPDATE users
SET email = $1, password = $2, updated_at = NOW()
WHERE id = $3;

-- Удалить пользователя
DELETE FROM users WHERE id = $1;