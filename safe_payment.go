package safe_payments

import (
	"fmt"
)

func Pay(PaymentDetails pd) {
	fmt.Println("sourceAccount:", pd.SourceAccount)
	fmt.Println("destinationAccount:", pd.DestinationAccount)
	fmt.Println("amount:", pd.Amount)
}
