package stores

import "zakroma_backend/schemas"

func GetDishWithId(id int) (schemas.Dish, error) {
	db, err := CreateConnection()
	if err != nil {
		return schemas.Dish{}, err
	}

	var dish schemas.Dish
	err = db.
		QueryRow(`
			select
				dish_id,
				dish_name,
				calories,
				proteins,
				fats,
				carbs,
				receipt
			from
			    dishes
			where
			    dish_id = $1`,
			id).
		Scan(&dish.Id,
			&dish.Name,
			&dish.Calories,
			&dish.Proteins,
			&dish.Fats,
			&dish.Carbs,
			&dish.Recipe)
	if err != nil {
		return schemas.Dish{}, err
	}

	productsRows, err := db.
		Query(`
			select
				products.product_id,
				products.product_name,				
				products_dishes.amount,
				products.unit_of_measurement
			from
			    products_dishes,
			    products
			where
			    products_dishes.dish_id = $1 and
			    products_dishes.product_id = products.product_id`,
			id)
	if err != nil {
		return schemas.Dish{}, err
	}

	defer productsRows.Close()

	for productsRows.Next() {
		var product schemas.DishProduct
		if err = productsRows.
			Scan(&product.ProductId,
				&product.Name,
				&product.Amount,
				&product.UnitOfMeasurement); err != nil {
			return schemas.Dish{}, err
		}
		dish.Products = append(dish.Products, product)
	}

	return dish, nil
}

func GetDishShortWithId(id int) (schemas.Dish, error) {
	db, err := CreateConnection()
	if err != nil {
		return schemas.Dish{}, err
	}

	var dish schemas.Dish
	err = db.
		QueryRow(`
			select
				dish_id,
				dish_name,
				calories,
				proteins,
				fats,
				carbs
			from
			    dishes
			where
			    dish_id = $1`,
			id).
		Scan(&dish.Id,
			&dish.Name,
			&dish.Calories,
			&dish.Proteins,
			&dish.Fats,
			&dish.Carbs)
	if err != nil {
		return schemas.Dish{}, err
	}

	return dish, nil
}

func checkMatch(name string, pattern string) bool {
	for i := 0; i < len(name)-len(pattern)+1; i++ {
		if name[i:i+len(pattern)] == pattern {
			return true
		}
	}
	return false
}

func GetDishesShortWithName(name string, rangeBegin int, rangeEnd int) []schemas.Dish {
	db, err := CreateConnection()
	if err != nil {
		return make([]schemas.Dish, 0)
	}

	var matchedDishes []int
	order := 0
	rows, err := db.Query(`
			select
				dishes.dish_id,
				dishes.dish_name
			from
			    dishes`)

	if err != nil {
		rows.Close()
		return make([]schemas.Dish, 0)
	}

	for rows.Next() {
		var dishId int
		var dishName string
		if err = rows.Scan(&dishId, &dishName); err != nil {
			return make([]schemas.Dish, 0)
		}

		if !checkMatch(dishName, name) {
			continue
		}

		order += 1
		if order > rangeEnd {
			break
		}
		if rangeBegin <= order {
			matchedDishes = append(matchedDishes, dishId)
		}
	}
	rows.Close()

	dishes := make([]schemas.Dish, 0)
	for id := range matchedDishes {
		dish, err := GetDishShortWithId(id)
		if err != nil {
			return make([]schemas.Dish, 0)
		}

		dishes = append(dishes, dish)
	}

	return dishes
}

func GetDishesShortWithTags(tags []string) []schemas.Dish {
	db, err := CreateConnection()
	if err != nil {
		return make([]schemas.Dish, 0)
	}

	var tagsId []int
	for i := range tags {
		var id int
		if err = db.QueryRow(`select tag_id from tags where tag = $1`, tags[i]).Scan(&id); err != nil {
			return make([]schemas.Dish, 0)
		}
		tagsId = append(tagsId, id)
	}

	var cnt = map[int]int{}
	for i := range tagsId {
		rows, err := db.Query(`
			select
				dishes_tags.dish_id
			from
			    dishes,
			    dishes_tags
			where
			    dishes_tags.tag_id = $1`, tagsId[i])

		if err != nil {
			rows.Close()
			return make([]schemas.Dish, 0)
		}

		for rows.Next() {
			var id int
			if err = rows.Scan(&id); err != nil {
				return make([]schemas.Dish, 0)
			}
			cnt[id] += 1
		}
		rows.Close()
	}

	var dishes []schemas.Dish
	for id, matched := range cnt {
		if matched == len(tagsId) {
			dish, err := GetDishShortWithId(id)
			if err != nil {
				return make([]schemas.Dish, 0)
			}

			dishes = append(dishes, dish)
		}
	}

	return dishes
}
