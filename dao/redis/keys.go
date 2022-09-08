package redis

// redis key
// redis key注意使用命名空间的方式，区分不同的key，方便业务查询、拆分

// 这些表示key
const (
	KeyPrefix          = "webProject:"
	KeyPostTimeZSet    = "post:time"   // zset：帖子及发帖时间为分数 ，value是时间+帖子的id，这个key是用来排序的
	KeyPostScoreZSet   = "post:score"  // zset：帖子及投票的分数
	KeyPostVotedZSetPF = "post:voted:" // zset:记录用户以及投票类型，参数是post id
	KeyCommunitySetPF  = "community:"  // set:保存每个分区下帖子的ID
)

// 给Redis key加上前缀
func getRedisKey(key string) string {
	return KeyPrefix + key
}
