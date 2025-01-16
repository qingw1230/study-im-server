package db

import (
	"github.com/gomodule/redigo/redis"
	"github.com/qingw1230/study-im-server/pkg/base_info"
	"github.com/qingw1230/study-im-server/pkg/common/log"
)

const (
	redisTimeOneMintue = 60
	redisTimeOneDay    = 60 * 60 * 24

	checkCode         = "STUDYIM:CHECK_CODE:"
	tokenToUserInfo   = "STUDYIM:TOKEN_TO_USER_INFO:"
	userIDToToken     = "STUDYIM:USER_ID_TO_TOKEN:"
	userIDTokenStatus = "STUDYIM:USER_ID_TOKEN_STATUS:"
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

func (d *DataBases) SetUserToken(token base_info.TokenToRedis) error {
	key1 := tokenToUserInfo + token.Token
	_, err1 := d.Exec("SET", key1, token, "EX", redisTimeOneDay)
	if err1 != nil {
		return err1
	}
	key2 := userIDToToken + token.UserID
	_, err2 := d.Exec("SET", key2, token.Token, "EX", redisTimeOneDay)
	return err2
}

// AddTokenFlag 添加用户 token
func (d *DataBases) AddTokenFlag(userID string, token string, flag int) error {
	key := userIDTokenStatus + userID
	log.Debug("add token key is ", key)
	_, err := d.Exec("HSET", key, token, flag)
	return err
}

func (d *DataBases) GetTokenMapByUid(userID string) (map[string]int, error) {
	key := userIDTokenStatus + userID
	log.Debug("get token key is ", key)
	return redis.IntMap(d.Exec("HGETALL", key))
}

func (d *DataBases) DeleteTokenByUid(userID string, fields []string) error {
	key := userIDTokenStatus + userID
	_, err := d.Exec("HDEL", key, redis.Args{}.Add().AddFlat(fields)...)
	return err
}
