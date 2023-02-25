package service

import (
	"github.com/RaymondCode/simple-douyin/dao"
	"sync"
)

type FollowServiceImp struct {
	UserService
}

// GetFollowerCnt 给定当前用户id，查询其粉丝数量。
func (*FollowServiceImp) GetFollowerCnt(userId int64) (int64, error) {
	// SQL中查询。
	ids, err := dao.NewFollowDaoInstance().GetFollowersIds(userId)
	if nil != err {
		return 0, err
	}
	return int64(len(ids)), err
}
func (*FollowServiceImp) GetFollowingCnt(userId int64) (int64, error) {
	// 用SQL查询。
	ids, err := dao.NewFollowDaoInstance().GetFollowersIds(userId)

	if nil != err {
		return 0, err
	}
	return int64(len(ids)), err
}

// GetFollowers 根据当前用户id来查询他的粉丝列表。
func (f *FollowServiceImp) getFollowers(userId int64) ([]User, error) {
	// 获取粉丝的id数组。
	ids, err := dao.NewFollowDaoInstance().GetFollowersIds(userId)
	// 查询出错
	if nil != err {
		return nil, err
	}
	// 没得粉丝
	if nil == ids {
		return nil, nil
	}
	// 根据每个id来查询用户信息。
	return f.getUserById(ids, userId)
}

func (f *FollowServiceImp) getUserById(ids []int64, userId int64) ([]User, error) {
	len := len(ids)
	if len > 0 {
		len -= 1
	}
	users := make([]User, len)
	var wg sync.WaitGroup
	wg.Add(len)
	i, j := 0, 0
	for ; i < len; j++ {
		// 越过-1
		if ids[j] == -1 {
			continue
		}
		//开启协程来查。
		go func(i int, idx int64) {
			defer wg.Done()
			users[i], _ = f.GetUserByIdWithCurId(idx, userId)
		}(i, ids[i])
		i++
	}
	wg.Wait()
	// 返回粉丝列表。
	return users, nil
}

func (*FollowServiceImp) IsFollowing(userId int64, targetId int64) (bool, error) {

	relation, err := dao.NewFollowDaoInstance().FindRelation(userId, targetId)

	if nil != err {
		return false, err
	}
	if nil == relation {
		return false, nil
	}

	return true, nil
}
