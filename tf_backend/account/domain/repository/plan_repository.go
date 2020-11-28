package repository

import "context"

type PlanRepository interface {
	FindCapacityById(context.Context, uint64) (uint64, error)
}
