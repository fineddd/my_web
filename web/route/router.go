package route

import (
	"my_web/web/handler"
//	"my_web/web/handler/active"
//	"my_web/web/handler/data_table"
//	"my_web/web/handler/order"
//	"my_web/web/handler/player_action"
//	"my_web/web/handler/real_time"
//	"my_web/web/handler/common_data_query"
	"net/http"
)

func StaticRoute() {
	http.Handle("/dist/", http.FileServer(http.Dir("view")))
	http.Handle("/bootstrap/", http.FileServer(http.Dir("view")))
	http.Handle("/pages/", http.FileServer(http.Dir("view")))
	http.Handle("/plugins/", http.FileServer(http.Dir("view")))
	http.Handle("/documentation/", http.FileServer(http.Dir("view")))
	http.Handle("/assets/", http.FileServer(http.Dir("view")))
}

func DynamicRoute() {
	//index
	http.HandleFunc("/", handler.RootHdl)
	http.HandleFunc("/index2", handler.Index2Hdl)
	http.HandleFunc("/index", handler.IndexHdl)
	http.HandleFunc("/login", handler.LoginHdl)

	http.HandleFunc("/account_manage", handler.AcntManageHdl)
	http.HandleFunc("/user/list", handler.AcntListHdl)
	http.HandleFunc("/right/list", handler.RightListHdl)
	http.HandleFunc("/platform/list", handler.PlatformListHdl)
	/*
	http.HandleFunc("/login", handler.LoginHdl)
	http.HandleFunc("/logout", handler.LogoutHdl)
	http.HandleFunc("/layout_language_bar", handler.LayoutLanguageBarHdl)
	http.HandleFunc("/table_base", handler.TableBaseHdl)
	http.HandleFunc("/chart", handler.ChartHdl)
	// platform
	http.HandleFunc("/platform/index", handler.PlatformIndexHdl)
	http.HandleFunc("/platform/list", handler.GetPlatformList)
	http.HandleFunc("/platform/add", handler.PlatformAdd)
	http.HandleFunc("/platform/del", handler.PlatformDel)
	http.HandleFunc("/platform/update", handler.PlatformUpdate)

	// source
	http.HandleFunc("/source/index", handler.SourceIndexHdl)
	http.HandleFunc("/source/list", handler.GetSourceList)
	http.HandleFunc("/source/add", handler.SourceAdd)
	http.HandleFunc("/source/del", handler.SourceDel)
	http.HandleFunc("/source/update", handler.SourceUpdate)

	// server
	http.HandleFunc("/server/index", handler.ServerIndexHdl)
	http.HandleFunc("/server/list", handler.GetServerList)
	http.HandleFunc("/server/get_list_and_refresh", handler.GetServerListAndRefresh)
	http.HandleFunc("/server/add", handler.ServerAdd)
	http.HandleFunc("/server/del", handler.ServerDel)
	http.HandleFunc("/server/update", handler.ServerUpdate)
	http.HandleFunc("/server/get_by_pfid", handler.GetServerListByPlatformID)
	// http.HandleFunc("/server/get_names_by_ids", handler.GetServerNamesByIDs)
	// user
	http.HandleFunc("/user/index", handler.UserIndexHdl)
	http.HandleFunc("/user/list", handler.GetUserList)
	http.HandleFunc("/user/add", handler.UserAdd)
	http.HandleFunc("/user/del", handler.UserDel)
	http.HandleFunc("/user/update", handler.UserUpdate)
	http.HandleFunc("/user/selfpasswordindex", handler.UserSelfPasswordIndexHdl)
	http.HandleFunc("/user/selfpasswordupdate", handler.UserSelfPasswordUpdate)
	// right
	http.HandleFunc("/right/index", handler.RightIndexHdl)
	http.HandleFunc("/right/list", handler.GetRightList)
	http.HandleFunc("/right/add", handler.RightAdd)
	http.HandleFunc("/right/del", handler.RightDel)
	http.HandleFunc("/right/update", handler.RightUpdate)
	//player
	http.HandleFunc("/player/gmindex", handler.GmCmdIndexHdl)
	http.HandleFunc("/player/gmcmd", handler.GmCmdHdl)
	http.HandleFunc("/player/gmhelpindex", handler.GmHelpIndexHdl)
	http.HandleFunc("/player/querygmhelp", handler.QueryGmHelpHdl)
	http.HandleFunc("/player/addgmhelpreply", handler.AddGmHelpReplyHdl)
	http.HandleFunc("/player/baseinfoindex", handler.PlayerBaseInfoIndexHdl)
	http.HandleFunc("/player/querybaseinfo", handler.QueryPlayerBaseInfo)
	http.HandleFunc("/player/itemindex", handler.PlayerItemIndexHdl)
	http.HandleFunc("/player/queryitem", handler.QueryPlayerItem)
	http.HandleFunc("/player/resindex", handler.PlayerResIndexHdl)
	http.HandleFunc("/player/queryres", handler.QueryPlayerRes)

	http.HandleFunc("/player/gmforbiddenindex", handler.GmForbiddenIndexHdl)
	http.HandleFunc("/player/gmcontrolforbidden", handler.GmControlForbidden)

	http.HandleFunc("/player/gmforbiddentalkindex", handler.GmForbiddenTalkIndexHdl)
	http.HandleFunc("/player/gmcontrolforbiddentalk", handler.GmControlForbiddenTalk)

	// chart
	http.HandleFunc("/chart/chart_data", handler.GetChartData)
	http.HandleFunc("/my_chart", handler.MyChartHdl)
	// compensation
	http.HandleFunc("/compensation/index", handler.CompensationIndexHdl)
	http.HandleFunc("/compensation/compensation_list", handler.GetCompensationList)
	http.HandleFunc("/compensation/add", handler.CompensationAdd)
	http.HandleFunc("/compensation/stop", handler.CompensationStop)
	// notice
	http.HandleFunc("/notice/index", handler.NoticeIndexHdl)
	http.HandleFunc("/notice/notice_list", handler.GetNoticeList)
	http.HandleFunc("/notice/add", handler.NoticeAdd)
	http.HandleFunc("/notice/stop", handler.NoticeStop)
	// data table
	http.HandleFunc("/datatable/online_index", data_table.OnlineIndex)
	http.HandleFunc("/datatable/online", data_table.OnlineDT)
	http.HandleFunc("/datatable/outflow_index", data_table.OutflowIndex)
	http.HandleFunc("/datatable/outflow_query", data_table.OutflowQuery)
	http.HandleFunc("/datatable/grade_index", data_table.GradeIndex)
	http.HandleFunc("/datatable/grade_query", data_table.GradeQuery)
	http.HandleFunc("/datatable/task_index", data_table.TaskIndex)
	http.HandleFunc("/datatable/task_query", data_table.TaskQuery)
	http.HandleFunc("/datatable/single_server_index", data_table.SingleServerIndex)
	http.HandleFunc("/datatable/single_server_query", data_table.SingleServerQuery)
	http.HandleFunc("/datatable/all_server_index", data_table.AllServerIndex)
	http.HandleFunc("/datatable/all_server_query", data_table.AllServerQuery)
	http.HandleFunc("/datatable/platform_index", data_table.PlatformIndex)
	http.HandleFunc("/datatable/platform_query", data_table.PlatformQuery)
	//cdk
	http.HandleFunc("/cdk/cdk_index", handler.CDKIndexHdl)
	http.HandleFunc("/cdk/add_hd_cdk", handler.CDKAddHDCDK)
	http.HandleFunc("/cdk/add_cdk", handler.CDKAddCDK)
	http.HandleFunc("/cdk/get_all_cdk_batch", handler.CDKGetAllCDKBatch)
	http.HandleFunc("/cdk/get_cdk_keys", handler.CDKGetCDKKeys)
	http.HandleFunc("/cdk/get_hd_cdk_keys", handler.CDKGetHDCDKKeys)
	http.HandleFunc("/cdk/modify_hd_batch", handler.CDKModifyHDBatch)
	http.HandleFunc("/cdk/modify_batch", handler.CDKModifyBatch)
	//real_time
	http.HandleFunc("/real_time/real_time_online_index", real_time.RealTimeOnlineIndexHdl)
	http.HandleFunc("/real_time/real_time_online_query", real_time.RealTimeOnlineQueryHdl)
	http.HandleFunc("/real_time/real_time_new_player_index", real_time.RealTimeNewPlayerIndexHdl)
	http.HandleFunc("/real_time/real_time_new_player_query", real_time.RealTimeNewPlayerQueryHdl)
	http.HandleFunc("/real_time/real_time_pay_index", real_time.RealTimePayIndexHdl)
	http.HandleFunc("/real_time/real_time_pay_query", real_time.RealTimePayQueryHdl)

	//player_action
	http.HandleFunc("/player_action/player_action_consume_index", player_action.PlayerActionConsumeIndexHdl)
	http.HandleFunc("/player_action/player_action_consume_query", player_action.PlayerActionConsumeQueryHdl)
	http.HandleFunc("/player_action/player_action_player_guide_index", player_action.PlayerActionPlayerGuideIndexHdl)
	http.HandleFunc("/player_action/player_action_player_guide_query", player_action.PlayerActionPlayerGuideQueryHdl)
	http.HandleFunc("/player_action/player_action_pay_distribute_index", player_action.PlayerActionPayDistributeIndexHdl)
	http.HandleFunc("/player_action/player_action_pay_distribute_query", player_action.PlayerActionPayDistributeQueryHdl)
	http.HandleFunc("/player_action/player_action_player_exist_index", player_action.PlayerActionPlayerExistIndexHdl)
	http.HandleFunc("/player_action/player_action_player_exist_query", player_action.PlayerActionPlayerExistQueryHdl)
	http.HandleFunc("/player_action/player_action_accumulate_pay_index", player_action.PlayerActionAccumulatePayIndexHdl)
	http.HandleFunc("/player_action/player_action_accumulate_pay_query", player_action.PlayerActionAccumulatePayQueryHdl)
	http.HandleFunc("/player_action/player_action_vip_index", player_action.PlayerActionVipIndexHdl)
	http.HandleFunc("/player_action/player_action_vip_query", player_action.PlayerActionVipQueryHdl)
	http.HandleFunc("/player_action/player_action_vip_count_index", player_action.PlayerActionVipCountIndexHdl)
	http.HandleFunc("/player_action/player_action_vip_count_query", player_action.PlayerActionVipCountQueryHdl)

	http.HandleFunc("/active/active_config_index", active.ActiveConfigIndexHdl)
	http.HandleFunc("/active/active_config_list", active.GetActiveConfigList)
	http.HandleFunc("/active/active_config_add", active.ActiveConfigAddHdl)
	http.HandleFunc("/active/active_config_stop", active.ActiveConfigStopHdl)

	http.HandleFunc("/order/order_index", order.OrderIndexHdl)
	http.HandleFunc("/order/order_query", order.QueryOrderListHdl)
	http.HandleFunc("/order/order_user_index", order.OrderUserIndexHdl)
	http.HandleFunc("/order/order_user_query", order.QueryOrderUserListHdl)
	
	http.HandleFunc("/common_mongo_query_index", common_data_query.CommonMongoQueryIndexHdl)
	http.HandleFunc("/common_mongo_query", common_data_query.CommonMongoQueryHdl)
	*/
}
