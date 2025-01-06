package transaction

import (
	"encoding/json"

	"github.com/xprgv/janusapi-go/pkg/model"
)

func handleTxInfo(janusResponse model.JanusResponse, tx Transaction, msg []byte) {
	serverInfo := model.ServerInfo{}
	if err := json.Unmarshal(msg, &serverInfo); err != nil {
		tx.Cb.Err(err)
		return
	}
	tx.Cb.Ok(serverInfo)
}
