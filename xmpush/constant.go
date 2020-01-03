package xmpush

const (
	ProductURL = "https://api.xmpush.xiaomi.com"
	SandboxURL = "https://sandbox.xmpush.xiaomi.com"

	ContentTypeFORM = "application/x-www-form-urlencoded;charset=UTF-8"
)

const (
	// pass_through的值可以为：
	// 0 表示通知栏消息
	// 1 表示透传消息
	PassThroughNotify = int64(0)
	PassThroughCustom = int64(1)
)
const (
	Version01 = "/v1"
	Version02 = "/v2"
	Version03 = "/v3"
)

const (
	URIRegId                             = "/message/regid"                // 向某个regid或一组regid列表推送某条消息
	URIMultiMessagesRegIds               = "/multi_messages/regids"        // 针对不同的regid推送不同的消息
	URIMultiMessagesAlias                = "/multi_messages/aliases"       // 针对不同的aliases推送不同的消息
	URIMultiMessagesUserAccount          = "/multi_messages/user_accounts" // 针对不同的accounts推送不同的消息
	URIMessageAlisa                      = "/message/alias"                // 根据alias，发送消息到指定设备上
	URIMessageUserAccount                = "/message/user_account"         // 根据account，发送消息到指定account上
	URIMultiPackageNameMessageMultiTopic = "/message/multi_topic"          // 根据topic，发送消息到指定一组设备上
	URIMessageMultiTopic                 = "/message/topic"                // 根据topic，发送消息到指定一组设备上
	URIMultiPackageNameMessageAll        = "/message/all"                  // 向所有设备推送某条消息
	URIMessageAll                        = "/message/all"                  // 向所有设备推送某条消息
	URIMultiTopic                        = "/message/multi_topic"          // 向多个topic广播消息
	URIScheduleJobExist                  = "/schedule_job/exist"           // 检测定时消息的任务是否存在。
	URIScheduleJobDelete                 = "/schedule_job/delete"          // 删除指定的定时消息。
	URIScheduleJobDeleteByJobKey         = "/schedule_job/delete"          // 删除指定的定时消息。

)

const (
	StatsURL          = "/stats/message/counters" // 统计push
	MessageStatusURL  = "/trace/message/status"   // 获取指定ID的消息状态
	MessagesStatusURL = "/trace/messages/status"  // 获取某个时间间隔内所有消息的状态
)

const (
	// notify_type的值可以是DEFAULT_ALL或者以下其他几种的OR组合：
	Notify_Type_All     = -1
	Notify_Type_Sound   = 1 // 使用默认提示音提示；
	Notify_Type_Vibrate = 2 // 使用默认震动提示；
	Notify_Type_Lights  = 4 // 使用默认led灯光提示；
)

const (
	// topic之间的操作关系。支持以下三种：
	// UNION并集
	// INTERSECTION交集
	// EXCEPT差集
	TopicUnion        = "UNION"
	TopicIntersection = "INTERSECTION"
	TopicExcept       = "EXCEPT"
)
