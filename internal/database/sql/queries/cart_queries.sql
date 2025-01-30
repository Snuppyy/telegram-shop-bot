-- Получить корзину пользователя по UserID
SELECT id, user_id, total, created_at, updated_at
FROM carts WHERE user_id = $1;

-- Создать корзину (для нового пользователя)
INSERT INTO carts (user_id, total, created_at, updated_at)
VALUES ($1, $2, NOW(), NOW()) RETURNING id;

-- Обновить корзину (например, пересчитать общую сумму)
UPDATE carts SET total = $1, updated_at = NOW() WHERE id = $2;

-- Удалить корзину (например, после заказа)
DELETE FROM carts WHERE id = $1;