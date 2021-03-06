package domain

import (
	"time"

	"github.com/hashicorp/go-uuid"
)

type Entity struct {
	id          string
	createdAt   time.Time
	modifiedAt  time.Time
	createdBy   string
	modifiedBy  string
	deletedDate time.Time
	deletedBy   string
}

func New() *Entity {
	return &Entity{createdAt: time.Now(), id: GenerateId()}
}

func NewEntityWithId(id string) *Entity {
	return &Entity{id: id, createdAt: time.Now()}
}

func NewEntity(id string,
	createdAt, modifiedAt time.Time,
	createdBy, modifiedBy string,
	deletedDate time.Time,
	deletedBy string,
) *Entity {
	return &Entity{id, createdAt, modifiedAt, createdBy, modifiedBy, deletedDate, deletedBy}
}

func GenerateId() string {
	id, err := uuid.GenerateUUID()
	if err != nil {
		return ""
	}
	return id
}

func (e *Entity) IsNew() bool {
	return e.id == ""
}

func (e *Entity) Id() string {
	return e.id
}

func (e *Entity) CreatedAt() time.Time {
	return e.createdAt
}

func (e *Entity) ModifiedAt() time.Time {
	return e.modifiedAt
}

func (e *Entity) CreatedBy() string {
	return e.createdBy
}

func (e *Entity) ModifiedBy() string {
	return e.modifiedBy
}

func (e *Entity) DeletedDate() time.Time {
	return e.deletedDate
}

func (e *Entity) DeletedBy() string {
	return e.deletedBy
}

func (e *Entity) IsDeleted() bool {
	return e.deletedDate.IsZero()
}

func (e *Entity) Delete(userId string) {
	e.deletedDate = time.Now()
	e.deletedBy = userId
}

func (e *Entity) Restore() {
	e.deletedDate = time.Time{}
	e.deletedBy = ""
}
