package insights

import (
	"context"
	"ddd/pkg/standard"

	"github.com/soreing/trex"
)

type traceExtractor struct {
}

func (ex *traceExtractor) ExtractTraceInfo(
	ctx context.Context,
) (ver, tid, pid, rid, flg string) {
  raw := ctx.Value(standard.TRACE_INFO_KEY)
  if raw == nil {
    ver = "00"
    tid = "0000000000000000"
    pid = "00000000"
    rid = "00000000"
    flg = "00"
    return
  }
  tr, ok := raw.(trex.TxModel); if !ok {
    ver = "00"
    tid = "0000000000000000"
    pid = "00000000"
    rid = "00000000"
    flg = "00"
    return
  }
  return tr.Ver, tr.Tid, tr.Pid, tr.Rid, tr.Flg
}
