CREATE TABLE "users" (
  "user_id" serial PRIMARY KEY,
  "user_name" varchar(32),
  "password_hash" varchar(64),
  "birth_date" date,
  "user_hash" varchar(64)
);

CREATE TABLE "users_groups" (
  "user_id" integer,
  "group_id" integer,
  "role" varchar(16),
  PRIMARY KEY ("user_id", "group_id")
);

CREATE TABLE "groups" (
  "group_id" serial PRIMARY KEY,
  "group_name" varchar(32),
  "current_diet_id" integer,
  "group_hash" varchar(64)
);

CREATE TABLE "dishes" (
  "dish_id" serial PRIMARY KEY,
  "dish_name" varchar(64),
  "recipe" varchar(2048),
  "proteins" numeric,
  "carbs" numeric,
  "fats" numeric,
  "calories" numeric,
  "image_path" varchar(64),
  "dish_hash" varchar(64)
);

CREATE TABLE "dishes_tags" (
  "dish_id" integer,
  "tag_id" integer,
  PRIMARY KEY ("dish_id", "tag_id")
);

CREATE TABLE "tags_for_dishes" (
  "tag_id" serial PRIMARY KEY,
  "tag" varchar(32)
);

CREATE TABLE "meals" (
  "meal_id" serial PRIMARY KEY,
  "meal_name" varchar(64),
  "meal_hash" varchar(64),
  "tag_id" integer
);

CREATE TABLE "meals_dishes" (
  "meal_id" integer,
  "dish_id" integer,
  "portions" numeric,
  "accepted" bool,
  "author_hash" varchar(64),
  PRIMARY KEY ("meal_id", "dish_id")
);

CREATE TABLE "diet" (
  "diet_id" serial PRIMARY KEY,
  "diet_name" varchar(64),
  "diet_hash" varchar(64)
);

CREATE TABLE "groups_diets" (
  "group_id" integer,
  "diet_id" integer,
  PRIMARY KEY ("group_id", "diet_id")
);

CREATE TABLE "products" (
  "product_id" serial PRIMARY KEY,
  "product_name" varchar(64),
  "proteins" numeric,
  "carbs" numeric,
  "fats" numeric,
  "calories" numeric,
  "unit_of_measurement" varchar(8)
);

CREATE TABLE "products_dishes" (
  "product_id" integer,
  "dish_id" integer,
  "amount" integer,
  PRIMARY KEY ("product_id", "dish_id")
);

CREATE TABLE "groups_products" (
  "product_id" integer,
  "group_id" integer,
  "amount" integer,
  "expiration_date" date,
  PRIMARY KEY ("product_id", "group_id")
);

CREATE TABLE "diet_day_meals" (
  "diet_day_id" integer,
  "meal_id" integer,
  "index" integer,
  PRIMARY KEY ("diet_day_id", "meal_id")
);

CREATE TABLE "diet_day_diet" (
  "diet_id" integer,
  "diet_day_id" integer,
  "index" integer,
  PRIMARY KEY ("diet_id", "diet_day_id")
);

CREATE TABLE "groups_dishes_created_by_groups" (
  "dish_id" integer,
  "group_id" integer,
  PRIMARY KEY ("dish_id", "group_id")
);

CREATE TABLE "groups_diet_created_by_groups" (
  "group_id" integer,
  "diet_id" integer,
  PRIMARY KEY ("group_id", "diet_id")
);

CREATE TABLE "diet_day" (
  "diet_day_id" serial PRIMARY KEY,
  "diet_day_name" varchar(64)
);

CREATE TABLE "tags_for_meals" (
  "tag_id" serial PRIMARY KEY,
  "tag" varchar(32)
);

CREATE TABLE "diet_tags" (
  "diet_id" integer,
  "tag_id" integer,
  PRIMARY KEY ("diet_id", "tag_id")
);

CREATE TABLE "tags_for_diet" (
  "tag_id" serial PRIMARY KEY,
  "tag" varchar(32)
);

ALTER TABLE "users_groups" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "users_groups" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("group_id");

ALTER TABLE "products_dishes" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("product_id");

ALTER TABLE "products_dishes" ADD FOREIGN KEY ("dish_id") REFERENCES "dishes" ("dish_id");

ALTER TABLE "groups_products" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("group_id");

ALTER TABLE "groups_products" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("product_id");

ALTER TABLE "groups_dishes_created_by_groups" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("group_id");

ALTER TABLE "groups_dishes_created_by_groups" ADD FOREIGN KEY ("dish_id") REFERENCES "dishes" ("dish_id");

ALTER TABLE "dishes_tags" ADD FOREIGN KEY ("dish_id") REFERENCES "dishes" ("dish_id");

ALTER TABLE "dishes_tags" ADD FOREIGN KEY ("tag_id") REFERENCES "tags_for_dishes" ("tag_id");

ALTER TABLE "groups_diets" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("group_id");

ALTER TABLE "meals_dishes" ADD FOREIGN KEY ("dish_id") REFERENCES "dishes" ("dish_id");

ALTER TABLE "meals_dishes" ADD FOREIGN KEY ("meal_id") REFERENCES "meals" ("meal_id");

ALTER TABLE "groups_diets" ADD FOREIGN KEY ("diet_id") REFERENCES "diet" ("diet_id");

ALTER TABLE "diet_day_meals" ADD FOREIGN KEY ("meal_id") REFERENCES "meals" ("meal_id");

ALTER TABLE "diet_day_diet" ADD FOREIGN KEY ("diet_id") REFERENCES "diet" ("diet_id");

ALTER TABLE "groups_diet_created_by_groups" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("group_id");

ALTER TABLE "groups_diet_created_by_groups" ADD FOREIGN KEY ("diet_id") REFERENCES "diet" ("diet_id");

ALTER TABLE "diet_day_diet" ADD FOREIGN KEY ("diet_day_id") REFERENCES "diet_day" ("diet_day_id");

ALTER TABLE "diet_day_meals" ADD FOREIGN KEY ("diet_day_id") REFERENCES "diet_day" ("diet_day_id");

ALTER TABLE "diet_tags" ADD FOREIGN KEY ("tag_id") REFERENCES "tags_for_diet" ("tag_id");

ALTER TABLE "diet_tags" ADD FOREIGN KEY ("diet_id") REFERENCES "diet" ("diet_id");

ALTER TABLE "meals" ADD FOREIGN KEY ("tag_id") REFERENCES "tags_for_meals" ("tag_id");