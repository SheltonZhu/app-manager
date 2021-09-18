package bcrypt

import (
	"fmt"
	"testing"
	"time"
)

func TestCheckPasswordHash(t *testing.T) {
	passwd := "docker.admin"
	h, _ := HashPassword(passwd)
	fmt.Println("Hash:", h)
	cost := GetHashingCost([]byte(h))
	fmt.Println("Cost: ", cost)
	hash := "$2a$11$XPxyTOksR.LXqqjVEs1vCefPFwGFruNXqZB.qkscfGQz90OTMNePy"
	n := time.Now()
	ok := CheckPasswordHash(passwd, hash)
	elapsed := time.Since(n)
	fmt.Println("Ok: ", ok)
	fmt.Println("Elapsed:", elapsed)
}
