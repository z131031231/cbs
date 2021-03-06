// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/zeromicro/cds/cmd/galaxy/internal/svc"
	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/user/add",
				Handler: addHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/dm-list",
				Handler: dmListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/dm-add",
				Handler: dmAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/list-tables",
				Handler: listTableHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/list-databases",
				Handler: listDatabasesHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/database-list",
				Handler: databaseListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/galaxy/html/default-config",
				Handler: defaultConfigHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/generate-create-sql",
				Handler: generateCreateSqlHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/exec-sql",
				Handler: execSqlHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/dm-stop",
				Handler: dmStopHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/dm-delete",
				Handler: dmDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/dm-redo",
				Handler: dmRedoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/rtu-list",
				Handler: rtuListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/rtu-add",
				Handler: rtuAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/rtu-stop",
				Handler: rtuStopHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/rtu-delete",
				Handler: rtuDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/rtu-redo",
				Handler: rtuRedoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/connector-list",
				Handler: connectorListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/connector-add",
				Handler: connectorAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/connector-delete",
				Handler: connectorDeleteHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/user/login",
				Handler: loginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/accesscenter/get-user-info",
				Handler: getUserInfoHandler(serverCtx),
			},
		},
	)
}
