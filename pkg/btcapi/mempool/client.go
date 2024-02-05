package mempool

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/b2network/b2committer/pkg/btcapi"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/pkg/errors"
)

type Client struct {
	baseURL string
}

func NewClient(netParams *chaincfg.Params) *Client {
	baseURL := ""
	switch netParams.Net {
	case wire.MainNet:
		baseURL = "https://mempool.space/api"
	case wire.TestNet3:
		baseURL = "https://mempool.space/testnet/api"
	case chaincfg.SigNetParams.Net:
		baseURL = "https://mempool.space/signet/api"
	default:
		log.Fatal("mem-pool don't support other netParams")
	}
	return &Client{
		baseURL: baseURL,
	}
}

func (c *Client) request(method, subPath string, requestBody io.Reader) ([]byte, error) {
	return btcapi.Request(method, c.baseURL, subPath, requestBody)
}

func (c *Client) GetRawTransaction(txHash *chainhash.Hash) (*wire.MsgTx, error) {
	res, err := c.request(http.MethodGet, fmt.Sprintf("/tx/%s/raw", txHash.String()), nil)
	if err != nil {
		return nil, err
	}

	tx := wire.NewMsgTx(wire.TxVersion)
	if err := tx.Deserialize(bytes.NewReader(res)); err != nil {
		return nil, err
	}
	return tx, nil
}

func (c *Client) BroadcastTx(tx *wire.MsgTx) (*chainhash.Hash, error) {
	var buf bytes.Buffer
	if err := tx.Serialize(&buf); err != nil {
		return nil, err
	}

	res, err := c.request(http.MethodPost, "/tx", strings.NewReader(hex.EncodeToString(buf.Bytes())))
	if err != nil {
		return nil, err
	}

	txHash, err := chainhash.NewHashFromStr(string(res))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to parse tx hash, %s", string(res)))
	}
	return txHash, nil
}

func (c *Client) GetTransactionByID(id string) (*btcapi.Transaction, error) {
	res, err := c.request(http.MethodGet, fmt.Sprintf("/tx/%s", id), nil)
	if err != nil {
		return nil, err
	}
	tx := &btcapi.Transaction{}
	err = json.Unmarshal(res, &tx)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (c *Client) GetCurrentBlockHash() (int64, error) {
	res, err := c.request(http.MethodGet, "/blocks/tip/height", nil)
	if err != nil {
		return 0, err
	}
	var blockHeight int64
	err = json.Unmarshal(res, &blockHeight)
	if err != nil {
		return 0, err
	}
	return blockHeight, nil
}

var _ btcapi.Client = (*Client)(nil)
