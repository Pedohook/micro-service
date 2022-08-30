package service_template_category

import (
	"avatar/services/center/config"
	"avatar/services/center/domain/repository"
	"avatar/services/center/domain/entity"
	pb "avatar/services/center/protos"
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/copier"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func Update(
	ctx context.Context,
	db neo4j.Driver,
	serviceTemplateCategoryRepository repository.ServiceTemplateCategoryRepository,
	input *pb.UpdateServiceTemplateCategoryRequest,
) (*pb.CreateServiceTemplateCategoryResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	data := entity.ServiceTemplateCategory{}
	err := copier.Copy(&data, input)
	if err != nil {
		return nil, err
	}
	ServiceTemplateCategoryRaw, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		ServiceTemplateCategory, err := serviceTemplateCategoryRepository.Update(ctx, data.Id, &data)
		if err != nil {
			log.Error("error when write transaction, error: ", err)
			return nil, err
		}
		return *ServiceTemplateCategory, nil
	})
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	ServiceTemplateCategory := ServiceTemplateCategoryRaw.(entity.ServiceTemplateCategory)
	var result pb.CreateServiceTemplateCategoryResponse
	err = copier.Copy(&result, ServiceTemplateCategory)
	if err != nil {
		return nil, err
	}
	result.CreatedAt = ServiceTemplateCategory.CreatedAt.Format(time.RFC3339)
	result.UpdatedAt = ServiceTemplateCategory.UpdatedAt.Format(time.RFC3339)
	return &result, nil
}
