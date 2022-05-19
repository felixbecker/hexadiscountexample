package internal

import (
	"fmt"
	"log"

	"github.com/felixbecker/hexadiscountexample/application"
	"github.com/felixbecker/hexadiscountexample/discounter"
	"github.com/felixbecker/hexadiscountexample/store"
	"github.com/felixbecker/hexadiscountexample/storeprovider"
)

type Redis struct {
	Addr string
}

func (r *Redis) Validate() error {

	if r.Addr == "" {
		return fmt.Errorf("Please provide configuration settigns for the redis [addr = address:port]")
	}
	return nil
}

type Postgres struct {
	User      string
	Password  string
	Host      string
	DB        string
	Tablename string
}

func (p *Postgres) Validate() error {

	if p.User != "" && p.Password != "" && p.Host != "" && p.DB != "" && p.Tablename != "" {
		return nil
	}
	return fmt.Errorf("Error: Please provide configuration settigns for the postgres db [user, password, host, db, table name ]")
}

type Config struct {
	StoreType string
	Redis     Redis
	Postgres  Postgres
}

type Factory struct {
	config *Config

	store       *store.Store
	discounter  *discounter.Discounter
	application *application.Application
}

func NewFactory(config *Config) *Factory {

	fmt.Println(config)
	return &Factory{
		config: config,
	}
}

func (f *Factory) Store() *store.Store {

	if f.store == nil {

		var p store.Provider
		switch f.config.StoreType {
		case "inMemory":
			p = storeprovider.NewInMemory()
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
		case "redis":
			err := f.config.Redis.Validate()
			if err != nil {
				panic(err)
			}
			p = storeprovider.NewRedisProvider(f.config.Redis.Addr)

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
