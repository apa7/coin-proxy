package util

import (
	// "../ethhelp"
	"bytes"
	"encoding/gob"
	"math/big"
	// "regexp"
	"strconv"
	"time"
	//"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereumproject/go-ethereum/common"
)

// var pow256 = ethhelp.BigPow(2, 256)
// var addressPattern = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
// var zeroHash = regexp.MustCompile("^0?x?0+$")

//
// func IsValidHexAddress(s string) bool {
// 	if IsZeroHash(s) || !addressPattern.MatchString(s) {
// 		return false
// 	}
// 	return true
// }
func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil

}

// func IsZeroHash(s string) bool {
// 	return zeroHash.MatchString(s)
// }

func MakeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// func GetTargetHex(diff int64) string {
// 	difficulty := big.NewInt(diff)
// 	diff1 := new(big.Int).Div(pow256, difficulty)
// 	return string(ethhelp.ToHex(diff1.Bytes()))
// }

// func TargetHexToDiff(targetHex string) *big.Int {
// 	targetBytes := ethhelp.FromHex(targetHex)
// 	return new(big.Int).Div(pow256, ethhelp.BytesToBig(targetBytes))
// }

func ToHex(n int64) string {
	return "0x0" + strconv.FormatInt(n, 16)
}

func FormatReward(reward *big.Int) string {
	return reward.String()
}

// func FormatRatReward(reward *big.Rat) string {
// 	wei := new(big.Rat).SetInt(ethhelp.Ether)
// 	reward = reward.Quo(reward, wei)
// 	return reward.FloatString(8)
// }

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func MustParseDuration(s string) time.Duration {
	value, err := time.ParseDuration(s)
	if err != nil {
		panic("util: Can't parse duration `" + s + "`: " + err.Error())
	}
	return value
}