package dao

import (
	"log"
	"sync"
)

type Follow struct {
	Id         int64
	UserId     int64
	FolloerwId int64
	Cancel     int8
}

type FollowDao struct {
}

func (Follow) TableName() string {
	return "follows"
}

var (
	followDao  *FollowDao //操作该dao层crud的结构体变量。
	followOnce sync.Once  //单例限定，去限定申请一个followDao结构体变量。
)

func NewFollowDaoInstance() *FollowDao {
	followOnce.Do(
		func() {
			followDao = &FollowDao{}
		})
	return followDao
}

func (*FollowDao) GetFollowerCnt(userId int64) (int64, error) {
	var cnt int64
	if err := Db.Model(Follow{}).
		Where("follower_id = ?", userId).
		Where("cancel = ?", 0).
		Count(&cnt).Error; nil != err {
		log.Println(err.Error())
		return 0, err
	}
	return cnt, nil
}

func (*FollowDao) GetFollowingCnt(userId int64) (int64, error) {
	// 用于存储当前用户关注了多少人。
	var cnt int64
	// 查询出错，日志打印err msg，并return err
	if err := Db.Model(Follow{}).
		Where("follower_id = ?", userId).
		Where("cancel = ?", 0).
		Count(&cnt).Error; nil != err {
		log.Println(err.Error())
		return 0, err
	}
	// 查询成功，返回人数。
	return cnt, nil
}

func (*FollowDao) GetFollowingIds(userId int64) ([]int64, error) {
	var ids []int64
	if err := Db.
		Model(Follow{}).
		Where("follower_id = ?", userId).
		Pluck("user_id", &ids).Error; nil != err {
		// 没有关注任何人，但是不能算错。
		if "record not found" == err.Error() {
			return nil, nil
		}
		// 查询出错。
		log.Println(err.Error())
		return nil, err
	}
	// 查询成功。
	return ids, nil
}

// GetFollowersIds 给定用户id，查询他关注了哪些人的id。
func (*FollowDao) GetFollowersIds(userId int64) ([]int64, error) {
	var ids []int64
	if err := Db.
		Model(Follow{}).
		Where("user_id = ?", userId).
		Where("cancel = ?", 0).
		Pluck("follower_id", &ids).Error; nil != err {
		// 没有粉丝，但是不能算错。
		if "record not found" == err.Error() {
			return nil, nil
		}
		// 查询出错。
		log.Println(err.Error())
		return nil, err
	}
	// 查询成功。
	return ids, nil
}
func (*FollowDao) FindRelation(userId int64, targetId int64) (*Follow, error) {
	// follow变量用于后续存储数据库查出来的用户关系。
	follow := Follow{}
	//当查询出现错误时，日志打印err msg，并return err.
	if err := Db.
		Where("user_id = ?", targetId).
		Where("follower_id = ?", userId).
		Where("cancel = ?", 0).
		Take(&follow).Error; nil != err {
		// 当没查到数据时，gorm也会报错。
		if "record not found" == err.Error() {
			return nil, nil
		}
		log.Println(err.Error())
		return nil, err
	}
	//正常情况，返回取到的值和空err.
	return &follow, nil
}
