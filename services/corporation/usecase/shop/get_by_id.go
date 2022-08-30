package shop

import (
	"avatar/services/corporation/config"
	"avatar/services/corporation/domain/entity"
	"avatar/services/corporation/domain/repository"
	pb "avatar/services/corporation/protos"
	"context"
	"time"

	"github.com/jinzhu/copier"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

func GetById(
	ctx context.Context,
	db neo4j.Driver,
	shopRepository repository.ShopRepository,
	input *pb.GetShopRequest,
) (*pb.CreateShopResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	shopRaw, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		shop, err := shopRepository.GetById(ctx, input.Id)

		if err != nil {
			log.Error("error when write transaction, error: ", err)
			return nil, err
		}
		return shop, nil
	})
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	shop := shopRaw.(*entity.Shop)
	var result pb.CreateShopResponse
	err = copier.Copy(&result, shop)
	if err != nil {
		return nil, err
	}
	result.CreatedAt = shop.CreatedAt.Format(time.RFC3339)
	if shop.UpdatedAt.IsZero() {
		result.UpdatedAt = ""
	} else {
		result.UpdatedAt = shop.UpdatedAt.Format(time.RFC3339)
	}
	return &result, nil
}
