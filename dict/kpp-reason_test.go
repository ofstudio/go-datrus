package dict

import "fmt"

func ExampleKPPReasonFind() {
	fmt.Println(KPPReasonFind("36"))
	fmt.Println(KPPReasonFind("00"))

	// Output: {36 Постановка на учет в налоговом органе организации при выполнении соглашения о разделе продукции}
	// {00 Неизвестная причина}
}
