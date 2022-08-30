package pos

import (
	"avatar/services/pos/config"
	"avatar/services/pos/domain/entity"
	"avatar/services/pos/domain/repository"
	pb "avatar/services/pos/protos"
	"context"
	"time"

	"github.com/jinzhu/copier"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

func Update(
	ctx context.Context,
	db neo4j.Driver,
	posRepository repository.PosRepository,
	input *pb.UpdatePosRequest,
) (*pb.CreatePosResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	data := entity.Pos{}
	err := copier.Copy(&data, input)
	if err != nil {
		return nil, err
	}
	posRaw, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		pos, err := posRepository.Update(ctx, data.Id, &data)
		if err != nil {
			log.Error("error when update pos, error: ", err)
			return nil, err
		}
		return *pos, nil
	})
	if err != nil {
		log.Error("error when update pos, error: ", err)
		return nil, err
	}
	pos := posRaw.(entity.Pos)
	var result pb.CreatePosResponse
	err = copier.Copy(&result, pos)
	if err != nil {
		return nil, err
	}
	result.CreatedAt = pos.CreatedAt.Format(time.RFC3339)
	result.UpdatedAt = pos.UpdatedAt.Format(time.RFC3339)
	return &result, nil
}
