package redis

import (
	"errors"
	"math"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432 // 每一票值多少分
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrVoteRepeated   = errors.New("不允许重复投票")
)

// CreatePost 创建帖子时，要对Redis进行的操作
func CreatePost(postID, communityID int64) error {
	pipeline := client.TxPipeline()
	// 帖子时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()), // 用时间戳表示帖子的分数
		Member: postID,                     // member存储帖子的id
	})

	// 帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()), //
		Member: postID,
	})
	// 更新：把帖子id加到社区的set
	cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(communityID)))
	pipeline.SAdd(cKey, postID)
	_, err := pipeline.Exec()
	return err
}

// VoteForPost 为帖子投票
func VoteForPost(userId, postId string, value float64) error {
	// 判断投票限制
	// 从Redis中获取帖子发布时间
	postTime := client.ZScore(getRedisKey(KeyPostTimeZSet), postId).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	//更新帖子的分数
	//查询之前的投票记录
	ov := client.ZScore(getRedisKey(KeyPostVotedZSetPF+postId), userId).Val()

	// 更新：如果这一次投票的值和之前投票的值一样，就提示不允许重复投票
	if value == ov {
		return ErrVoteRepeated
	}
	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value) // 计算两次投票的差值
	pipeline := client.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postId)

	// 3. 记录用户为该贴子投票的数据
	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZSetPF+postId), userId)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZSetPF+postId), redis.Z{
			Score:  value, // 赞成票还是反对票
			Member: userId,
		})
	}
	_, err := pipeline.Exec()
	return err
}
