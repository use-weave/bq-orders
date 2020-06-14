// Package p contains a Pub/Sub Cloud Function.
package p

import (
	"bytes"
	"context"
	"io"
	"os"

	"github.com/use-weave/discounts"
	"github.com/use-weave/events"
	limits "github.com/use-weave/legal-limits"
	"github.com/use-weave/order-api/pkg/elastic"
	"github.com/use-weave/order-api/pkg/mysql"
	"github.com/use-weave/order-api/pkg/neo4j"
	"github.com/use-weave/order-api/pkg/order"
	"github.com/use-weave/pricing"
)

type receiverInteractor struct {
	interactor events.ReceiverInteractor
}

func (r *receiverInteractor) Close() error {
	return nil
}

func (r *receiverInteractor) Receive(reader io.Reader, attr map[string]string) error {
	return r.interactor.Receive(reader, attr)
}

// PubSubMessage is the payload of a Pub/Sub event.
type PubSubMessage struct {
	Data       []byte            `json:"data"`
	Attributes map[string]string `json:"attributes"`
}

// HandleEvent consumes a Pub/Sub message.
func HandleEvent(ctx context.Context, m PubSubMessage) error {

	searchPathValue := os.Getenv("SEARCH_PATH")

	c := elastic.New(searchPathValue)
	if err := c.Open(); err != nil {
		panic(err)
	}

	defer c.Close()

	neo4jPathValue := "bolt://localhost:7687"

	neo4jPathEnv := os.Getenv("NEO4J_PATH")
	if neo4jPathEnv != "" {
		neo4jPathValue = neo4jPathEnv
	}

	mysqlPathValue := "root:root@/weave?parseTime=true"

	mysqlPathEnv := os.Getenv("MYSQL_PATH")
	if mysqlPathEnv != "" {
		mysqlPathValue = mysqlPathEnv
	}

	db := mysql.New(mysqlPathValue)
	if err := db.Open(); err != nil {
		panic(err)
	}
	defer db.Close()

	cli := &client{DefaultReceiver: events.NewReceiver(), indexer: c, db: neo4j.New(neo4jPathValue), order: order.New(
		db,
		pricing.New(db.PricingService()),
		discounts.New(),
		limits.New(),
	)}

	receiver := &receiverInteractor{interactor: events.NewReceiverInteractor(cli)}
	receiver.Receive(bytes.NewReader(m.Data), m.Attributes)

	return nil
}
