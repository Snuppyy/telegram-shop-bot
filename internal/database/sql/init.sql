-- -- Создание схемы (если необходимо)
-- CREATE SCHEMA IF NOT EXISTS myapp;
--
-- -- Настройка начальных переменных, если нужно
-- -- Например, создать роль базы данных, если используются ограничения доступа
-- DO $$
-- BEGIN
--     CREATE ROLE myapp_user;
-- EXCEPTION
--     WHEN DUPLICATE_OBJECT THEN NULL; -- Роль уже существует, ничего не делаем
-- END $$;