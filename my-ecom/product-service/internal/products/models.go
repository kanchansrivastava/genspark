package products

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Product struct {
	ID          string    `json:"id"` // UUID
	Name        string    `db:"name" json:"name" binding:"required"`
	Description string    `db:"description" json:"description"`
	Price       string    `db:"price" json:"price" binding:"required"`
	Category    string    `db:"category" json:"category"`
	Stock       int       `db:"stock" json:"stock" binding:"required"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"` // Creation timestamp
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"` // Last updated timestamp
}

type NewProduct struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"omitempty"`
	Price       string `json:"price" binding:"required"`
	Category    string `json:"category" binding:"omitempty"`
	Stock       int    `json:"stock" binding:"required,gte=0"`
}

type StripeProductPricing struct {
	ID              int       `json:"id"`                // Unique identifier
	ProductID       string    `json:"product_id"`        // Foreign key referencing products table
	StripeProductID string    `json:"stripe_product_id"` // Stripe product ID
	PriceID         string    `json:"price_id"`          // Stripe price ID
	Price           uint64    `json:"price"`
	CreatedAt       time.Time `json:"created_at"` // Timestamp when the record was created
	UpdatedAt       time.Time `json:"updated_at"` // Timestamp when the record was last updated
}

func (p *NewProduct) ValidatePrice() error {
	price := strings.TrimSpace(p.Price)
	parts := strings.Split(price, ".")
	if len(parts) > 2 {
		return fmt.Errorf("invalid price format: too many parts")
	}

	rsPart := parts[0]
	if len(rsPart) == 0 || !isUint(rsPart) {
		return fmt.Errorf("invalid rupee part: must be a valid positive integer")
	}

	rupee, _ := strconv.Atoi(rsPart)
	if rupee < 100 || rupee > 1000000 {
		return fmt.Errorf("rupee part must be between 100 and 1,000,000")
	}

	paisaPart := "00"
	if len(parts) == 2 {
		paisaPart = parts[1]
		if !isUint(paisaPart) || len(paisaPart) > 2 {
			return fmt.Errorf("invalid paisa part: must be a valid number with at most two digits")
		}
		if len(paisaPart) == 1 {
			paisaPart = "0" + paisaPart
		}
	}

	// Calculate the final price in paisa (rupee * 100 + paisa)
	paisa, _ := strconv.Atoi(paisaPart)
	finalPrice := rupee*100 + paisa

	p.Price = strconv.Itoa(finalPrice)

	return nil
}

func isUint(value string) bool {
	_, err := strconv.Atoi(value)
	return err == nil
}
