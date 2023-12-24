package stores

import "zakroma_backend/schemas"

func GetGroupCartList(groupHash string) ([]schemas.DishProduct, error) {
	db, err := CreateConnection()
	if err == nil {
		defer db.Close()
	}
	if err != nil {
		return []schemas.DishProduct{}, err
	}

	productsRows, err := db.Query(`
		select
		    group_cart.product_id,
		    products.product_name,
		    group_cart.amount,
		    products.unit_of_measurement
		from
		    group_cart,
		    groups,
		    products
		where
		    group_cart.group_id = groups.group_id and
		    groups.group_hash = $1 and
		    group_cart.product_id = products.product_id
		order by
		    product_id`,
		groupHash)
	if err != nil {
		return []schemas.DishProduct{}, err
	}
	defer productsRows.Close()

	var products []schemas.DishProduct
	for productsRows.Next() {
		var product schemas.DishProduct
		if err := productsRows.Scan(
			&product.ProductId,
			&product.Name,
			&product.Amount,
			&product.UnitOfMeasurement); err != nil {
			return []schemas.DishProduct{}, err
		}

		products = append(products, product)
	}

	return products, nil
}

func AddGroupCartProduct(groupHash string, productId int, amount float32) error {
	db, err := CreateConnection()
	if err == nil {
		defer db.Close()
	}
	if err != nil {
		return err
	}

	if err := db.QueryRow(`
		insert into
			group_cart(group_id, product_id, amount)
		values
		    ((
		        select
		            group_id
		        from
		            groups
		        where
		            group_hash = $1
		    ), $2, $3) 
		returning
			amount`,
		groupHash,
		productId,
		amount).Scan(
		&amount); err != nil {
		return err
	}

	return nil
}

func RemoveGroupCartProduct(groupHash string, productId int) error {
	db, err := CreateConnection()
	if err == nil {
		defer db.Close()
	}
	if err != nil {
		return err
	}

	var amount float32
	if err := db.QueryRow(`
		delete from
			group_cart
		where
		    product_id = $2 and
		    group_cart.group_id = (
		        select
		            group_id
		        from
		            groups
		        where
		            group_hash = $1
		    )
		returning
			amount`,
		groupHash,
		productId).Scan(&amount); err != nil {
		return err
	}

	return nil
}

func ChangeGroupCartProduct(groupHash string, productId int, amount float32) error {
	db, err := CreateConnection()
	if err == nil {
		defer db.Close()
	}
	if err != nil {
		return err
	}

	if err := db.QueryRow(`
		update
			group_cart
		set
		    amount = $3
		where
		    product_id = $2 and
		    group_cart.group_id = (
		        select
		            group_id
		        from
		            groups
		        where
		            group_hash = $1
		    )
		returning
			amount`,
		groupHash,
		productId,
		amount).Scan(
		&amount); err != nil {
		return err
	}

	return nil
}
