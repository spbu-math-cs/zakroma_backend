package stores

import (
	"sort"
	"zakroma_backend/schemas"
	"zakroma_backend/utils"
)

var DISHES_BY_ID_CACHE_SIZE = 100000
var DishesByIdCache = map[int]schemas.Dish{}
var DishesByIdCacheQueue = utils.CreateQueue(DISHES_BY_ID_CACHE_SIZE)

var DISHES_BY_HASH_CACHE_SIZE = 1000
var DishesByHashCache = map[string]schemas.Dish{}
var DishesByHashCacheQueue = utils.CreateQueue(DISHES_BY_HASH_CACHE_SIZE)

func GetDishIdByHash(hash string) (int, error) {
	db, err := CreateConnection()
	if err == nil {
		defer db.Close()
	}
	if err != nil {
		return -1, err
	}

	var dishId int
	if err = db.QueryRow(`
		select
		    dish_id
		from
			dishes
		where
			dish_hash = $1`,
		hash).Scan(
		&dishId); err != nil {
		return -1, err
	}

	return dishId, nil
}

func GetDishByHash(hash string) (schemas.Dish, error) {
	db, err := CreateConnection()
	if err == nil {
		defer db.Close()
	}
	if err != nil {
		return schemas.Dish{}, err
	}

	dish, err := GetDishShortWithRecipeByHash(hash)
	if err != nil {
		return schemas.Dish{}, err
	}

	productsRows, err := db.Query(`
		select
			products.product_hash,
		from
			products_dishes,
			products
		where
			products_dishes.dish_id = $1 and
			products_dishes.product_id = products.product_id`,
		dish.Id)
	if err != nil {
		return schemas.Dish{}, err
	}
	defer productsRows.Close()

	for productsRows.Next() {
		var productHash string
		if err = productsRows.Scan(
			&productHash); err != nil {
			return schemas.Dish{}, err
		}
		dish.Products = append(dish.Products, productHash)
		//TODO (вроде сделал)
	}

	return dish, nil
}

func GetDishShortByHash(hash string) (schemas.Dish, error) {
	var dish schemas.Dish

	dish, found := DishesByHashCache[hash]
	if found {
		return dish, nil
	}

	db, err := CreateConnection()
	if err == nil {
		defer db.Close()
	}
	if err != nil {
		return schemas.Dish{}, err
	}

	if err = db.QueryRow(`
		select
		    dish_id,
		    dish_hash,
    		dish_name,
    		calories,
    		proteins,
    		fats,
    		carbs,
    		image_path
		from
			dishes
		where
			dish_hash = $1`,
		hash).Scan(
		&dish.Id,
		&dish.Hash,
		&dish.Name,
		&dish.Calories,
		&dish.Proteins,
		&dish.Fats,
		&dish.Carbs,
		&dish.ImagePath); err != nil {
		return schemas.Dish{}, err
	}

	DishesByHashCache[hash] = dish
	removedDish, flag := DishesByHashCacheQueue.Push(hash)
	if flag {
		delete(DishesByHashCache, removedDish.(string))
	}

	return dish, nil
}

func GetDishShortById(id int) (schemas.Dish, error) {
	var dish schemas.Dish

	dish, found := DishesByIdCache[id]
	if found {
		return dish, nil
	}

	db, err := CreateConnection()
	if err == nil {
		defer db.Close()
	}
	if err != nil {
		return schemas.Dish{}, err
	}

	if err = db.QueryRow(`
		select
		    dish_id,
		    dish_hash,
    		dish_name,
    		calories,
    		proteins,
    		fats,
    		carbs,
    		image_path
		from
			dishes
		where
			dish_id = $1`,
		id).Scan(
		&dish.Id,
		&dish.Hash,
		&dish.Name,
		&dish.Calories,
		&dish.Proteins,
		&dish.Fats,
		&dish.Carbs,
		&dish.ImagePath); err != nil {
		return schemas.Dish{}, err
	}

	DishesByIdCache[id] = dish
	removedDish, flag := DishesByIdCacheQueue.Push(id)
	if flag {
		delete(DishesByIdCache, removedDish.(int))
	}

	return dish, nil
}

