package resources

import (
	_ "embed"
	"github.com/ClickHouse/clickhouse-go/v2/lib/proto"
	"gopkg.in/yaml.v3"
	"strings"
)

type Meta struct {
	ClickhouseVersions []proto.Version `yaml:"clickhouse_versions"`
	GoVersions         []proto.Version `yaml:"go_versions"`
}

//go:embed meta.yml
var metaFile []byte
var ClientMeta Meta

func init() {
	if err := yaml.Unmarshal(metaFile, &ClientMeta); err != nil {
		panic(err)
	}
}

func (m *Meta) IsSupportedClickHouse(v proto.Version) bool {
	return true
}

func (m *Meta) SupportedVersions() string {
	versions := make([]string, len(m.ClickhouseVersions), len(m.ClickhouseVersions))
	for i := range m.ClickhouseVersions {
		versions[i] = m.ClickhouseVersions[i].String()
	}
	return strings.Join(versions, ",")
}
