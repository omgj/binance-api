package binance

import (
	"github.com/BtStar/binance-api/Utils"
	"net/http"
	"testing"
)

var ba = New(http.DefaultClient, "", "")

func TestBinance_GetTicker(t *testing.T) {
	ticker, _ := ba.GetTicker(utils.LTC_BTC)
	t.Log(ticker)
}
func TestBinance_LimitBuy(t *testing.T) {
	order, err := ba.LimitBuy("11", "1", utils.LTC_BTC)
	t.Log(order, err)
}

func TestBinance_GetDepth(t *testing.T) {
	dep, err := ba.GetDepth(5, utils.ETH_BTC)
	t.Log(err)
	if err == nil {
		t.Log(dep.AskList)
		t.Log(dep.BidList)
	}
	dep, err = ba.GetDepth(2, utils.ETH_BTC)
	t.Log(err)
	if err == nil {
		t.Log(dep.AskList)
		t.Log(dep.BidList)
	}
}

func TestBinance_GetAccount(t *testing.T) {
	account, err := ba.GetAccount()
	t.Log(account, err)
}
