package stores

import "zakroma_backend/schemas"

func GetGroupStoreList(groupHash string) ([]schemas.DishProduct, error) {
	db, err := CreateConnection()
	if err == nil {
		defer db.Close()
	}
	if err != nil {
		return []schemas.DishProduct{}, err
	}

	productsRows, err := db.Query(`
		select
		    group_store.product_id,
		    products.product_name,
		    group_store.amount,
		    products.unit_of_measurement
		from
		    group_store,
		    groups,
		    products
		where
		    group_store.group_id = groups.group_id and
		    groups.group_hash = $1 and
		    group_store.product_id = products.product_id
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

func AddGroupStoreProduct(groupHash string, productId int, amount float32) error {
	db, err := CreateConnection()
	if err == nil {
		defer db.Close()
	}
	if err != nil {
		return err
	}

	if err := db.QueryRow(`
		insert into
			group_store(group_id, product_id, amount)
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

func RemoveGroupStoreProduct(groupHash string, productId int) error {
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
			group_store
		where
		    product_id = $2 and
		    group_store.group_id = (
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

func ChangeGroupStoreProduct(groupHash string, productId int, amount float32) error {
	db, err := CreateConnection()
	if err == nil {
		defer db.Close()
	}
	if err != nil {
		return err
	}

	if err := db.QueryRow(`
		update
			group_store
		set
		    amount = $3
		where
		    product_id = $2 and
		    group_store.group_id = (
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
