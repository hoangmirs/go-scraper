package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/hoangmirs/go-scraper/controllers:Keyword"] = append(beego.GlobalControllerRouter["github.com/hoangmirs/go-scraper/controllers:Keyword"],
        beego.ControllerComments{
            Method: "Show",
            Router: "/keyword/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
