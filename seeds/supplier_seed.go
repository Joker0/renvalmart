package seeds

import (
	"github.com/joker0/renvalmart/internal/app/models"
	"gorm.io/gorm"
)

func SeedSuppliers(db *gorm.DB) {
	suppliers := []models.Supplier{
		{Name: "Supplier 1", ContactPerson: "John Doe", Phone: "123-456-7890", Email: "supplier1@example.com", Address: "123 Main St"},
		{Name: "Supplier 2", ContactPerson: "Jane Smith", Phone: "987-654-3210", Email: "supplier2@example.com", Address: "456 Elm St"},
		{Name: "Supplier 3", ContactPerson: "Bob Johnson", Phone: "555-555-5555", Email: "supplier3@example.com", Address: "789 Oak Ave"},
		{Name: "Supplier 4", ContactPerson: "Alice Brown", Phone: "777-777-7777", Email: "supplier4@example.com", Address: "321 Pine Rd"},
		{Name: "Supplier 5", ContactPerson: "Eve White", Phone: "888-888-8888", Email: "supplier5@example.com", Address: "654 Birch Ln"},
		{Name: "Supplier 6", ContactPerson: "Sam Green", Phone: "333-333-3333", Email: "supplier6@example.com", Address: "234 Cedar Dr"},
		{Name: "Supplier 7", ContactPerson: "Lisa Black", Phone: "222-222-2222", Email: "supplier7@example.com", Address: "876 Red St"},
		{Name: "Supplier 8", ContactPerson: "Mike Gray", Phone: "666-666-6666", Email: "supplier8@example.com", Address: "432 Maple Ave"},
		{Name: "Supplier 9", ContactPerson: "Sara Brown", Phone: "111-111-1111", Email: "supplier9@example.com", Address: "987 Elm Rd"},
		{Name: "Supplier 10", ContactPerson: "Tom Wilson", Phone: "444-444-4444", Email: "supplier10@example.com", Address: "765 Oak Dr"},
	}

	for _, supplier := range suppliers {
		db.Create(&supplier)
	}
}
