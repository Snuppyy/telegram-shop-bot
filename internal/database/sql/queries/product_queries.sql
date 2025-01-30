-- Создать продукт
INSERT INTO products (category_id, name, description, price, stock, image_url, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW()) RETURNING id;

-- Обновить продукт
UPDATE products
SET category_id = $1, name = $2, description = $3, price = $4, stock = $5, image_url = $6, updated_at = NOW()
WHERE id = $7;

-- Получить продукт по ID
SELECT id, category_id, name, description, price, stock, image_url, created_at, updated_at
FROM products WHERE id = $1;

-- Получить список всех продуктов
SELECT id, category_id, name, description, price, stock, image_url, created_at, updated_at FROM products;

-- Удалить продукт
DELETE FROM products WHERE id = $1;