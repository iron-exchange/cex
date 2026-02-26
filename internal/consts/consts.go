package consts

const (
	// RedisLockPrefix 分布式锁通用前缀
	RedisLockPrefix = "CEX:LOCK:"

	// RedisAssetLockPrefix 资金防穿透专属锁前缀 (例如: CEX:LOCK:USER_WALLET:8888)
	RedisAssetLockPrefix = RedisLockPrefix + "USER_WALLET:"

	// LockWatchDogTimeout 锁的默认持有时间 (续期前的兜底时间)
	LockWatchDogTimeout = 8 * 1000 // 8 秒

	// LoginUserCachePrefix 登录用户缓存前缀
	LoginUserCachePrefix = "CEX:USER:LOGIN:"
)
