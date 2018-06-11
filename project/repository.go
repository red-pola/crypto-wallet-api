package project

import (
	"github.com/red-pola/crypto-wallet-api/entity"
)

// Repository project repository inteface
type Repository interface {
	Find(id entity.ID) (*entity.Project, error)
	FindByName(name string) ([]*entity.Project, error)
	FindAll() ([]*entity.Project, error)
	Save(project *entity.Project) (entity.ID, error)
	Update(project *entity.Project) error
	Delete(id entity.ID) error
}
