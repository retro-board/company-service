package company

import (
	"context"
	"fmt"

	bugLog "github.com/bugfixes/go-bugfixes/logs"
	"github.com/mrz1836/go-sanitize"
	"github.com/retro-board/company-service/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Config *config.Config
	CTX    context.Context
}

type DataSet struct {
	CompanyID string `json:"company_id" bson:"company_id"`
	Generated int64  `json:"generated" bson:"generated"`
	Details   struct {
		Name   string `json:"name" bson:"name"`
		Domain string `json:"domain" bson:"domain"`
	} `json:"details" bson:"details"`
}

func NewMongo(c *config.Config) *Mongo {
	return &Mongo{
		Config: c,
		CTX:    context.Background(),
	}
}

func (m *Mongo) Get(domain string) (*DataSet, error) {
	client, err := mongo.Connect(
		m.CTX,
		options.Client().ApplyURI(fmt.Sprintf(
			"mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority",
			m.Config.Mongo.Username,
			m.Config.Mongo.Password,
			m.Config.Mongo.Host)),
	)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := client.Disconnect(m.CTX); err != nil {
			bugLog.Info(err)
		}
	}()

	var dataSet DataSet
	err = client.Database("company").Collection("company").FindOne(m.CTX, map[string]string{"domain": sanitize.AlphaNumeric(domain, false)}).Decode(&dataSet)
	if err != nil {
		return nil, err
	}

	return &dataSet, nil
}

func (m *Mongo) Create(data DataSet) error {
	return nil
}
