package handler

import (
	"github.com/alifradityar/votr/votr-api/domain/topic"
)

// Root should list all the handler that we will use
type Root struct {
	*topic.Handler `inject:""`
}
