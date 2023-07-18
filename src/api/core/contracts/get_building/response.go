package get_building

import "flay-api-v3.0/src/api/core/contracts/common"

type Response struct {
	Buildings []common.Building `json:"buildings"`
}
