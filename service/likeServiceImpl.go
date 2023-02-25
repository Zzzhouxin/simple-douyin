package service

type LikeServiceImpl struct {
	VideoService
	UserService
}

// GetLikeService 解决likeService调videoService,videoService调userService,useService调likeService循环依赖的问题
func GetLikeService() LikeServiceImpl {
	var userService UserServiceImpl
	var videoService VideoServiceImpl
	var likeService LikeServiceImpl
	userService.LikeService = &likeService
	likeService.VideoService = &videoService
	videoService.UserService = &userService
	return likeService
}
