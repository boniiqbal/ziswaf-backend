package create_transaction

import (
	"strconv"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	CreateTransactionResponse struct {
		base.BaseResponse
		Data CreateTransactionResponseData `json:"data"`
	}

	CreateTransactionResponseData struct {
		ID            uint64 `json:"id"`
		TransactionID string `json:"transaction_id"`
	}
)

func SetResponse(data CreateTransactionResponseData, message string, success bool) CreateTransactionResponse {
	return CreateTransactionResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: data,
	}
}

func ResponseMapper(trxID uint64, dvsID uint16) CreateTransactionResponseData {
	transactionID := TransactionFormat(trxID, dvsID)
	return CreateTransactionResponseData{
		ID:            trxID,
		TransactionID: transactionID,
	}
}

func TransactionFormat(ID uint64, divisionID uint16) string {
	transactionID := strconv.FormatUint(ID, 10)

	switch divisionID {
	case 1:
		return transactionID + "-UPZ"
	case 2:
		return transactionID + "-RTL"
	case 3:
		return transactionID + "-CP"
	default:
		return transactionID + "-RTL"
	}
}
