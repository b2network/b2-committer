package types

type BTCTxOutputs struct {
	TxID         string         `json:"txid"`
	Vout         uint64         `json:"vout"`
	Address      string         `json:"address"`
	CodeType     uint64         `json:"codeType"`
	Satoshi      uint64         `json:"satoshi"`
	ScriptType   string         `json:"scriptType"`
	Height       uint64         `json:"height"`
	Idx          uint64         `json:"idx"`
	Inscriptions []Inscriptions `json:"inscriptions"`
}

type Inscriptions struct {
	InscriptionNumber uint64 `json:"inscriptionNumber"`
	InscriptionID     string `json:"inscriptionId"`
	Offset            uint64 `json:"offset"`
	Moved             bool   `json:"moved"`
	Sequence          uint64 `json:"sequence"`
	IsBRC20           bool   `json:"isBRC20"`
}

type APIBTCTxOutputs struct {
	Data []BTCTxOutputs `json:"data"`
}
