//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/rudsonsouzas/pos-fullcycle-clean-architecture/internal/entity"
	"github.com/rudsonsouzas/pos-fullcycle-clean-architecture/internal/event"
	"github.com/rudsonsouzas/pos-fullcycle-clean-architecture/internal/infra/database"
	"github.com/rudsonsouzas/pos-fullcycle-clean-architecture/internal/infra/web"
	"github.com/rudsonsouzas/pos-fullcycle-clean-architecture/internal/usecase"
	"github.com/rudsonsouzas/pos-fullcycle-clean-architecture/pkg/events"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	event.NewOrdersListed,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventInterface), new(*event.OrdersListed)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

var setOrdersListedEvent = wire.NewSet(
	event.NewOrdersListed,
	wire.Bind(new(events.EventInterface), new(*event.OrdersListed)),
)

func NewListOrdersUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.ListOrdersUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrdersListedEvent,
		usecase.NewListOrdersUseCase,
	)
	return &usecase.ListOrdersUseCase{}
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		setOrdersListedEvent,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}
