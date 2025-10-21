package loader

import (
	"github.com/ryantangit/sjsubark/etl/transform"
)

type loader interface {
	Upload(cgr transform.CompleteGarageRecord)
}
