package clickhouse

import (
	"github.com/kshvakov/clickhouse/lib/binary"
	"github.com/kshvakov/clickhouse/lib/data"
	"github.com/kshvakov/clickhouse/lib/protocol"
)

func (ch *clickhouse) readBlock(decoder *binary.Decoder) (*data.Block, error) {
	if ch.ServerInfo.Revision >= protocol.DBMS_MIN_REVISION_WITH_TEMPORARY_TABLES {
		if _, err := decoder.String(); err != nil {
			return nil, err
		}
	}
	if ch.compress {

	}
	var block data.Block
	if err := block.Read(&ch.ServerInfo, decoder); err != nil {
		return nil, err
	}
	return &block, nil
}
