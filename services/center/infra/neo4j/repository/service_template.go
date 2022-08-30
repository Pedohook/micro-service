package repository

import (
	"avatar/pkg/base_repository"
	"avatar/services/center/config"
	"avatar/services/center/domain/entity"
	"avatar/services/center/domain/repository"
	"context"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

type serviceTemplateRepositoryImpl struct {
	baseRepository base_repository.BaseRepository
}

func NewServiceTemplateRepository(
	baseRepository base_repository.BaseRepository,
) repository.ServiceTemplateRepository {
	return &serviceTemplateRepositoryImpl{
		baseRepository: baseRepository,
	}
}

func (r *serviceTemplateRepositoryImpl) Create(ctx context.Context, input *entity.ServiceTemplate) (*entity.ServiceTemplate, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.Create(transaction, input)
	if err != nil {
		log.Error("error when create service template, error: ", err)
		return nil, err
	}
	return input, nil
}

func (r *serviceTemplateRepositoryImpl) GetList(ctx context.Context, input entity.GetListServiceTemplateOption) ([]*entity.ServiceTemplate, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	result := []*entity.ServiceTemplate{}
	query := r.baseRepository.Model(&entity.ServiceTemplate{})
	if input.PerPage > 0 {
		query = query.Limit(input.PerPage)
	}
	if input.Page > 0 {
		query = query.Skip((input.Page - 1) * input.PerPage)
	}
	if input.ServiceTemplateCategory > 0 {
		query = query.Where("serviceTemplateCategory = ?", input.ServiceTemplateCategory)
	}
	if input.CorporationId > 0 {
		query = query.Where("corporationId = ?", input.CorporationId)
	}
	err := query.Find(transaction, &result)
	if err != nil {
		log.Error("error when get list service template, error: ", err)
		return nil, err
	}
	return result, nil
}

func (r *serviceTemplateRepositoryImpl) GetById(ctx context.Context, id int64) (*entity.ServiceTemplate, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	var result entity.ServiceTemplate
	err := r.baseRepository.Model(&entity.ServiceTemplate{}).Where("id = ?", id).FindOne(transaction, &result)
	if err != nil {
		log.Error("error when get service template, error: ", err)
		return nil, err
	}
	return &result, nil
}

func (r *serviceTemplateRepositoryImpl) Update(ctx context.Context, id int64, input *entity.ServiceTemplate) (*entity.ServiceTemplate, error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	err := r.baseRepository.Model(&entity.ServiceTemplate{}).Where("id = ?", id).Select("*").Update(transaction, input)
	if err != nil {
		log.Error("error when update service template, error: ", err)
		return nil, err
	}
	return input, nil
}

func (r *serviceTemplateRepositoryImpl) Delete(ctx context.Context, id int64) (rowsAffected int64, err error) {
	transaction := ctx.Value(config.Neo4jTransactionContextKey).(neo4j.Transaction)
	rowsAffected, err = r.baseRepository.Model(&entity.ServiceTemplate{}).Where("id = ?", id).Delete(transaction)
	if err != nil {
		log.Error("error when delete service template, error: ", err)
		return 0, err
	}
	return rowsAffected, nil
}
