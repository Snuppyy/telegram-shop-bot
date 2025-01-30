-- Создать новый адрес
INSERT INTO addresses (user_id, country, city, street, building, apartment, zip_code, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW()) RETURNING id;

-- Обновить адрес
UPDATE addresses
SET country = $1, city = $2, street = $3, building = $4, apartment = $5, zip_code = $6, updated_at = NOW()
WHERE id = $7;

-- Получить адреса пользователя
SELECT id, user_id, country, city, street, building, apartment, zip_code, created_at, updated_at
FROM addresses WHERE user_id = $1;

-- Удалить адрес
DELETE FROM addresses WHERE id = $1;