package transaction

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/xprgv/janusapi-go/pkg/model"
	"github.com/xprgv/janusapi-go/pkg/plugin"
)

// nolint dupl
func handleTxListMountpoint(janusResponse model.JanusResponse, tx Transaction, msg []byte) {
	switch janusResponse.Janus {
	case model.JanusMessageSuccess:
		janusResponseSuccess := model.JanusResponseSuccess{}
		if err := json.Unmarshal(msg, &janusResponseSuccess); err != nil {
			tx.Cb.Err(err)
			return
		}
		switch janusResponseSuccess.PluginData.Plugin {
		case plugin.JanusPluginStreaming:
			switch janusResponseSuccess.PluginData.Data["streaming"] {
			case "list":
				bin, err := json.Marshal(janusResponseSuccess.PluginData.Data["list"])
				if err != nil {
					tx.Cb.Err(err)
					return
				}
				mountpointsList := &[]model.MountpointShort{}
				if err := json.Unmarshal(bin, mountpointsList); err != nil {
					tx.Cb.Err(err)
					return
				}
				tx.Cb.Ok(mountpointsList)
			case model.JanusMessageEvent:
				errMsg, ok := janusResponseSuccess.PluginData.Data["error"].(string)
				if !ok {
					tx.Cb.Err(errors.New("can't convert error message to string"))
					return
				}
				tx.Cb.Err(errors.New(errMsg))
			default:
				tx.Cb.Err(fmt.Errorf("not implemented"))
			}
		default:
			tx.Cb.Err(fmt.Errorf("handlers for plugin [%s] not implemented", janusResponseSuccess.PluginData.Plugin))
		}
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

// nolint dupl
func handleTxInfoMountpoint(janusResponse model.JanusResponse, tx Transaction, msg []byte) {
	switch janusResponse.Janus {
	case model.JanusMessageSuccess:
		janusResponseSuccess := model.JanusResponseSuccess{}
		if err := json.Unmarshal(msg, &janusResponseSuccess); err != nil {
			tx.Cb.Err(err)
			return
		}
		switch janusResponseSuccess.PluginData.Plugin {
		case plugin.JanusPluginStreaming:
			switch janusResponseSuccess.PluginData.Data["streaming"] {
			case "info":
				bin, err := json.Marshal(janusResponseSuccess.PluginData.Data["info"])
				if err != nil {
					tx.Cb.Err(err)
					return
				}
				mountpointInfo := &model.Mountpoint{}
				if err := json.Unmarshal(bin, mountpointInfo); err != nil {
					tx.Cb.Err(err)
					return
				}
				tx.Cb.Ok(mountpointInfo)
			case model.JanusMessageEvent:
				errMsg, ok := janusResponseSuccess.PluginData.Data["error"].(string)
				if !ok {
					tx.Cb.Err(errors.New("can't convert error message to string"))
					return
				}
				tx.Cb.Err(errors.New(errMsg))
			default:
				tx.Cb.Err(fmt.Errorf("not implemented"))
			}
		default:
			tx.Cb.Err(fmt.Errorf("handlers for plugin [%s] not implemented", janusResponseSuccess.PluginData.Plugin))
		}
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

func handleTxDestroyMountpoint(janusResponse model.JanusResponse, tx Transaction, msg []byte) {
	switch janusResponse.Janus {
	case model.JanusMessageSuccess:
		janusResponseSuccess := model.JanusResponseSuccess{}
		if err := json.Unmarshal(msg, &janusResponseSuccess); err != nil {
			tx.Cb.Err(err)
			return
		}
		switch janusResponseSuccess.PluginData.Plugin {
		case plugin.JanusPluginStreaming:
			switch janusResponseSuccess.PluginData.Data["streaming"] {
			case "destroyed":
				tx.Cb.Ok(nil)
			case model.JanusMessageEvent:
				errMsg, ok := janusResponseSuccess.PluginData.Data["error"].(string)
				if !ok {
					tx.Cb.Err(errors.New("can't convert error message to string"))
					return
				}
				tx.Cb.Err(errors.New(errMsg))
			default:
				tx.Cb.Err(fmt.Errorf("not implemented"))
			}
		default:
			tx.Cb.Err(fmt.Errorf("handlers for plugin [%s] not implemented", janusResponseSuccess.PluginData.Plugin))
		}
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

// nolint dupl
func handleTxCreateMountpoint(janusResponse model.JanusResponse, tx Transaction, msg []byte) {
	switch janusResponse.Janus {
	case model.JanusMessageSuccess:
		janusResponseSuccess := model.JanusResponseSuccess{}
		if err := json.Unmarshal(msg, &janusResponseSuccess); err != nil {
			tx.Cb.Err(err)
			return
		}
		switch janusResponseSuccess.PluginData.Plugin {
		case plugin.JanusPluginStreaming:
			switch janusResponseSuccess.PluginData.Data["streaming"] {
			case "created":
				bin, err := json.Marshal(janusResponseSuccess.PluginData.Data["stream"])
				if err != nil {
					tx.Cb.Err(err)
					return
				}
				result := &model.PluginDataStreamingStream{}
				if err := json.Unmarshal(bin, result); err != nil {
					tx.Cb.Err(err)
					return
				}
				tx.Cb.Ok(result)
			case model.JanusMessageEvent:
				errMsg, ok := janusResponseSuccess.PluginData.Data["error"].(string)
				if !ok {
					tx.Cb.Err(errors.New("can't convert error message to string"))
					return
				}
				tx.Cb.Err(errors.New(errMsg))
			default:
				tx.Cb.Err(fmt.Errorf("not implemented"))
			}
		default:
			tx.Cb.Err(fmt.Errorf("handlers for plugin [%s] not implemented", janusResponseSuccess.PluginData.Plugin))
		}
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
