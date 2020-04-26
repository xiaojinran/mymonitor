package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["mymonitor/controllers:ConfigController"] = append(beego.GlobalControllerRouter["mymonitor/controllers:ConfigController"],
        beego.ControllerComments{
            Method: "SshConfig",
            Router: `/ssh`,
            AllowHTTPMethods: []string{"Get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mymonitor/controllers:ConfigController"] = append(beego.GlobalControllerRouter["mymonitor/controllers:ConfigController"],
        beego.ControllerComments{
            Method: "AddSshHost",
            Router: `/ssh`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mymonitor/controllers:LinuxController"] = append(beego.GlobalControllerRouter["mymonitor/controllers:LinuxController"],
        beego.ControllerComments{
            Method: "LinuxDetails",
            Router: `/details`,
            AllowHTTPMethods: []string{"Get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mymonitor/controllers:LinuxController"] = append(beego.GlobalControllerRouter["mymonitor/controllers:LinuxController"],
        beego.ControllerComments{
            Method: "LinuxOps",
            Router: `/ops`,
            AllowHTTPMethods: []string{"Get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mymonitor/controllers:LinuxController"] = append(beego.GlobalControllerRouter["mymonitor/controllers:LinuxController"],
        beego.ControllerComments{
            Method: "ExecCommands",
            Router: `/ops`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mymonitor/controllers:ProductController"] = append(beego.GlobalControllerRouter["mymonitor/controllers:ProductController"],
        beego.ControllerComments{
            Method: "GetProduct",
            Router: `/`,
            AllowHTTPMethods: []string{"Get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mymonitor/controllers:ProductController"] = append(beego.GlobalControllerRouter["mymonitor/controllers:ProductController"],
        beego.ControllerComments{
            Method: "AddProduct",
            Router: `/add`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mymonitor/controllers:SalesController"] = append(beego.GlobalControllerRouter["mymonitor/controllers:SalesController"],
        beego.ControllerComments{
            Method: "GetSales",
            Router: `/sales`,
            AllowHTTPMethods: []string{"Get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mymonitor/controllers:WechatController"] = append(beego.GlobalControllerRouter["mymonitor/controllers:WechatController"],
        beego.ControllerComments{
            Method: "WechatAPI",
            Router: `/send`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
