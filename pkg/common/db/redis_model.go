package db

import (
	"github.com/gomodule/redigo/redis"
)

const (
	redisTimeOneMintue = 60
	redisTimeOneDay    = 60 * 60 * 24

	checkCode                     = "STUDYIM:CHECK_CODE:"
	userIdTokenStatus             = "STUDYIM:USER_ID_TOKEN_STATUS:"
	userIncrSeq                   = "STUDYIM:USER_INCR_SEQ:"
	conversationReceiveMessageOpt = "STUDYIM:CON_RECV_MSG_OPT:"
)

func (d *DataBases) Exec(cmd string, key interface{}, args ...interface{}) (interface{}, error) {
	conn := d.redisPool.Get()
	if err := conn.Err(); err != nil {
		return nil, err
	}
	defer conn.Close()

	params := make([]interface{}, 0)
	params = append(params, key)

	if len(args) > 0 {
		params = append(params, args...)
	}
	return conn.Do(cmd, params...)
}

func (d *DataBases) SetCheckCode(id, ans string) error {
	key := checkCode + id
	_, err := d.Exec("SET", key, ans, "EX", 5*redisTimeOneMintue)
	return err
}

func (d *DataBases) GetCheckCode(id string) (string, error) {
	key := checkCode + id
	// TODO(qingw1230): 将 GET 和 DEL 放在事务中
	reply, err := d.Exec("GET", key)
	// 获取后直接删除，防止多次尝试
	d.Exec("DEL", key)
	ans, _ := redis.String(reply, err)
	return ans, err
}

func (d *DataBases) IncrUserSeq(userId string) (uint64, error) {
	key := userIncrSeq + userId
	return redis.Uint64(d.Exec("INCR", key))
}

func (d *DataBases) GetUserMaxSeq(userId string) (uint64, error) {
	key := userIncrSeq + userId
	return redis.Uint64(d.Exec("GET", key))
}

// AddTokenFlag 添加用户 token
func (d *DataBases) AddTokenFlag(userId string, token string, flag int) error {
	key := userIdTokenStatus + userId
	_, err := d.Exec("HSET", key, token, flag)
	// TODO(qingw1230): 为 token 设置过期时间
	return err
}

func (d *DataBases) GetTokenMapByUid(userId string) (map[string]int, error) {
	key := userIdTokenStatus + userId
	return redis.IntMap(d.Exec("HGETALL", key))
}

func (d *DataBases) DeleteTokenByUid(userId string, fields []string) error {
	key := userIdTokenStatus + userId
	_, err := d.Exec("HDEL", key, redis.Args{}.Add().AddFlat(fields)...)
	return err
}

func (d *DataBases) SetSingleConversationMsgOpt(userId, conversationId string, opt int) error {
	key := conversationReceiveMessageOpt + userId
	_, err := d.Exec("HSET", key, conversationId, opt)
	return err
}

func (d *DataBases) GetSingleConversationMsgOpt(userId, conversationId string) (int, error) {
	key := conversationReceiveMessageOpt + userId
	return redis.Int(d.Exec("HGET", key, conversationId))
}
