// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
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
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheDb2UserIdPrefix       = "cache:db2:user:id:"
	cacheDb2UserUserIdPrefix   = "cache:db2:user:userId:"
	cacheDb2UserUsernamePrefix = "cache:db2:user:username:"
)

type (
	userModel interface {
		Insert(ctx context.Context, data *User) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*User, error)
		FindOneByUserId(ctx context.Context, userId int64) (*User, error)
		FindOneByUsername(ctx context.Context, username string) (*User, error)
		Update(ctx context.Context, data *User) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}

	User struct {
		Id         int64          `db:"id"`
		UserId     int64          `db:"user_id"`
		Username   string         `db:"username"`
		Password   string         `db:"password"`
		Email      sql.NullString `db:"email"`
		Gender     int64          `db:"gender"`
		CreateTime time.Time      `db:"create_time"`
		UpdateTime time.Time      `db:"update_time"`
	}
)

func newUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultUserModel {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`user`",
	}
}

func (m *defaultUserModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	db2UserIdKey := fmt.Sprintf("%s%v", cacheDb2UserIdPrefix, id)
	db2UserUserIdKey := fmt.Sprintf("%s%v", cacheDb2UserUserIdPrefix, data.UserId)
	db2UserUsernameKey := fmt.Sprintf("%s%v", cacheDb2UserUsernamePrefix, data.Username)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, db2UserIdKey, db2UserUserIdKey, db2UserUsernameKey)
	return err
}

func (m *defaultUserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	db2UserIdKey := fmt.Sprintf("%s%v", cacheDb2UserIdPrefix, id)
	var resp User
	err := m.QueryRowCtx(ctx, &resp, db2UserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
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

func (m *defaultUserModel) FindOneByUserId(ctx context.Context, userId int64) (*User, error) {
	db2UserUserIdKey := fmt.Sprintf("%s%v", cacheDb2UserUserIdPrefix, userId)
	var resp User
	err := m.QueryRowIndexCtx(ctx, &resp, db2UserUserIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", userRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userId); err != nil {
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

func (m *defaultUserModel) FindOneByUsername(ctx context.Context, username string) (*User, error) {
	db2UserUsernameKey := fmt.Sprintf("%s%v", cacheDb2UserUsernamePrefix, username)
	var resp User
	err := m.QueryRowIndexCtx(ctx, &resp, db2UserUsernameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `username` = ? limit 1", userRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, username); err != nil {
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

func (m *defaultUserModel) Insert(ctx context.Context, data *User) (sql.Result, error) {
	db2UserIdKey := fmt.Sprintf("%s%v", cacheDb2UserIdPrefix, data.Id)
	db2UserUserIdKey := fmt.Sprintf("%s%v", cacheDb2UserUserIdPrefix, data.UserId)
	db2UserUsernameKey := fmt.Sprintf("%s%v", cacheDb2UserUsernamePrefix, data.Username)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.Username, data.Password, data.Email, data.Gender)
	}, db2UserIdKey, db2UserUserIdKey, db2UserUsernameKey)
	return ret, err
}

func (m *defaultUserModel) Update(ctx context.Context, newData *User) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	db2UserIdKey := fmt.Sprintf("%s%v", cacheDb2UserIdPrefix, data.Id)
	db2UserUserIdKey := fmt.Sprintf("%s%v", cacheDb2UserUserIdPrefix, data.UserId)
	db2UserUsernameKey := fmt.Sprintf("%s%v", cacheDb2UserUsernamePrefix, data.Username)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.UserId, newData.Username, newData.Password, newData.Email, newData.Gender, newData.Id)
	}, db2UserIdKey, db2UserUserIdKey, db2UserUsernameKey)
	return err
}

func (m *defaultUserModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheDb2UserIdPrefix, primary)
}

func (m *defaultUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserModel) tableName() string {
	return m.table
}
