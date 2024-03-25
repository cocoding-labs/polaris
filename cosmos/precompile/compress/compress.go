package compress

import (
	"bytes"
	"compress/gzip"
	"context"
	"io"

	generated "github.com/berachain/polaris/contracts/bindings/forma/precompile/compress"
	ethprecompile "github.com/berachain/polaris/eth/core/precompile"
	"github.com/ethereum/go-ethereum/common"
)

type Compress struct {
	ethprecompile.BaseContract
}

func NewCompress() *Compress {
	return &Compress{
		BaseContract: ethprecompile.NewBaseContract(
			generated.CompressABI,
			common.HexToAddress("0x0F043A0000000000000000000000000000000002"),
		),
	}
}

func (c *Compress) Deflate(ctx context.Context, input []byte) ([]byte, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	_, err := gz.Write(input)
	if err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (c *Compress) Inflate(ctx context.Context, input []byte) ([]byte, error) {
	var buf bytes.Buffer
	gz, err := gzip.NewReader(bytes.NewBuffer(input))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(&buf, gz)
	if err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
