package stores

import "zakroma_backend/schemas"

func GetProductWithId(id int) (schemas.Product, error) {
	db, err := CreateConnection()
	if err != nil {
		return schemas.Product{}, err
	}

	var product schemas.Product
	err = db.
		QueryRow(`
			select
				product_id,
				product_name,
				calories,
				proteins,
				fats,
				carbs,
				unit_of_measurement
			from
			    products
			where
			    product_id = $1`,
			id).
		Scan(&product.Id,
			&product.Name,
			&product.Calories,
			&product.Proteins,
			&product.Fats,
			&product.Carbs,
			&product.UnitOfMeasurement)
	if err != nil {
		return schemas.Product{}, err
	}

	return product, nil
}
