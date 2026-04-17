package example_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

// ============================================================
// ARRANGE / ACT / ASSERT — one clear outcome per test
// ============================================================

func TestCreateUser_ReturnsCreatedUser(t *testing.T) {
	// Arrange
	db := setupDB(t)
	repo := NewUserRepository(db)
	input := CreateUserInput{Name: "Alice", Email: "alice@example.com"}

	// Act
	user, err := repo.CreateUser(context.Background(), input)

	// Assert
	require.NoError(t, err)
	assert.Equal(t, "Alice", user.Name)
	assert.Equal(t, "alice@example.com", user.Email)
	assert.NotEmpty(t, user.ID)
}

// ============================================================
// TEST ERROR PATHS — not just the happy path
// ============================================================

func TestCreateUser_DuplicateEmail_ReturnsError(t *testing.T) {
	// Arrange
	db := setupDB(t)
	repo := NewUserRepository(db)
	input := CreateUserInput{Name: "Alice", Email: "alice@example.com"}
	_, err := repo.CreateUser(context.Background(), input)
	require.NoError(t, err)

	// Act — create again with same email
	_, err = repo.CreateUser(context.Background(), input)

	// Assert
	require.Error(t, err)
	assert.ErrorIs(t, err, ErrDuplicateEmail)
}

// ============================================================
// TESTCONTAINERS — real Postgres, not SQLite or mocks
// No shared state between tests — each gets a fresh DB
// ============================================================

func setupDB(t *testing.T) *sql.DB {
	t.Helper()

	ctx := context.Background()

	container, err := postgres.Run(ctx,
		"postgres:16",
		postgres.WithDatabase("testdb"),
		postgres.WithUsername("user"),
		postgres.WithPassword("password"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections"),
		),
	)
	require.NoError(t, err)

	t.Cleanup(func() {
		_ = container.Terminate(ctx)
	})

	connStr, err := container.ConnectionString(ctx, "sslmode=disable")
	require.NoError(t, err)

	db, err := sql.Open("postgres", connStr)
	require.NoError(t, err)

	// Run migrations
	err = runMigrations(db)
	require.NoError(t, err)

	return db
}
