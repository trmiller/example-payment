package safe_payments

import (
	"fmt"
)

type PaymentDetails struct {
	SourceAccount      string
	DestinationAccount string
	Amount             string
}

func Pay(pd *PaymentDetails) {
	fmt.Println("sourceAccount:", pd.SourceAccount)
	fmt.Println("destinationAccount:", pd.DestinationAccount)
	fmt.Println("amount:", pd.Amount)
}
