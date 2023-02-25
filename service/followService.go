package service

type FollowService interface {
	// IsFollowing 根据当前用户id和目标用户id来判断当前用户是否关注了目标用户
	IsFollowing(userId int64, targetId int64) (bool, error)
	// GetFollowerCnt 根据用户id来查询用户被多少其他用户关注
	GetFollowerCnt(userId int64) (int64, error)
	// GetFollowingCnt 根据用户id来查询用户关注了多少其它用户
	GetFollowingCnt(userId int64) (int64, error)
}
