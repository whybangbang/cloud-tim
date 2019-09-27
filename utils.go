package cloud_tim

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)


var (
	randMax = big.NewInt(1000000000)
)

func buildSig(appKey, mobile, random, time string) string {
	str := fmt.Sprintf("appkey=%v&random=%v&time=%v&mobile=%v", appKey, random, time, mobile)
	sha := sha256.New()
	sha.Write([]byte(str))
	return fmt.Sprintf("%x", sha.Sum(nil))
}

func random() int64 {
	//p, err := rand.Prime(rand.Reader, 20)
	p, err := rand.Int(rand.Reader, randMax)
	if err != nil {
		// it's so rare
		fmt.Println("random err")
		return 0
	}

	return p.Int64()
}