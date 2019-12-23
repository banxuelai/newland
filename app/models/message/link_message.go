package message

import (
	"time"
)

type LinkMessage struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11) 'id'"`
	Type         int       `json:"type" xorm:"not null default 0 comment('消息类型 0未知 1进度提醒\u00a0 2共享提醒\u00a0 3超期提醒\u00a0 4网签提醒') TINYINT(4) 'type'"`
	SubType      int       `json:"sub_type" xorm:"not null default 0 TINYINT(4) 'sub_type'"`
	CustomerId   int64     `json:"customer_id" xorm:"not null default 0 comment('客户ID') BIGINT(20) 'customer_id'"`
	AgentId      int64     `json:"agent_id" xorm:"not null default 0 comment('经纪人ID') BIGINT(20) 'agent_id'"`
	RelationId   int64     `json:"relation_id" xorm:"not null default 0 comment('关联id') BIGINT(20) 'relation_id'"`
	RelationType int       `json:"relation_type" xorm:"not null default 0 comment('关联的主体类型，详见当前页面的：link 数据库实体表枚举值') TINYINT(3) 'relation_type'"`
	Content      string    `json:"content" xorm:"not null default '' comment('消息内容') VARCHAR(512) 'content'"`
	Images       string    `json:"images" xorm:"not null default '' comment('图片') VARCHAR(255) 'images'"`
	Url          string    `json:"url" xorm:"not null default '' comment('连接') VARCHAR(255) 'url'"`
	SendType     int       `json:"send_type" xorm:"not null default 0 comment('类型 0:文本消息1:通知号消息') TINYINT(4) 'send_type'"`
	Extend       string    `json:"extend" xorm:"not null default '' comment('扩展字段') VARCHAR(512) 'extend'"`
	ImUrl        string    `json:"im_url" xorm:"not null default '' comment('IM内的跳转连接') VARCHAR(255) 'im_url'"`
	Status       int       `json:"status" xorm:"not null default 0 comment('状态 ') TINYINT(4) 'status'"`
	Ctime        time.Time `json:"ctime" xorm:"not null default 'current_timestamp()' TIMESTAMP 'ctime'"`
	Mtime        time.Time `json:"mtime" xorm:"not null default 'current_timestamp()' TIMESTAMP 'mtime'"`
	Title        string    `json:"title" xorm:"not null default '' comment('消息标题') VARCHAR(50) 'title'"`
	AgentType    int       `json:"agent_type" xorm:"not null default 0 comment('经纪人类型：0 经纪人\u00a0 1 案场人员 2 渠道人员') TINYINT(2) 'agent_type'"`
}
