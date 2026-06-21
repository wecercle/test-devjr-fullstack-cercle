package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/aggregate"
	domainException "github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/exception"
	"github.com/wecercle/test-devjr-fullstack-cercle/core/modules/resale/domain/valueobject"
)

const (
	testCPF     = "10579550001"
	testOrderID = "6ae28fac-280e-4879-bbe3-9277decb4a06"
	testItemID  = "b14e75f8-a3d6-4f8d-8a2b-5fbb0cf51001"
)

type fakeResaleOrderQueryRepository struct {
	orderExists bool
	item        *aggregate.ResaleOrderItem
}

func (r *fakeResaleOrderQueryRepository) ExistsOrderByCPFAndOrderID(ctx context.Context, cpf, orderID string) (bool, error) {
	return r.orderExists, nil
}

func (r *fakeResaleOrderQueryRepository) SelectOrderItemsByCPFAndOrderID(ctx context.Context, cpf, orderID string) ([]*aggregate.ResaleOrderItem, error) {
	return nil, nil
}

func (r *fakeResaleOrderQueryRepository) SelectOrderItemByOrderIDAndItemID(ctx context.Context, orderID, itemID string) (*aggregate.ResaleOrderItem, error) {
	if r.item == nil {
		return nil, domainException.ErrOrderItemNotFound
	}
	return r.item, nil
}

type fakeResaleOrderCommandRepository struct {
	updated      bool
	updateCalled bool
}

func (r *fakeResaleOrderCommandRepository) MarkOrderItemAsReturned(ctx context.Context, orderID, itemID string) (bool, error) {
	r.updateCalled = true
	return r.updated, nil
}

func TestCancelOrderItemUseCaseExecute(t *testing.T) {
	now := time.Now().UTC()
	deliveredRecently := now.AddDate(0, 0, -2)
	deliveredTooLongAgo := now.AddDate(0, 0, -8)

	tests := []struct {
		name             string
		orderID          string
		itemID           string
		orderExists      bool
		item             *aggregate.ResaleOrderItem
		updateResult     bool
		wantErr          error
		wantUpdateCalled bool
	}{
		{
			name:    "invalid order_id returns ErrInvalidOrderID",
			orderID: "invalid",
			itemID:  testItemID,
			wantErr: domainException.ErrInvalidOrderID,
		},
		{
			name:    "invalid item_id returns ErrInvalidOrderItemID",
			orderID: testOrderID,
			itemID:  "invalid",
			wantErr: domainException.ErrInvalidOrderItemID,
		},
		{
			name:        "order does not belong to cpf returns ErrOrderNotFound",
			orderID:     testOrderID,
			itemID:      testItemID,
			orderExists: false,
			wantErr:     domainException.ErrOrderNotFound,
		},
		{
			name:        "item already returned is idempotent and does not call update",
			orderID:     testOrderID,
			itemID:      testItemID,
			orderExists: true,
			item:        newOrderItem(valueobject.ShippingStatusReturned, &deliveredRecently),
		},
		{
			name:        "posted item is not returnable",
			orderID:     testOrderID,
			itemID:      testItemID,
			orderExists: true,
			item:        newOrderItem(valueobject.ShippingStatusPosted, nil),
			wantErr:     domainException.ErrOrderItemNotReturnable,
		},
		{
			name:        "label generated item is not returnable",
			orderID:     testOrderID,
			itemID:      testItemID,
			orderExists: true,
			item:        newOrderItem(valueobject.ShippingStatusLabelGenerated, nil),
			wantErr:     domainException.ErrOrderItemNotReturnable,
		},
		{
			name:        "cancelled item is not returnable",
			orderID:     testOrderID,
			itemID:      testItemID,
			orderExists: true,
			item:        newOrderItem(valueobject.ShippingStatusCancelled, nil),
			wantErr:     domainException.ErrOrderItemNotReturnable,
		},
		{
			name:        "delivered item with nil delivered_at is not returnable",
			orderID:     testOrderID,
			itemID:      testItemID,
			orderExists: true,
			item:        newOrderItem(valueobject.ShippingStatusDelivered, nil),
			wantErr:     domainException.ErrOrderItemNotReturnable,
		},
		{
			name:        "delivered item older than seven days returns ErrOrderItemReturnExpired",
			orderID:     testOrderID,
			itemID:      testItemID,
			orderExists: true,
			item:        newOrderItem(valueobject.ShippingStatusDelivered, &deliveredTooLongAgo),
			wantErr:     domainException.ErrOrderItemReturnExpired,
		},
		{
			name:             "delivered item within seven days updates and succeeds",
			orderID:          testOrderID,
			itemID:           testItemID,
			orderExists:      true,
			item:             newOrderItem(valueobject.ShippingStatusDelivered, &deliveredRecently),
			updateResult:     true,
			wantUpdateCalled: true,
		},
		{
			name:             "update returns false returns ErrOrderItemNotReturnable",
			orderID:          testOrderID,
			itemID:           testItemID,
			orderExists:      true,
			item:             newOrderItem(valueobject.ShippingStatusDelivered, &deliveredRecently),
			updateResult:     false,
			wantErr:          domainException.ErrOrderItemNotReturnable,
			wantUpdateCalled: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryRepo := &fakeResaleOrderQueryRepository{
				orderExists: tt.orderExists,
				item:        tt.item,
			}
			commandRepo := &fakeResaleOrderCommandRepository{updated: tt.updateResult}
			useCase := NewCancelOrderItemUseCase(queryRepo, commandRepo)

			err := useCase.Execute(context.Background(), testCPF, tt.orderID, tt.itemID)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected error %v, got %v", tt.wantErr, err)
			}
			if commandRepo.updateCalled != tt.wantUpdateCalled {
				t.Fatalf("expected updateCalled %v, got %v", tt.wantUpdateCalled, commandRepo.updateCalled)
			}
		})
	}
}

func newOrderItem(status valueobject.ShippingStatus, deliveredAt *time.Time) *aggregate.ResaleOrderItem {
	return aggregate.NewResaleOrderItem(
		testItemID,
		testOrderID,
		"LAB-CAMISETA-001",
		"Camiseta Dry Fit Feminina",
		1,
		"219.00",
		nil,
		status,
		deliveredAt,
	)
}
