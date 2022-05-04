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
	userVipFieldNames          = builder.RawFieldNames(&UserVip{})
	userVipRows                = strings.Join(userVipFieldNames, ",")
	userVipRowsExpectAutoSet   = strings.Join(stringx.Remove(userVipFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	userVipRowsWithPlaceHolder = strings.Join(stringx.Remove(userVipFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheUserVipIdPrefix = "cache:userVip:id:"
)

type (
	UserVipModel interface {
		Insert(data *UserVip) (sql.Result, error)
		FindOne(id int64) (*UserVip, error)
		FindAllByUid(id int64) ([]*UserVip, error)
		Update(data *UserVip) error
		Delete(id int64) error
	}

	defaultUserVipModel struct {
		sqlc.CachedConn
		table string
	}

	UserVip struct {
		Id            int64     `db:"id"`
		UserId        int64     `db:"user_id"`
		VipId         int64     `db:"vip_id"`
		VipType       int64     `db:"vip_type"`   // 对应 vip 表 type
		VideoId       int64     `db:"video_id"`   // 视频 id（单张购买使用）
		OrderId       int64     `db:"order_id"`   // 订单 id
		LicenseId     int64     `db:"license_id"` // cd_user_vip_license表id
		StartTime     int64     `db:"start_time"` // 起效时间
		StartAt       time.Time `db:"start_at"`
		EndTime       int64     `db:"end_time"` // 失效时间
		EndAt         time.Time `db:"end_at"`
		DayLimit      int64     `db:"day_limit"`       // 日下载限制
		TotalLimit    int64     `db:"total_limit"`     // 总下载限制
		LastAdminUser string    `db:"last_admin_user"` // 最后一个操作vip的人
		Remark        string    `db:"remark"`          // 备注{为什么添加}
		IsDel         int64     `db:"is_del"`          // 是否删除
		CreatedTime   int64     `db:"created_time"`
		CreatedAt     time.Time `db:"created_at"`
		UpdatedTime   int64     `db:"updated_time"`
		UpdatedAt     time.Time `db:"updated_at"`
		DeletedTime   int64     `db:"deleted_time"`
		DeletedAt     time.Time `db:"deleted_at"`
	}
)

func NewUserVipModel(conn sqlx.SqlConn, c cache.CacheConf) UserVipModel {
	return &defaultUserVipModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user_vip`",
	}
}

func (m *defaultUserVipModel) Insert(data *UserVip) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userVipRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.UserId, data.VipId, data.VipType, data.VideoId, data.OrderId, data.LicenseId, data.StartTime, data.StartAt, data.EndTime, data.EndAt, data.DayLimit, data.TotalLimit, data.LastAdminUser, data.Remark, data.IsDel, data.CreatedTime, data.CreatedAt, data.UpdatedTime, data.UpdatedAt, data.DeletedTime, data.DeletedAt)

	return ret, err
}

func (m *defaultUserVipModel) FindOne(id int64) (*UserVip, error) {
	userVipIdKey := fmt.Sprintf("%s%v", cacheUserVipIdPrefix, id)
	var resp UserVip
	err := m.QueryRow(&resp, userVipIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userVipRows, m.table)
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

func (m *defaultUserVipModel) FindAllByUid(uid int64) ([]*UserVip, error) {
	var resp []*UserVip
	query := fmt.Sprintf("select %s from %s where `user_id` = ?", userVipRows, m.table)
	err := m.QueryRowsNoCache(&resp, query, uid)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserVipModel) Update(data *UserVip) error {
	userVipIdKey := fmt.Sprintf("%s%v", cacheUserVipIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userVipRowsWithPlaceHolder)
		return conn.Exec(query, data.UserId, data.VipId, data.VipType, data.VideoId, data.OrderId, data.LicenseId, data.StartTime, data.StartAt, data.EndTime, data.EndAt, data.DayLimit, data.TotalLimit, data.LastAdminUser, data.Remark, data.IsDel, data.CreatedTime, data.CreatedAt, data.UpdatedTime, data.UpdatedAt, data.DeletedTime, data.DeletedAt, data.Id)
	}, userVipIdKey)
	return err
}

func (m *defaultUserVipModel) Delete(id int64) error {
	userVipIdKey := fmt.Sprintf("%s%v", cacheUserVipIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, userVipIdKey)
	return err
}

func (m *defaultUserVipModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserVipIdPrefix, primary)
}

func (m *defaultUserVipModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userVipRows, m.table)
	return conn.QueryRow(v, query, primary)
}
