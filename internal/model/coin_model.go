package model

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"weather_mgr/helper"
	"weather_mgr/pkg/xprometheus"
	"weather_mgr/pkg/xzap"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

type coinModel struct {
}

// 金币的总数记录
type CoinAccount struct {
	UserId                uint64    `gorm:"column:user_id;primary_key"` // 用户 id
	AppName               string    `gorm:"column:app_name;NOT NULL"`   // app 名称
	Balance               int64     `gorm:"column:balance;default:0"`   // 金币余额
	AccountType           *string   `gorm:"column:account_type"`        // 账号类型, PHONE,OAUTH2_GOOGLE...
	AccountName           *string   `gorm:"column:account_name"`        // 账号名称
	HasRewardedFirstLogin bool      `gorm:"column:has_rewarded_first_login;default:0"`
	CreateTime            time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 数据创建的时间
	UpdateTime            time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 数据更新的时间
}

type CoinAccountByUid struct {
	Id                    uint64    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`   // 自增 ID
	UserId                uint64    `gorm:"column:user_id;NOT NULL" json:"user_id"`           // 用户 id
	AppName               string    `gorm:"column:app_name;NOT NULL" json:"app_name"`         // app 名称
	Balance               int       `gorm:"column:balance;default:0;NOT NULL" json:"balance"` // 金币余额
	AccountType           string    `gorm:"column:account_type;NOT NULL" json:"account_type"` // 账号类型, PHONE,OAUTH2_GOOGLE...
	AccountName           string    `gorm:"column:account_name;NOT NULL" json:"account_name"` // 账号名称
	HasRewardedFirstLogin int       `gorm:"column:has_rewarded_first_login;default:0;NOT NULL" json:"has_rewarded_first_login"`
	CreatedAt             time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"created_at"` // 数据创建的时间
	UpdatedAt             time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"updated_at"` // 数据更新的时间
}

func (c *coinModel) IncrCoin(ctx context.Context, uid uint64, appName string, coinNum int64, label string) (err error) {
	return c.IncrCoinNew(ctx, uid, appName, coinNum, label)
}

func (c *coinModel) IncrCoinNew(ctx context.Context, uid uint64, appName string, coinNum int64, label string) (err error) {
	xzap.Debug("", zap.Any("uid", uid))
	var id uint64
	tName := c.getCoinAccountTableName(uid)
	if appName == "" {
		row := migRwDbHandle.QueryRowContext(ctx, "select id,appName from "+tName+" where user_id=?", uid)
		err = row.Scan(&id, &appName)
		if err != nil {
			xzap.ErrorContext(ctx, "", zap.Error(err))
			return err
		}
	} else {
		row := migRwDbHandle.QueryRowContext(ctx, "select id from "+tName+" where user_id=? and app_name=?", uid, appName)
		err = row.Scan(&id)
		if err != nil {
			xzap.ErrorContext(ctx, "", zap.Error(err))
			return err
		}
	}

	if label == "first_login" {
		_, err = migRwDbHandle.ExecContext(ctx, "update "+tName+" set balance=balance+?,has_rewarded_first_login=1 where id=?", coinNum, id)
	} else {
		_, err = migRwDbHandle.ExecContext(ctx, "update "+tName+" set balance=balance+? where id=?", coinNum, id)
	}
	if err == nil {
		go xprometheus.AddCoinLabelCouter(label, appName, coinNum)
	}
	//go c.syncFromNewToOld(uid, appName)
	return err
}

func (c *coinModel) GetOrCreateCoinAccount(ctx context.Context, uid uint64, appName string, accountName string, accountType string) (account *CoinAccount, err error) {
	return c.GetOrCreateCoinAccountNew(ctx, uid, appName, accountName, accountType)
}

// 读新表
func (c *coinModel) GetOrCreateCoinAccountNew(ctx context.Context, uid uint64, appName string, accountName string, accountType string) (account *CoinAccount, err error) {
	account = new(CoinAccount)
	//accountByUid := new(CoinAccountByUid)
	var row *sql.Row
	tName := c.getCoinAccountTableName(uid)
	if appName == "" {
		row = migRwDbHandle.QueryRowContext(ctx, "select user_id,balance,app_name,has_rewarded_first_login,created_at,updated_at  from "+tName+" where user_id=? limit 1", uid)
	} else {
		row = migRwDbHandle.QueryRowContext(ctx, "select user_id,balance,app_name,has_rewarded_first_login,created_at,updated_at  from "+tName+" where user_id=? and app_name=? limit 1", uid, appName)
	}

	err = row.Scan(&account.UserId, &account.Balance, &account.AppName, &account.HasRewardedFirstLogin, &account.CreateTime, &account.UpdateTime)

	if err == sql.ErrNoRows {
		err = nil
		if appName == "" {
			return nil, helper.ERROR_INVALID
		}
		account = &CoinAccount{
			UserId:                uid,
			AppName:               appName,
			Balance:               0,
			HasRewardedFirstLogin: false,
			AccountType:           &accountType,
			AccountName:           &accountName,
			CreateTime:            time.Now(),
			UpdateTime:            time.Now(),
		}
		migRwDbHandle.ExecContext(ctx, "insert "+tName+" set user_id=?, app_name=?, balance=?,has_rewarded_first_login=?,account_type=?,account_name=?,created_at=?,updated_at=?", account.UserId, account.AppName, account.Balance, account.HasRewardedFirstLogin,
			account.AccountType, account.AccountName, account.CreateTime, account.UpdateTime)
	} else if err != nil {
		xzap.ErrorContext(ctx, "", zap.Error(err))
		return nil, err
	}
	return account, nil
}

func (c *coinModel) GetOrCreateCoinInfo(ctx context.Context, uid uint64, appName string, accountName string, accountType string) (balance int64, hasRewardedFirstLogin bool, appName2 string, err error) {
	return c.GetOrCreateCoinInfoNew(ctx, uid, appName, accountName, accountType)

}

func (c *coinModel) GetOrCreateCoinInfoNew(ctx context.Context, uid uint64, appName string, accountName string, accountType string) (balance int64, hasRewardedFirstLogin bool, appName2 string, err error) {
	ant, err := c.GetOrCreateCoinAccountNew(ctx, uid, appName, accountName, accountType)
	if err != nil {
		xzap.ErrorContext(ctx, "", zap.Error(err))
		return 0, false, appName, err
	}
	return ant.Balance, ant.HasRewardedFirstLogin, ant.AppName, nil
}

func (c *coinModel) GetCoinTotalStats(ctx context.Context, uid uint64) (totalAwardCoins int64, err error) {
	NewTName := c.getCoinStatsTableName(uid)

	row := migRwDbHandle.QueryRowContext(ctx, "select total_reward_coins from "+NewTName+" where user_id=? limit 1", uid)
	err = row.Scan(&totalAwardCoins)
	if err == nil || err == sql.ErrNoRows {
		err = nil
		return totalAwardCoins, nil
	}
	xzap.ErrorContext(ctx, "GetCoinTotalStats", zap.Error(err))
	return 0, err
}

func (c *coinModel) IncrCoinTotalStats(ctx context.Context, uid uint64, IncrCoin int64) (totalAwardCoins int64, err error) {
	NewTName := c.getCoinStatsTableName(uid)
	row := migRwDbHandle.QueryRowContext(ctx, "select total_reward_coins from  "+NewTName+" where user_id=? limit 1", uid)
	err = row.Scan(&totalAwardCoins)
	if err == sql.ErrNoRows {
		if IncrCoin > 0 {
			_, err = migRwDbHandle.ExecContext(ctx, "insert into  "+NewTName+" set total_reward_coins=? , user_id=?", IncrCoin, uid)
		} else {
			_, err = migRwDbHandle.ExecContext(ctx, "insert into  "+NewTName+" set total_consume_coins=? , user_id=?", IncrCoin, uid)
		}
		if err != nil {
			xzap.ErrorContext(ctx, "", zap.Error(err))
			return 0, err
		}
		return IncrCoin, nil
	} else if err == nil {
		if IncrCoin > 0 {
			_, err = migRwDbHandle.ExecContext(ctx, "update  "+NewTName+" set total_reward_coins=total_reward_coins+? where user_id=? limit 1", IncrCoin, uid)
		} else {
			_, err = migRwDbHandle.ExecContext(ctx, "update "+NewTName+" set total_consume_coins=total_consume_coins+? where user_id=? limit 1", IncrCoin, uid)
		}
		if err != nil {
			xzap.ErrorContext(ctx, "", zap.Error(err))
			return 0, err
		}
		return IncrCoin + totalAwardCoins, nil
	}
	xzap.ErrorContext(ctx, "", zap.Error(err))
	return 0, err
}

func (c *coinModel) getCoinAccountTableName(uid uint64) string {
	return fmt.Sprintf("coin_account_%.3d", uid%256)
}

func (c *coinModel) getCoinStatsTableName(uid uint64) string {
	return fmt.Sprintf("coin_trade_total_stats_%.3d", uid%256)
}

func (c *coinModel) getCoinAwardDailyStatsRedisKey(uid uint64, today string, label string) string {
	return fmt.Sprintf("coinAwardStats:%d:%s:%s", uid, today, label)
}

func (c *coinModel) GetDailyAwardStats(ctx context.Context, uid uint64, label string, today string) (stats map[string]int64, err error) {

	stats = map[string]int64{}
	k := c.getCoinAwardDailyStatsRedisKey(uid, today, label)
	v, err := redisClient.HGetAll(k).Result()
	if err == redis.Nil {
		return stats, err
	}
	if err != nil {
		xzap.ErrorContext(ctx, "", zap.Error(err))
		return nil, err
	}

	for kk, vv := range v {
		stats[kk], _ = helper.Str2int64(vv)
	}

	return stats, nil
}

func (c *coinModel) IncrDailyAwardStats(ctx context.Context, uid uint64, label string, today string, num int64) (stats map[string]int64, err error) {
	stats = map[string]int64{}
	k := c.getCoinAwardDailyStatsRedisKey(uid, today, label)
	stats["amount"], err = redisClient.HIncrBy(k, "amount", num).Result()
	if err != nil {
		xzap.ErrorContext(ctx, "", zap.Error(err))
		return nil, err
	}
	stats["times"], err = redisClient.HIncrBy(k, "times", 1).Result()
	if err != nil {
		xzap.ErrorContext(ctx, "", zap.Error(err))
		return nil, err
	}
	redisClient.Expire(k, 24*time.Hour)
	return stats, err
}
