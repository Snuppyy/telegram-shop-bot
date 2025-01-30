-- Заполнение таблицы пользователей
INSERT INTO users (username, email, password, role, created_at, updated_at)
VALUES
('admin', 'admin@example.com', 'hashed_password_1', 'admin', NOW(), NOW()),
('user', 'user@example.com', 'hashed_password_2', 'user', NOW(), NOW());

-- Заполнение таблицы категорий
INSERT INTO categories (name, parent_id, created_at, updated_at)
VALUES
('Electronics', NULL, NOW(), NOW()),
('Books', NULL, NOW(), NOW()),
('Phones', 1, NOW(), NOW()),
('Laptops', 1, NOW(), NOW());

-- Заполнение таблицы продуктов
INSERT INTO products (category_id, name, description, price, stock, image_url, created_at, updated_at)
VALUES
(3, 'iPhone 14', 'Latest Apple smartphone', 999.99, 50, 'https://example.com/iphone14.jpg', NOW(), NOW()),
(4, 'MacBook Pro', 'Powerful laptop by Apple', 1999.99, 20, 'https://example.com/macbookpro.jpg', NOW(), NOW());