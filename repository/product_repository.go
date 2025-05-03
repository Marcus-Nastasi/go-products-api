package repository

import (
	"github.com/Marcus-Nastasi/go-products-api/model"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) GetProducts() ([]model.Product, error) {
	var products []model.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) GetProduct(id string) (model.Product, error) {
	var product model.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return model.Product{}, err
	}
	return product, nil
}

func (r *ProductRepository) Create(product model.Product) (model.Product, error) {
	if err := r.db.Create(&product).Error; err != nil {
		return model.Product{}, err
	}
	return product, nil
}

// func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
// 	rows, err := pr.connection.Query("SELECT p.* FROM products p;")
// 	if err != nil {
// 		fmt.Println("Error getting data " + err.Error())
// 		return []model.Product{}, nil
// 	}
// 	var productsList []model.Product
// 	var product model.Product
// 	for rows.Next() {
// 		err = rows.Scan(
// 			&product.ID,
// 			&product.Name,
// 			&product.Price,
// 		)
// 		if err != nil {
// 			fmt.Println(err)
// 			return []model.Product{}, nil
// 		}
// 		productsList = append(productsList, product)
// 	}
// 	rows.Close()
// 	return productsList, nil
// }

// func (pr *ProductRepository) Create(new_product model.Product) (int, error) {
// 	var id int
// 	rows, err := pr.connection.Prepare("INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id;")
// 	if err != nil {
// 		return 0, err
// 	}
// 	err = rows.QueryRow(new_product.Name, new_product.Price).Scan(&id)
// 	if err != nil {
// 		return 0, err
// 	}
// 	rows.Close()
// 	return id, nil
// }
