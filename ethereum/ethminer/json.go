package ethminer

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"net/http"
	"strconv"
)

type jsonRequest struct {
	Method  string          `json:"method"`
	Version string          `json:"jsonrpc"`
	Id      json.RawMessage `json:"id,omitempty"`
	Payload json.RawMessage `json:"params,omitempty"`
}

type jsonError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type jsonErrResponse struct {
	Version string      `json:"jsonrpc"`
	Id      interface{} `json:"id,omitempty"`
	Error   jsonError   `json:"error"`
}

type jsonSuccessResponse struct {
	Version string      `json:"jsonrpc"`
	Id      interface{} `json:"id,omitempty"`
	Result  interface{} `json:"result"`
}

func createResponse(id interface{}, reply interface{}) interface{} {
	return &jsonSuccessResponse{Version: jsonrpcVersion, Id: id, Result: reply}
}

func createErrorResponse(id interface{}, err Error) interface{} {
	return &jsonErrResponse{Version: jsonrpcVersion, Id: id, Error: jsonError{Code: err.ErrorCode(), Message: err.Error()}}
}

func extractRPCMsg(r *http.Request) (string, json.RawMessage, interface{}, Error) {
	var incomingMsg json.RawMessage
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&incomingMsg); err != nil {
		return "", json.RawMessage{}, nil, &invalidRequestError{err.Error()}
	}

	var in jsonRequest
	if err := json.Unmarshal(incomingMsg, &in); err != nil {
		return "", json.RawMessage{}, nil, &invalidMessageError{err.Error()}
	}

	if err := checkReqId(in.Id); err != nil {
		return "", json.RawMessage{}, nil, &invalidMessageError{err.Error()}
	}
	return in.Method, in.Payload, in.Id, nil
}

func checkReqId(reqId json.RawMessage) error {
	if len(reqId) == 0 {
		return fmt.Errorf("missing request id")
	}
	if _, err := strconv.ParseFloat(string(reqId), 64); err == nil {
		return nil
	}
	var str string
	if err := json.Unmarshal(reqId, &str); err == nil {
		return nil
	}
	return fmt.Errorf("invalid request id")
}

func parseHashrateArguments(payload json.RawMessage) (hexutil.Uint64, common.Hash, Error) {
	args := [2]json.RawMessage{}
	if e := json.Unmarshal(payload, &args); e != nil {
		return 0, common.Hash{}, &invalidMessageError{e.Error()}
	}
	hashrate := hexutil.Uint64(0)
	if e := json.Unmarshal(args[0], &hashrate); e != nil {
		return 0, common.Hash{}, &invalidMessageError{e.Error()}
	}
	id := common.Hash{}
	if e := json.Unmarshal(args[1], &id); e != nil {
		return 0, common.Hash{}, &invalidMessageError{e.Error()}
	}
	return hashrate, id, nil
}

func parseWorkArguments(payload json.RawMessage) (types.BlockNonce, common.Hash, common.Hash, Error) {
	args := [3]json.RawMessage{}
	if e := json.Unmarshal(payload, &args); e != nil {
		return types.BlockNonce{}, common.Hash{}, common.Hash{}, &invalidMessageError{e.Error()}
	}
	nonce := types.BlockNonce{}
	if e := json.Unmarshal(args[0], &nonce); e != nil {
		return types.BlockNonce{}, common.Hash{}, common.Hash{}, &invalidMessageError{e.Error()}
	}
	hash := common.Hash{}
	if e := json.Unmarshal(args[1], &hash); e != nil {
		return types.BlockNonce{}, common.Hash{}, common.Hash{}, &invalidMessageError{e.Error()}
	}
	mixDigest := common.Hash{}
	if e := json.Unmarshal(args[2], &mixDigest); e != nil {
		return types.BlockNonce{}, common.Hash{}, common.Hash{}, &invalidMessageError{e.Error()}
	}
	return nonce, hash, mixDigest, nil
}
