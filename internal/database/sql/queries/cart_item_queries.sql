-- Получить элементы корзины по CartID
SELECT id, cart_id, product_id, quantity, price
FROM cart_items WHERE cart_id = $1;

-- Добавить элемент в корзину
INSERT INTO cart_items (cart_id, product_id, quantity, price)
VALUES ($1, $2, $3, $4) RETURNING id;

-- Обновить количество товара в корзине
UPDATE cart_items
SET quantity = $1, price = $2
WHERE id = $3;

-- Удалить элемент корзины
DELETE FROM cart_items WHERE id = $1;

-- Удалить все элементы корзины (например, после создания заказа)
DELETE FROM cart_items WHERE cart_id = $1;