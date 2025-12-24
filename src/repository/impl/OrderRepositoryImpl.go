package impl

import (
	"context"
	"fmt"
	"strings"
	"time"

	"cloud.google.com/go/firestore"

	"github.com/google/uuid"
	"github.com/ruiborda/go-service-common/database"
	"github.com/ruiborda/go-service-common/dto"
	"github.com/ruiborda/pos-api/src/model"
	"github.com/ruiborda/pos-api/src/repository"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrderRepositoryImpl struct {
	collectionName string
}

func NewOrderRepositoryImpl() repository.OrderRepository {
	return &OrderRepositoryImpl{collectionName: "orders"}
}

func (r *OrderRepositoryImpl) Create(order *model.Order) (*model.Order, error) {
	ctx := context.Background()
	client := database.GetFirestoreClient()

	if order.Id == "" {
		order.Id = uuid.New().String()
	}

	_, err := client.Collection(r.collectionName).Doc(order.Id).Set(ctx, order)
	if err != nil {
		return nil, fmt.Errorf("failed to create order: %v", err)
	}
	return order, nil
}

func (r *OrderRepositoryImpl) FindById(id string) (*model.Order, error) {
	ctx := context.Background()
	client := database.GetFirestoreClient()
	docSnap, err := client.Collection(r.collectionName).Doc(id).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get order: %v", err)
	}

	var order model.Order
	if err := docSnap.DataTo(&order); err != nil {
		return nil, fmt.Errorf("failed to convert document to order: %v", err)
	}
	order.Id = docSnap.Ref.ID
	return &order, nil
}

func (r *OrderRepositoryImpl) Update(order *model.Order) (*model.Order, error) {
	ctx := context.Background()
	client := database.GetFirestoreClient()
	_, err := client.Collection(r.collectionName).Doc(order.Id).Set(ctx, order)
	if err != nil {
		return nil, fmt.Errorf("failed to update order: %v", err)
	}
	return order, nil
}

func (r *OrderRepositoryImpl) Delete(id string) error {
	ctx := context.Background()
	client := database.GetFirestoreClient()
	_, err := client.Collection(r.collectionName).Doc(id).Delete(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete order: %v", err)
	}
	return nil
}

func (r *OrderRepositoryImpl) applyFilters(q firestore.Query, startDate, endDate *time.Time, filters []*dto.PageRequestFilter) firestore.Query {
	if startDate != nil {
		q = q.Where("orderDate", ">=", *startDate)
	}
	if endDate != nil {
		q = q.Where("orderDate", "<=", *endDate)
	}

	for _, f := range filters {
		q = q.Where(f.Field, string(f.Operator), f.Value.Export())
	}
	return q
}

func (r *OrderRepositoryImpl) FindAllByPageAndSize(page, size int, sorts []*dto.PageRequestOrder, startDate, endDate *time.Time, filters []*dto.PageRequestFilter) ([]*model.Order, error) {
	ctx := context.Background()
	client := database.GetFirestoreClient()
	offset := page * size
	q := client.Collection(r.collectionName).Query
	q = r.applyFilters(q, startDate, endDate, filters)

	if len(sorts) > 0 {
		for _, s := range sorts {
			direction := firestore.Asc
			if strings.ToLower(s.Order) == "desc" {
				direction = firestore.Desc
			}
			q = q.OrderBy(s.By, direction)
		}
	} else {
		q = q.OrderBy("createdAt", firestore.Desc)
	}

	iter := q.Documents(ctx)
	defer iter.Stop()

	var orders []*model.Order
	index := 0
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to iterate orders: %v", err)
		}

		if index < offset {
			index++
			continue
		}
		if len(orders) >= size {
			break
		}
		var order model.Order
		if err := doc.DataTo(&order); err != nil {
			continue
		}
		order.Id = doc.Ref.ID
		orders = append(orders, &order)
		index++
	}
	return orders, nil
}

func (r *OrderRepositoryImpl) Count(startDate, endDate *time.Time, filters []*dto.PageRequestFilter) (int64, error) {
	ctx := context.Background()
	client := database.GetFirestoreClient()
	q := client.Collection(r.collectionName).Query
	q = r.applyFilters(q, startDate, endDate, filters)

	iter := q.Documents(ctx)
	defer iter.Stop()

	var count int64
	for {
		_, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return 0, fmt.Errorf("failed to count orders: %v", err)
		}
		count++
	}
	return count, nil
}

func (r *OrderRepositoryImpl) GetNextDailyOrderNumber() (int, error) {
	ctx := context.Background()
	client := database.GetFirestoreClient()
	now := time.Now()
	year, month, day := now.Date()
	startOfDay := time.Date(year, month, day, 0, 0, 0, 0, now.Location())

	query := client.Collection(r.collectionName).
		Where("orderDate", "==", startOfDay).
		OrderBy("dailyOrderNumber", firestore.Desc).
		Limit(1)

	iter := query.Documents(ctx)
	defer iter.Stop()

	doc, err := iter.Next()
	if err == iterator.Done {
		return 1, nil
	}
	if err != nil {
		return 0, fmt.Errorf("failed to query for next daily order number: %v", err)
	}

	var lastOrder model.Order
	if err := doc.DataTo(&lastOrder); err != nil {
		return 0, fmt.Errorf("failed to convert document to order: %v", err)
	}
	return lastOrder.DailyOrderNumber + 1, nil
}

func (r *OrderRepositoryImpl) GetNextInvoiceNumber() (int64, error) {
	ctx := context.Background()
	client := database.GetFirestoreClient()
	query := client.Collection(r.collectionName).
		OrderBy("invoiceNumber", firestore.Desc).
		Limit(1)

	iter := query.Documents(ctx)
	defer iter.Stop()

	doc, err := iter.Next()
	if err == iterator.Done {
		return 1, nil
	}
	if err != nil {
		return 0, fmt.Errorf("failed to query for next invoice number: %v", err)
	}

	var lastOrder model.Order
	if err := doc.DataTo(&lastOrder); err != nil {
		return 0, fmt.Errorf("failed to convert document to order for invoice number: %v", err)
	}
	if lastOrder.InvoiceNumber == nil {
		return 1, nil
	}
	return *lastOrder.InvoiceNumber + 1, nil
}
