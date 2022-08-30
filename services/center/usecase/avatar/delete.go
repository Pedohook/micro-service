package avatar

import (
	"avatar/services/center/config"
	"avatar/services/center/domain/repository"
	pb "avatar/services/center/protos"
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func Delete(
	ctx context.Context,
	db neo4j.Driver,
	avatarRepository repository.AvatarRepository,
	input *pb.DeleteByIdRequest,
) (*pb.DeleteResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	rowsAffectedRaw, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		rowsAffected, err := avatarRepository.Delete(ctx, input.Id)
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
	return &pb.DeleteResponse{
		RowsAffected: rowsAffected,
	}, nil
}
