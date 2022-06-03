package internal

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/felixbecker/hexadiscountexample/application"
	"github.com/felixbecker/hexadiscountexample/discounter"
	"github.com/felixbecker/hexadiscountexample/store"
	"github.com/felixbecker/hexadiscountexample/storeprovider"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

var (
	ErrNoRedisAdr               string = "Please provide configuration settigns for the redis [addr = address:port]"
	ErrNoPostgresConfigSettings string = "Error: Please provide configuration settigns for the postgres db [user, password, host, db, table name ]"
)

type Redis struct {
	Addr string `env:"REDIS_ADDR,default=localhost:6379"`
}

func (r *Redis) Validate() error {

	if r.Addr == "" {
		return fmt.Errorf(ErrNoRedisAdr)
	}
	return nil
}

type Postgres struct {
	User      string `env:"POSTGRES_USER"`
	Password  string `env:"POSTGRES_PASSWORD"`
	Host      string `env:"POSTGRES_HOST"`
	DB        string `env:"POSTGRES_DB,default=discounter"`
	Tablename string `env:"POSTGRES_TABLE,default=rates"`
}

func (p *Postgres) Validate() error {

	if p.User != "" && p.Password != "" && p.Host != "" && p.DB != "" && p.Tablename != "" {
		return nil
	}
	return fmt.Errorf(ErrNoPostgresConfigSettings)
}

type Config struct {
	ENV       string `env:"ENV,default=development"`
	StoreType string `env:"STORE_TYPE,default=inmemory"`
	Redis     Redis
	Postgres  Postgres
}

type Factory struct {
	config *Config

	store       *store.Store
	discounter  *discounter.Discounter
	application *application.Application
}

func NewFactory() (*Factory, error) {
	var config Config
	const EnvDevelopment = "development"

	env := os.Getenv("ENV")
	if env == "" {
		env = EnvDevelopment
	}

	env = strings.ToLower(env)
	config.ENV = env

	log.Printf("Figured out your are running in %s env", config.ENV)
	_ = godotenv.Load(".env." + env + ".local")

	if env != "test" {
		_ = godotenv.Load(".env.local")
	}
	_ = godotenv.Load(".env." + env)
	_ = godotenv.Load(".env")
	_ = godotenv.Load()

	err := envconfig.Process(context.Background(), &config)
	if err != nil {
		return nil, err
	}

	fmt.Println(config)
	return &Factory{
		config: &config,
	}, nil
}

func (f *Factory) Store() *store.Store {

	if f.store == nil {

		var p store.Provider
		switch f.config.StoreType {
		case "inmemory":
			p = storeprovider.NewInMemory()
			log.Println("Assigned InMemory Store")
		case "postgres":
			err := f.config.Postgres.Validate()
			if err != nil {
				panic(err)
			}
			p, err = storeprovider.NewPostgresProvider(
				f.config.Postgres.User,
				f.config.Postgres.Password,
				f.config.Postgres.Host,
				f.config.Postgres.DB, f.config.Postgres.Tablename)
			if err != nil {
				panic(err)
			}
			log.Println("Assigned Postgres Store")
		case "redis":
			err := f.config.Redis.Validate()
			if err != nil {
				panic(err)
			}
			p = storeprovider.NewRedisProvider(f.config.Redis.Addr)
			log.Println("Assigned Redis Store")

		default:
			log.Println("Not a valid store type")
			panic("YIKES")
		}

		f.store = store.New(p)

	}

	return f.store
}

func (f *Factory) Discounter() *discounter.Discounter {

	if f.discounter == nil {
		f.discounter = discounter.New(f.Store())
	}

	return f.discounter
}

func (f *Factory) Application() *application.Application {
	if f.application == nil {
		f.application = application.New(f.Discounter())
	}
	return f.application
}
