package loader
import (
	"github.com/ryantangit/sjsubark/etl/transform"	
)

type loader interface {
	upload(cgr transform.CompleteGarageRecord)
}

