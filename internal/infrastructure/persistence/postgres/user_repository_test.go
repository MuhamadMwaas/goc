package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"path/filepath" // New import
	"runtime"       // New import
	"testing"

	"github.com/futek/donation-campaign/internal/domain"
	"github.com/futek/donation-campaign/internal/infrastructure/migration"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	db         *sql.DB
	container  testcontainers.Container
	repository *UserRepository
}

func (suite *UserRepositoryTestSuite) SetupSuite() {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "postgres:14-alpine",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "test",
			"POSTGRES_PASSWORD": "test",
			"POSTGRES_DB":       "test",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatalf("Could not start container: %v", err)
	}

	host, err := container.Host(ctx)
	if err != nil {
		log.Fatalf("Could not get container host: %v", err)
	}

	port, err := container.MappedPort(ctx, "5432")
	if err != nil {
		log.Fatalf("Could not get container port: %v", err)
	}

	dsn := fmt.Sprintf("postgres://test:test@%s:%s/test?sslmode=disable", host, port.Port())
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	suite.db = db
	suite.container = container
	suite.repository = NewUserRepository(db)

	// Dynamically get the project root and then the migration path
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)                           // This is internal/infrastructure/persistence/postgres
	projectRoot := filepath.Join(basePath, "../../../..") // Navigate up to project root
	migrationsPath := filepath.Join(projectRoot, "migration")
	migrationsURL := "file://" + migrationsPath
	log.Printf("Migrations URL: %s\n", migrationsURL) // Debug print

	// Run migrations using the migration package
	if err := migration.Run(dsn, migrationsURL, "up"); err != nil {
		log.Fatalf("Could not run migrations: %v", err)
	}
}

func (suite *UserRepositoryTestSuite) TearDownSuite() {
	if err := suite.container.Terminate(context.Background()); err != nil {
		log.Fatalf("Could not terminate container: %v", err)
	}
}

func (suite *UserRepositoryTestSuite) TearDownTest() {
	_, err := suite.db.Exec("DELETE FROM users")
	if err != nil {
		log.Fatalf("Could not clean users table: %v", err)
	}
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}

func (suite *UserRepositoryTestSuite) TestCreateUser() {
	user := &domain.User{
		Email: "test@example.com",
		Name:  "Test User",
		Phone: "1234567890",
	}

	err := suite.repository.Create(context.Background(), user)
	assert.NoError(suite.T(), err)
	assert.NotZero(suite.T(), user.ID)
	assert.NotZero(suite.T(), user.CreatedAt)
	assert.NotZero(suite.T(), user.UpdatedAt)
}

func (suite *UserRepositoryTestSuite) TestGetUserByID() {
	user := &domain.User{
		Email: "test@example.com",
		Name:  "Test User",
		Phone: "1234567890",
	}
	err := suite.repository.Create(context.Background(), user)
	assert.NoError(suite.T(), err)

	foundUser, err := suite.repository.GetByID(context.Background(), user.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), user.ID, foundUser.ID)
	assert.Equal(suite.T(), user.Email, foundUser.Email)
	assert.Equal(suite.T(), user.Name, foundUser.Name)
	assert.Equal(suite.T(), user.Phone, foundUser.Phone)
}

func (suite *UserRepositoryTestSuite) TestUpdateUser() {
	user := &domain.User{
		Email: "test@example.com",
		Name:  "Test User",
		Phone: "1234567890",
	}
	err := suite.repository.Create(context.Background(), user)
	assert.NoError(suite.T(), err)

	user.Name = "Updated Test User"
	err = suite.repository.Update(context.Background(), user)
	assert.NoError(suite.T(), err)

	foundUser, err := suite.repository.GetByID(context.Background(), user.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Updated Test User", foundUser.Name)
}

func (suite *UserRepositoryTestSuite) TestDeleteUser() {
	user := &domain.User{
		Email: "test@example.com",
		Name:  "Test User",
		Phone: "1234567890",
	}
	err := suite.repository.Create(context.Background(), user)
	assert.NoError(suite.T(), err)

	err = suite.repository.Delete(context.Background(), user.ID)
	assert.NoError(suite.T(), err)

	_, err = suite.repository.GetByID(context.Background(), user.ID)
	assert.Error(suite.T(), err)
}
