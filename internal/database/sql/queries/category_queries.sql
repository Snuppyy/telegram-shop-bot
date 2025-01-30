-- Создать категорию
INSERT INTO categories (name, parent_id, created_at, updated_at)
VALUES ($1, $2, NOW(), NOW()) RETURNING id;

-- Обновить категорию
UPDATE categories
SET name = $1, parent_id = $2, updated_at = NOW()
WHERE id = $3;

-- Получить категорию по ID
SELECT id, name, parent_id, created_at, updated_at
FROM categories WHERE id = $1;

-- Получить все категории
SELECT id, name, parent_id, created_at, updated_at FROM categories;

-- Удалить категорию
DELETE FROM categories WHERE id = $1;