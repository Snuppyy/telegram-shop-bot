-- Создать заказ
INSERT INTO orders (user_id, total_amount, status, created_at, updated_at)
VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id;

-- Обновить статус заказа
UPDATE orders SET status = $1, updated_at = NOW() WHERE id = $2;

-- Получить заказы пользователя
SELECT id, user_id, total_amount, status, created_at, updated_at
FROM orders WHERE user_id = $1;

-- Получить заказ по ID
SELECT id, user_id, total_amount, status, created_at, updated_at
FROM orders WHERE id = $1;

-- Удалить заказ
DELETE FROM orders WHERE id = $1;