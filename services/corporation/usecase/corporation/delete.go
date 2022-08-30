package corporation

import (
	"avatar/services/corporation/config"
	"avatar/services/corporation/domain/repository"
	pb "avatar/services/corporation/protos"
	"context"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

func Delete(
	ctx context.Context,
	db neo4j.Driver,
	corporationRepository repository.CorporationRepository,
	input *pb.DeleteCorporationRequest,
) (*pb.DeleteCorporationResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	rowsAffectedRaw, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		rowsAffected, err := corporationRepository.Delete(ctx, input.Id)
		if err != nil {
			log.Error("error when write transaction, error: ", err)
			return nil, err
		}
		return rowsAffected, nil
	})
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	rowsAffected := rowsAffectedRaw.(int64)
	return &pb.DeleteCorporationResponse{
		RowsAffected: rowsAffected,
	}, nil
}