func GetDishShortWithRecipeByHash(hash string) (schemas.Dish, error) {
	db, err := CreateConnection()
	if err == nil {
		defer db.Close()
	}
	if err != nil {
		return schemas.Dish{}, err
	}

	var dish schemas.Dish
	if err = db.QueryRow(`
		select
		    dish_id,
		    dish_hash,
    		dish_name,
    		calories,
    		proteins,
    		fats,
    		carbs,
    		image_path,
    		recipe
		from
			dishes
		where
			dish_hash = $1`,
		hash).Scan(
		&dish.Id,
		&dish.Hash,
		&dish.Name,
		&dish.Calories,
		&dish.Proteins,
		&dish.Fats,
		&dish.Carbs,
		&dish.ImagePath,
		&dish.Recipe); err != nil {
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

func GetDishesShortByName(name string, rangeBegin int, rangeEnd int) []schemas.Dish {
	db, err := CreateConnection()
	if err == nil {
		defer db.Close()
	}
	if err != nil {
		return make([]schemas.Dish, 0)
	}

	var matchedDishes []int
	order := 0
	dishesRows, err := db.Query(`
		select
			dish_id,
			dish_name
		from
			dishes
		order by
		    dish_id`)
	if err != nil {
		return make([]schemas.Dish, 0)
	}
	defer dishesRows.Close()

	for dishesRows.Next() {
		var dishId int
		var dishName string
		if err = dishesRows.Scan(
			&dishId,
			&dishName); err != nil {
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

	dishes := make([]schemas.Dish, 0)
	for i := range matchedDishes {
		dish, err := GetDishShortById(matchedDishes[i])
		if err != nil {
			return make([]schemas.Dish, 0)
		}
		dishes = append(dishes, dish)
	}

	return dishes
}

func GetDishesShortByTags(tags []string, rangeBegin int, rangeEnd int) []schemas.Dish {
	db, err := CreateConnection()
	if err == nil {
		defer db.Close()
	}
	if err != nil {
		return make([]schemas.Dish, 0)
	}

	order := 0
	var matchedDishes []int

	if len(tags) == 0 {
		dishesRows, err := db.Query(`
			select
				dish_id
			from
			    dishes
			order by
			    dish_id`)
		if err != nil {
			return make([]schemas.Dish, 0)
		}
		defer dishesRows.Close()

		for dishesRows.Next() {
			var dishId int
			if err = dishesRows.Scan(
				&dishId); err != nil {
				return make([]schemas.Dish, 0)
			}

			order += 1
			if order > rangeEnd {
				break
			}
			if rangeBegin <= order {
				matchedDishes = append(matchedDishes, dishId)
			}
		}
	} else {
		var tagsId []int
		for i := range tags {
			var id int
			if err = db.QueryRow(`
				select
				    tag_id
				from
				    tags_for_dishes
				where
				    tag = $1`,
				tags[i]).Scan(
				&id); err != nil {
				return make([]schemas.Dish, 0)
			}
			tagsId = append(tagsId, id)
		}

		var cnt = map[int]int{}
		for i := range tagsId {
			if err := func() error {
				dishesRows, err := db.Query(`
					select
						distinct(dishes_tags.dish_id)
					from
						dishes,
						dishes_tags
					where
						dishes_tags.tag_id = $1`,
					tagsId[i])
				if err != nil {
					return err
				}
				defer dishesRows.Close()

				for dishesRows.Next() {
					var dishId int
					if err = dishesRows.Scan(
						&dishId); err != nil {
						return err
					}
					cnt[dishId] += 1
				}

				return nil
			}(); err != nil {
				return make([]schemas.Dish, 0)
			}
		}

		for id, matched := range cnt {
			if matched == len(tagsId) {
				order += 1
				if order > rangeEnd {
					break
				}
				if rangeBegin <= order {
					matchedDishes = append(matchedDishes, id)
				}
			}
		}
	}

	sort.Ints(matchedDishes)

	dishes := make([]schemas.Dish, 0)
	for i := range matchedDishes {
		dish, err := GetDishShortById(matchedDishes[i])
		if err != nil {
			return make([]schemas.Dish, 0)
		}
		dishes = append(dishes, dish)
	}

	return dishes
}
