package aggregate

import (
	"time"

	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/domain/exception"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/retailer/domain/valueobject"
	sharedValueobject "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/shared/domain/valueobject"
)

type Retailer struct {
	id             sharedValueobject.UUID
	documentNumber valueobject.CNPJ
	name           valueobject.RetailerName
	tradeName      valueobject.RetailerTradeName
	createdAt      time.Time
	updatedAt      time.Time
	deletedAt      *time.Time
}

// NewRetailer Cria um novo Retailer, garantindo que os campos obrigatórios sejam validados e que o ID seja fornecido.
func NewRetailer(id string, documentNumber string, name string, tradeName string) (*Retailer, error) {
	retailerID, err := sharedValueobject.NewUUID(id)
	if err != nil {
		return nil, err
	}
	cnpj, err := valueobject.NewCNPJ(documentNumber)
	if err != nil {
		return nil, err
	}
	retailerName, err := valueobject.NewRetailerName(name)
	if err != nil {
		return nil, err
	}
	retailerTradeName, err := valueobject.NewRetailerTradeName(tradeName)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	return &Retailer{
		id:             retailerID,
		documentNumber: cnpj,
		name:           retailerName,
		tradeName:      retailerTradeName,
		createdAt:      now,
		updatedAt:      now,
	}, nil
}

// ReconstituteRetailer reconstrói um Retailer a partir de dados persistidos, garantindo que os campos sejam validados. O ID e o DocumentNumber são obrigatórios para reconstituição, e os timestamps são necessários para manter a integridade do estado do agregado.
func ReconstituteRetailer(id string, documentNumber string, name string, tradeName string, createdAt time.Time, updatedAt time.Time, deletedAt *time.Time) (*Retailer, error) {
	retailerID, err := sharedValueobject.NewUUID(id)
	if err != nil {
		return nil, err
	}
	cnpj, err := valueobject.NewCNPJ(documentNumber)
	if err != nil {
		return nil, err
	}
	retailerName, err := valueobject.NewRetailerName(name)
	if err != nil {
		return nil, err
	}
	retailerTradeName, err := valueobject.NewRetailerTradeName(tradeName)
	if err != nil {
		return nil, err
	}

	return &Retailer{
		id:             retailerID,
		documentNumber: cnpj,
		name:           retailerName,
		tradeName:      retailerTradeName,
		createdAt:      createdAt,
		updatedAt:      updatedAt,
		deletedAt:      deletedAt,
	}, nil
}

// Update atualiza os campos mutáveis do Retailer, garantindo que o agregado não esteja marcado como deletado. O DocumentNumber não pode ser atualizado, pois é um identificador único e imutável do Retailer.
func (r *Retailer) Update(name string, tradeName string) error {
	retailerName, err := valueobject.NewRetailerName(name)
	if err != nil {
		return err
	}
	retailerTradeName, err := valueobject.NewRetailerTradeName(tradeName)
	if err != nil {
		return err
	}
	if r.deletedAt != nil {
		return domainException.ErrRetailerDeleted
	}

	r.name = retailerName
	r.tradeName = retailerTradeName
	r.updatedAt = time.Now()
	return nil
}

func (r *Retailer) SoftDelete() error {
	if r.deletedAt != nil {
		return domainException.ErrRetailerDeleted
	}

	now := time.Now()
	r.deletedAt = &now
	r.updatedAt = now
	return nil
}

func (r *Retailer) ID() string             { return r.id.String() }
func (r *Retailer) DocumentNumber() string { return r.documentNumber.String() }
func (r *Retailer) Name() string           { return r.name.String() }
func (r *Retailer) TradeName() string      { return r.tradeName.String() }
func (r *Retailer) CreatedAt() time.Time   { return r.createdAt }
func (r *Retailer) UpdatedAt() time.Time   { return r.updatedAt }
func (r *Retailer) DeletedAt() *time.Time  { return r.deletedAt }
