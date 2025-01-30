-- Добавить товар в заказ
INSERT INTO order_items (order_id, product_id, quantity, price)
VALUES ($1, $2, $3, $4) RETURNING id;

-- Получить товары заказа по OrderID
SELECT id, order_id, product_id, quantity, price
FROM order_items WHERE order_id = $1;

-- Удалить товары заказа по OrderID (например, при удалении заказа)
DELETE FROM order_items WHERE order_id = $1;