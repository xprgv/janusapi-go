package transaction

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/xprgv/janusapi-go/pkg/model"
)

func handleTxHandleInfo(janusResponse model.JanusResponse, tx Transaction, msg []byte) {
	switch janusResponse.Janus {
	case model.JanusMessageSuccess:
		handleInfo := model.HandleInfo{}
		if err := json.Unmarshal(msg, &handleInfo); err != nil {
			tx.Cb.Err(fmt.Errorf("error while unparse response: %w", err))
			return
		}
		tx.Cb.Ok(handleInfo)
	case model.JanusMessageError:
		janusResponseError := model.JanusResponseError{}
		if err := json.Unmarshal(msg, &janusResponseError); err != nil {
			tx.Cb.Err(err)
			return
		}
		tx.Cb.Err(errors.New(janusResponseError.ErrorData.Reason))
	default:
		tx.Cb.Err(fmt.Errorf("unknown response type: [%s]", janusResponse.Janus))
	}
}
