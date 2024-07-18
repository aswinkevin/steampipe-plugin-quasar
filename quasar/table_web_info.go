package quasar

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableWebInfo() *plugin.Table {
	return &plugin.Table{
		Name:        "web_info",
		Description: "Website reachability and ssl information",
		List: &plugin.ListConfig{
			Hydrate: listWebInfo,
		},
		Columns: []*plugin.Column{
			{Name: "hostname", Type: proto.ColumnType_STRING, Description: "The hostname of the website"},
			{Name: "status_code", Type: proto.ColumnType_STRING, Description: "The HTTP status code of the website"},
		},
	}
}

func listWebInfo(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, ok := d.Connection.Config.(quasarConfig)
	if !ok {
		return nil, nil
	}

	for _, domain := range conn.GetDomains() {
		d.StreamListItem(ctx, map[string]interface{}{
			"hostname":    domain,
			"status_code": "200",
		})
	}

	return nil, nil
}
