package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	vipFieldNames          = builder.RawFieldNames(&Vip{})
	vipRows                = strings.Join(vipFieldNames, ",")
	vipRowsExpectAutoSet   = strings.Join(stringx.Remove(vipFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	vipRowsWithPlaceHolder = strings.Join(stringx.Remove(vipFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheVipIdPrefix = "cache:vip:id:"
)

type (
	VipModel interface {
		Insert(data *Vip) (sql.Result, error)
		FindOne(id int64) (*Vip, error)
		Update(data *Vip) error
		Delete(id int64) error
	}

	defaultVipModel struct {
		sqlc.CachedConn
		table string
	}

	Vip struct {
		Id            int64     `db:"id"`
		Title         string    `db:"title"`          // 标题
		TitleSimple   string    `db:"title_simple"`   // 简写标题
		Desc          string    `db:"desc"`           // 描述
		Appliance     string    `db:"appliance"`      // 适用范围
		Tp            int64     `db:"type"`           // 1:免费 100:个人 200:企业 300:单张个人 400:单张企业
		ClassType     int64     `db:"class_type"`     // 分类 vip 类型
		AuthId        int64     `db:"auth_id"`        // 授权类型
		Level         int64     `db:"level"`          // 等级
		Sort          int64     `db:"sort"`           // 排序
		Price         float64   `db:"price"`          // 价格
		OriginalPrice float64   `db:"original_price"` // 原价
		DayLimit      int64     `db:"day_limit"`      // 日下载限制
		TotalLimit    int64     `db:"total_limit"`    // 总下载限制
		ResetCycle    int64     `db:"reset_cycle"`    // 重置周期（单位/天）
		ParentId      int64     `db:"parent_id"`      // 父级id
		SingleId      int64     `db:"single_id"`      // VIP对应单张 id (0代表单个购买价格，大于0代表VIP价格(其中1代表无对应售卖的单款价格)）
		IsUsable      int64     `db:"is_usable"`      // 是否可用
		CreatedAt     time.Time `db:"created_at"`
		UpdatedTime   int64     `db:"updated_time"`
		UpdatedAt     time.Time `db:"updated_at"`
		CreatedTime   int64     `db:"created_time"`
	}
)

func NewVipModel(conn sqlx.SqlConn, c cache.CacheConf) VipModel {
	return &defaultVipModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`vip`",
	}
}

func (m *defaultVipModel) Insert(data *Vip) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, vipRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Title, data.TitleSimple, data.Desc, data.Appliance, data.Tp, data.ClassType, data.AuthId, data.Level, data.Sort, data.Price, data.OriginalPrice, data.DayLimit, data.TotalLimit, data.ResetCycle, data.ParentId, data.SingleId, data.IsUsable, data.CreatedAt, data.UpdatedTime, data.UpdatedAt, data.CreatedTime)

	return ret, err
}

func (m *defaultVipModel) FindOne(id int64) (*Vip, error) {
	vipIdKey := fmt.Sprintf("%s%v", cacheVipIdPrefix, id)
	var resp Vip
	err := m.QueryRow(&resp, vipIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", vipRows, m.table)
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultVipModel) Update(data *Vip) error {
	vipIdKey := fmt.Sprintf("%s%v", cacheVipIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, vipRowsWithPlaceHolder)
		return conn.Exec(query, data.Title, data.TitleSimple, data.Desc, data.Appliance, data.Tp, data.ClassType, data.AuthId, data.Level, data.Sort, data.Price, data.OriginalPrice, data.DayLimit, data.TotalLimit, data.ResetCycle, data.ParentId, data.SingleId, data.IsUsable, data.CreatedAt, data.UpdatedTime, data.UpdatedAt, data.CreatedTime, data.Id)
	}, vipIdKey)
	return err
}

func (m *defaultVipModel) Delete(id int64) error {
	vipIdKey := fmt.Sprintf("%s%v", cacheVipIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, vipIdKey)
	return err
}

func (m *defaultVipModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheVipIdPrefix, primary)
}

func (m *defaultVipModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", vipRows, m.table)
	return conn.QueryRow(v, query, primary)
}
