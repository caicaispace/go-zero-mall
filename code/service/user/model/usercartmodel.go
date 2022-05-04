package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"mall/common/library/gorm"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userCartFieldNames          = builder.RawFieldNames(&UserCart{})
	userCartRows                = strings.Join(userCartFieldNames, ",")
	userCartRowsExpectAutoSet   = strings.Join(stringx.Remove(userCartFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	userCartRowsWithPlaceHolder = strings.Join(stringx.Remove(userCartFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheUserCartIdPrefix                                  = "cache:userCart:id:"
	cacheUserCartUserIdSourceIdSourceTypeLicenseTypePrefix = "cache:userCart:userId:sourceId:sourceType:licenseType:"
)

type (
	UserCartModel interface {
		Insert(data *UserCart) (sql.Result, error)
		FindOne(id int64) (*UserCart, error)
		FindAllByUid(id int64) ([]*UserCart, error)
		FindAllByUidWithUserInfo(id int64) ([]*UserCart, error)
		FindOneByUserIdSourceIdSourceTypeLicenseType(userId int64, sourceId int64, sourceType int64, licenseType int64) (*UserCart, error)
		Update(data *UserCart) error
		Delete(id int64) error
	}

	defaultUserCartModel struct {
		sqlc.CachedConn
		table string
	}

	UserCart struct {
		Id          int64     `db:"id"`
		UserId      int64     `db:"user_id"`
		Username    string    `db:"username"`
		SourceId    int64     `db:"source_id"`
		SourceType  int64     `db:"source_type"`
		LicenseType int64     `db:"license_type"` // 授权类型(100:个人,200:企业,210企业plus,300:单张)
		VideoRate   int64     `db:"video_rate"`   // 等级\\0 : 其他\\1 : 1280x720\\2 : 1920x1080\\3 : 4096x2169(4k)\\4 : 8192x4320(8k)\\5 : 2048x1080(2k)\\6 : 3840x2160(UHD 4K)\\7 : 7680x4320(UHD 8K)
		SourceNum   int64     `db:"source_num"`   // 素材数量(用于购买多个)
		IsDel       int64     `db:"is_del"`
		DeletedTime int64     `db:"deleted_time"`
		DeletedAt   time.Time `db:"deleted_at"`
		CreatedTime int64     `db:"created_time"`
		CreatedAt   time.Time `db:"created_at"`
		UpdatedTime int64     `db:"updated_time"`
		UpdatedAt   time.Time `db:"updated_at"`
	}
)

func NewUserCartModel(conn sqlx.SqlConn, c cache.CacheConf) UserCartModel {
	return &defaultUserCartModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user_cart`",
	}
}

func (m *defaultUserCartModel) Insert(data *UserCart) (sql.Result, error) {
	userCartIdKey := fmt.Sprintf("%s%v", cacheUserCartIdPrefix, data.Id)
	userCartUserIdSourceIdSourceTypeLicenseTypeKey := fmt.Sprintf("%s%v:%v:%v:%v", cacheUserCartUserIdSourceIdSourceTypeLicenseTypePrefix, data.UserId, data.SourceId, data.SourceType, data.LicenseType)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userCartRowsExpectAutoSet)
		return conn.Exec(query, data.UserId, data.SourceId, data.SourceType, data.LicenseType, data.VideoRate, data.SourceNum, data.IsDel, data.DeletedTime, data.DeletedAt, data.CreatedTime, data.CreatedAt, data.UpdatedTime, data.UpdatedAt)
	}, userCartUserIdSourceIdSourceTypeLicenseTypeKey, userCartIdKey)
	return ret, err
}

func (m *defaultUserCartModel) FindOne(id int64) (*UserCart, error) {
	userCartIdKey := fmt.Sprintf("%s%v", cacheUserCartIdPrefix, id)
	var resp UserCart
	err := m.QueryRow(&resp, userCartIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userCartRows, m.table)
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

func (m *defaultUserCartModel) FindAllByUid(uid int64) ([]*UserCart, error) {
	var resp []*UserCart
	query := fmt.Sprintf("select %s from %s where `user_id` = ?", userCartRows, m.table)
	err := m.QueryRowsNoCache(&resp, query, uid)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
	// var resp []*UserCart
	// query := fmt.Sprintf(`
	// select %s,user.name as username from %s inner join user on user_cart.user_id = user.id where user_id = ?
	// `, helper.SqlFieldWithTableName(userCartRows, m.table, "username"), m.table)
	// log.Println("query", query)
	// err := m.QueryRowsNoCache(&resp, query, uid)
	// switch err {
	// case nil:
	// 	return resp, nil
	// case sqlc.ErrNotFound:
	// 	return nil, ErrNotFound
	// default:
	// 	return nil, err
	// }
}

func (m *defaultUserCartModel) FindAllByUidWithUserInfo(uid int64) ([]*UserCart, error) {
	resp := make([]*UserCart, 0)
	table := gorm.DB().Table(m.table)
	total := int64(0)
	table.Select("count(*)").Count(&total)
	table.Select(`
user_cart.id AS id,
user_cart.source_id AS source_id,
user.name AS username,
user_cart.user_id as user_id,
user_cart.source_id as source_id,
user_cart.source_type as source_type,
user_cart.license_type as license_type,
user_cart.video_rate as video_rate,
user_cart.source_num as source_num,
user_cart.is_del as is_del,
user_cart.deleted_time as deleted_time,
user_cart.deleted_at as deleted_at,
user_cart.created_time as created_time,
user_cart.created_at as created_at,
user_cart.updated_time as updated_time,
user_cart.updated_at as updated_at
`)
	table.Joins("left join user ON user_cart.user_id = user.id")
	table.Limit(10)
	table.Find(&resp)
	return resp, nil
}

func (m *defaultUserCartModel) FindOneByUserIdSourceIdSourceTypeLicenseType(userId int64, sourceId int64, sourceType int64, licenseType int64) (*UserCart, error) {
	userCartUserIdSourceIdSourceTypeLicenseTypeKey := fmt.Sprintf("%s%v:%v:%v:%v", cacheUserCartUserIdSourceIdSourceTypeLicenseTypePrefix, userId, sourceId, sourceType, licenseType)
	var resp UserCart
	err := m.QueryRowIndex(&resp, userCartUserIdSourceIdSourceTypeLicenseTypeKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? and `source_id` = ? and `source_type` = ? and `license_type` = ? limit 1", userCartRows, m.table)
		if err := conn.QueryRow(&resp, query, userId, sourceId, sourceType, licenseType); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserCartModel) Update(data *UserCart) error {
	userCartIdKey := fmt.Sprintf("%s%v", cacheUserCartIdPrefix, data.Id)
	userCartUserIdSourceIdSourceTypeLicenseTypeKey := fmt.Sprintf("%s%v:%v:%v:%v", cacheUserCartUserIdSourceIdSourceTypeLicenseTypePrefix, data.UserId, data.SourceId, data.SourceType, data.LicenseType)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userCartRowsWithPlaceHolder)
		return conn.Exec(query, data.UserId, data.SourceId, data.SourceType, data.LicenseType, data.VideoRate, data.SourceNum, data.IsDel, data.DeletedTime, data.DeletedAt, data.CreatedTime, data.CreatedAt, data.UpdatedTime, data.UpdatedAt, data.Id)
	}, userCartUserIdSourceIdSourceTypeLicenseTypeKey, userCartIdKey)
	return err
}

func (m *defaultUserCartModel) Delete(id int64) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}

	userCartUserIdSourceIdSourceTypeLicenseTypeKey := fmt.Sprintf("%s%v:%v:%v:%v", cacheUserCartUserIdSourceIdSourceTypeLicenseTypePrefix, data.UserId, data.SourceId, data.SourceType, data.LicenseType)
	userCartIdKey := fmt.Sprintf("%s%v", cacheUserCartIdPrefix, id)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, userCartIdKey, userCartUserIdSourceIdSourceTypeLicenseTypeKey)
	return err
}

func (m *defaultUserCartModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserCartIdPrefix, primary)
}

func (m *defaultUserCartModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userCartRows, m.table)
	return conn.QueryRow(v, query, primary)
}
