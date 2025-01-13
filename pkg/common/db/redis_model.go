package db

import "github.com/gomodule/redigo/redis"

const (
	checkCode = "STUDYIM:CHECK_CODE:"
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
	_, err := d.Exec("SET", key, ans, "EX", 5*60)
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
