package transaction

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/xprgv/janusapi-go/pkg/errs"
	"github.com/xprgv/janusapi-go/pkg/model"
)

func handleTxCreate(janusResponse model.JanusResponse, tx Transaction, msg []byte) {
	switch janusResponse.Janus {
	case model.JanusMessageSuccess:
		janusResponseSuccess := model.JanusResponseSuccess{}
		if err := json.Unmarshal(msg, &janusResponseSuccess); err != nil {
			tx.Cb.Err(err)
			return
		}
		tx.Cb.Ok(janusResponseSuccess.Data.Id)
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

func handleTxAttach(janusResponse model.JanusResponse, tx Transaction, msg []byte) {
	switch janusResponse.Janus {
	case model.JanusMessageSuccess:
		janusResponseSuccess := model.JanusResponseSuccess{}
		if err := json.Unmarshal(msg, &janusResponseSuccess); err != nil {
			tx.Cb.Err(err)
			return
		}
		tx.Cb.Ok(janusResponseSuccess.Data.Id)
	case model.JanusMessageError:
		janusResponseError := model.JanusResponseError{}
		if err := json.Unmarshal(msg, &janusResponseError); err != nil {
			tx.Cb.Err(err)
			return
		}
		tx.Cb.Err(errors.New(janusResponseError.ErrorData.Reason))
		switch janusResponseError.ErrorData.Code {
		case errs.JanusErrorSessionNotFound:
			tx.Cb.Err(errs.ErrSessionNotFound)
		default:
			tx.Cb.Err(errors.New(janusResponseError.ErrorData.Reason))
		}
	default:
		tx.Cb.Err(fmt.Errorf("unknown response type: [%s]", janusResponse.Janus))
	}
}

func handleTxKeepalive(janusResponse model.JanusResponse, tx Transaction, msg []byte) {
	switch janusResponse.Janus {
	case model.JanusMessageAck:
		tx.Cb.Ok(nil)
	default:
		tx.Cb.Err(fmt.Errorf("unknown response type: [%s]", janusResponse.Janus))
	}
}

func handleTxDestroy(janusResponse model.JanusResponse, tx Transaction, msg []byte) {
	switch janusResponse.Janus {
	case model.JanusMessageSuccess:
		tx.Cb.Ok(nil)
	default:
		tx.Cb.Err(fmt.Errorf("unknown response type: [%s]", janusResponse.Janus))
	}
}
