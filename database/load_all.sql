-- Скрипт для автоматизации заполнения бд
-- !!!
-- !!! ВНИМАНИЕ! Запуск load_all.sql УДАЛИТ ВСЕ ТАБЛИЦЫ в целевой базе
-- !!!

DO $$ DECLARE
    r RECORD;
BEGIN
    FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = current_schema()) LOOP
            EXECUTE 'DROP TABLE IF EXISTS ' || quote_ident(r.tablename) || ' CASCADE';
        END LOOP;
END $$;

\ir scripts/0_DB_creation.sql
\ir scripts/1_Diet_table_filling.sql
\ir scripts/2_Dishes_table_filling.sql
\ir scripts/3_Groups_table_filling.sql
\ir scripts/4_Meals_table_filling.sql
\ir scripts/5_Products_table_filling.sql
\ir scripts/6_Tags_table_filling.sql
\ir scripts/7_Users_table_filling.sql
\ir scripts/8_Diet_day_table_filling.sql
\ir scripts/9_Diet_day_diet_table_filling.sql
\ir scripts/10_Diet_day_meals_table_filling.sql
\ir scripts/11_Dishes_tags_table_filling.sql
\ir scripts/12_Groups_diet_created_by_groups_table_filling.sql
\ir scripts/13_Groups_diets_table_filling.sql
\ir scripts/14_Groups_dishes_created_by_groups_table_filling.sql
\ir scripts/15_Groups_products_table_filling.sql
\ir scripts/16_Meals_dishes_table_filling.sql
\ir scripts/17_Products_dishes_table_filling.sql
\ir scripts/18_Users_groups_table_filling.sql
