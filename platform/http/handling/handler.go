package handling

import "platform/pipeline"

type Handler interface {
	Execute(*pipeline.ComponentContext)
}
