// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/tal-tech/cds/galaxy/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/user/login",
				Handler: LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/user/add",
				Handler: AddHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/galaxy/user/get",
				Handler: GetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/dm-list",
				Handler: DmListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/dm-add",
				Handler: DmAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/list-tables",
				Handler: ListTableHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/list-databases",
				Handler: ListDatabasesHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/database-list",
				Handler: DatabaseListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/galaxy/html/default-config",
				Handler: DefaultConfigHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/generate-create-sql",
				Handler: GenerateCreateSqlHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/exec-sql",
				Handler: ExecSqlHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/dm-stop",
				Handler: DmStopHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/dm-delete",
				Handler: DmDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/dm-redo",
				Handler: DmRedoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/rtu-list",
				Handler: RtuListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/rtu-add",
				Handler: RtuAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/rtu-stop",
				Handler: RtuStopHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/rtu-delete",
				Handler: RtuDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/rtu-redo",
				Handler: RtuRedoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/connector-list",
				Handler: ConnectorListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/connector-add",
				Handler: ConnectorAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/galaxy/html/connector-delete",
				Handler: ConnectorDeleteHandler(serverCtx),
			},
		},
	)
}
