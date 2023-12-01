package dict

import (
	"fmt"
)

func ExampleAccountCodeFind() {
	fmt.Println(AccountCodeFind("408"))
	fmt.Println(AccountCodeFind("40817"))
	fmt.Println(AccountCodeFind("000"))

	// Output: {408 Прочие счета}
	// {40817 Физические лица}
	// {000 Неизвестный код счета}

}
