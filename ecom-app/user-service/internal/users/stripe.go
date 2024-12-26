package users

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/customer"
	"log/slog"
	"os"
	"time"
	"user-service/pkg/logkey"
)

func (c *Conf) CreateCustomerStripe(ctx context.Context, userId, name, email string) error {
	sKey := os.Getenv("STRIPE_TEST_KEY")
	if sKey == "" {
		return fmt.Errorf("STRIPE_TEST_KEY not set")
	}

	stripe.Key = sKey

	err := c.withTx(ctx, func(tx *sql.Tx) error {
		// we will check if customer is present in the users_stripe table
		sqlQuery := `
				SELECT stripe_customer_id 
				FROM users_stripe
				WHERE user_id = $1
				`

		var stripeCustomerId string
		err := tx.QueryRowContext(ctx, sqlQuery, userId).Scan(&stripeCustomerId)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("failed to fetch Stripe customer ID: %w", err)
			}
			params := &stripe.CustomerParams{
				Name:  stripe.String(name),
				Email: stripe.String(email),
			}

			customerResult, err := customer.New(params)
			if err != nil {
				slog.Error("failed to create Stripe customer", slog.Any(logkey.ERROR, err))
				return fmt.Errorf("failed to create Stripe customer: %w", err)
			}

			// Define the SQL query
			query := `
		INSERT INTO users_stripe (user_id,email, stripe_customer_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`
			createdAt := time.Now().UTC()
			updatedAt := createdAt
			res, err := tx.ExecContext(ctx, query, userId, email, customerResult.ID, createdAt, updatedAt)
			if err != nil {
				slog.Error("failed to insert Stripe customer ID", slog.Any(logkey.ERROR, err))
				return fmt.Errorf("failed to insert Stripe customer ID: %w", err)
			}
			if num, err := res.RowsAffected(); num == 0 || err != nil {
				return fmt.Errorf("failed to insert Stripe customer ID: %w", err)
			}
			return nil

		}

		// success case
		return nil
	})

	if err != nil {
		return err
	}
	return nil

}
