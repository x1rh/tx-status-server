package ton

import "testing"

func TestDefaultTestnetGetTransaction(t *testing.T) {
	lt := uint64(26725328000003)
	address := "UQCwxyBrDh2FL3MycXF0-XgInMp1aMbND_1SIggGS2cVjAlV"
	txHash := "746AA9D08540713EEC53646655ED216F5CACCC6CDA042D6BD865FE6B8A2C4C1C"

	tx, err := DefaultTestnetGetTransaction(lt, address, txHash)
	if err != nil {
		t.Error(err)
	}

	t.Logf("tx: %+v\n", tx)
}

func TestDefaultMainnetGetTransaction(t *testing.T) {

}
