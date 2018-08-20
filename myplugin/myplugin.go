package myplugin

import (
	"github.com/Peripli/service-manager/pkg/web"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type MyPlugin struct{}

func (p *MyPlugin) Name() string { return "MyPlugin" }

func (p *MyPlugin) FetchCatalog(req *web.Request, next web.Handler) (*web.Response, error) {
	res, err := next.Handle(req)
	if err != nil {
		return nil, err
	}
	serviceName := gjson.GetBytes(res.Body, "services.0.name").String()
	res.Body, err = sjson.SetBytes(res.Body, "services.0.name", serviceName+"-suffix")
	return res, err
}
