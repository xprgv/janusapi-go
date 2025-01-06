package transaction

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/xprgv/janusapi-go/pkg/model"
)

func handleTxWatch(janusResponse model.JanusResponse, tx Transaction, msg []byte) {
	switch janusResponse.Janus {
	case model.JanusMessageAck: // janus successfully received request
	case model.JanusMessageEvent:
		janusMessage := model.JanusResponseMessage{}
		if err := json.Unmarshal(msg, &janusMessage); err != nil {
			tx.Cb.Err(err)
			return
		}
		tx.Cb.Ok(janusMessage.Jsep)
	default:
		tx.Cb.Err(fmt.Errorf("unknown response type: [%s]", janusResponse.Janus))
	}
}

func handleTxStart(janusResponse model.JanusResponse, tx Transaction, msg []byte) {
	switch janusResponse.Janus {
	case model.JanusMessageAck: // janus successfully received request
	case model.JanusMessageEvent:
		janusEvent := model.JanusResponseEvent{}
		if err := json.Unmarshal(msg, &janusEvent); err != nil {
			tx.Cb.Err(err)
			return
		}
		tx.Cb.Ok(nil)
	default:
		tx.Cb.Err(fmt.Errorf("unknown response type: [%s]", janusResponse.Janus))
	}
}

func handleTxPause(janusResponse model.JanusResponse, tx Transaction, msg []byte) {
	switch janusResponse.Janus {
	case model.JanusMessageAck: // janus successfully received request
	case model.JanusMessageEvent:
		janusEvent := model.JanusResponseEvent{}
		if err := json.Unmarshal(msg, &janusEvent); err != nil {
			tx.Cb.Err(err)
			return
		}
		tx.Cb.Ok(nil)
	default:
		tx.Cb.Err(fmt.Errorf("unknown response type: [%s]", janusResponse.Janus))
	}
}

func handleTxResume(janusResponse model.JanusResponse, tx Transaction, msg []byte) {
	switch janusResponse.Janus {
	case model.JanusMessageAck: // janus successfully received request
	case model.JanusMessageEvent:
		janusEvent := model.JanusResponseEvent{}
		if err := json.Unmarshal(msg, &janusEvent); err != nil {
			tx.Cb.Err(err)
			return
		}
		tx.Cb.Ok(nil)
	default:
		tx.Cb.Err(fmt.Errorf("unknown response type: [%s]", janusResponse.Janus))
	}
}

func handleTxSwitch(janusResponse model.JanusResponse, tx Transaction, msg []byte) {
	switch janusResponse.Janus {
	case model.JanusMessageAck: // janus successfully received request
	case model.JanusMessageEvent:
		janusEvent := model.JanusResponseEvent{}
		if err := json.Unmarshal(msg, &janusEvent); err != nil {
			tx.Cb.Err(err)
			return
		}
		if err, exist := janusEvent.PluginData.Data["error"]; exist {
			errMsg, ok := err.(string)
			if !ok {
				tx.Cb.Err(errors.New("can't convert error message to string"))
				return
			}
			tx.Cb.Err(fmt.Errorf("switch error: %s", errMsg))
			return
		}
		tx.Cb.Ok(nil)
	default:
		tx.Cb.Err(fmt.Errorf("unknown response type: [%s]", janusResponse.Janus))
	}
}

func handleTxIceCandidate(janusResponse model.JanusResponse, tx Transaction, msg []byte) {
	switch janusResponse.Janus {
	case model.JanusMessageAck: // do nothing, janus successfully received the trickle
		tx.Cb.Ok(nil)
	// case janusMessageSuccess:
	//	tx.Cb.Ok(nil)
	case model.JanusMessageError:
		janusResponseError := model.JanusResponseError{}
		if err := json.Unmarshal(msg, &janusResponseError); err != nil {
			tx.Cb.Err(err)
			return
		}
		tx.Cb.Err(errors.New(janusResponseError.ErrorData.Reason))
	}
}
