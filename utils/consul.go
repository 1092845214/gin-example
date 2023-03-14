package utils

import (
	consul "github.com/hashicorp/consul/api"
	"github.com/yangkaiyue/gin-exp/global"
	"strings"
)

type ConsulCli struct {
	Client   *consul.KV
	URL      string
	User     string
	Password string
}

func NewConsulCli(url, user, password string) (syncer *ConsulCli, err error) {

	cli, err := consul.NewClient(&consul.Config{
		Address: url,
		HttpAuth: &consul.HttpBasicAuth{
			Username: user,
			Password: password,
		},
	})

	if err != nil {
		global.Logger.Error("Create Consul Client Failed. Err: ", err.Error())
		return
	}

	return &ConsulCli{
		URL:      url,
		User:     user,
		Password: password,
		Client:   cli.KV(),
	}, nil
}

// GetKey 获取 key 信息
func (sy *ConsulCli) GetKey(key string) (bool, []byte, error) {

	// 请求 kv 的值
	r, _, err := sy.Client.Get(key, nil)

	// 如果 err 不为空说明请求或链接失败
	if err != nil {
		global.Logger.Errorf("Get Consul Key %v Failed. Err: %v", key, err.Error())
		return false, nil, err
	}

	// 如果结果为空说明 key 不存在
	if r == nil {
		global.Logger.Infof("Consul Key %v Not Exist", key)
		return false, nil, nil
	}

	return true, r.Value, nil
}

// GetKeys 根据 pre 获取所有 key
func (sy *ConsulCli) GetKeys(prefix, sep string) (rs []string, err error) {

	keys, _, err := sy.Client.Keys(prefix, sep, nil)

	// 如果 err 不为空说明请求或链接失败
	if err != nil {
		global.Logger.Errorf("Get Consul Keys %v Failed. Err: %v", prefix, err.Error())
		return nil, err
	}

	// 如果结果为空说明 key 不存在
	if keys == nil {
		global.Logger.Infof("Consul Keys %v Not Exist", prefix)
		return nil, err
	}

	// 去头尾
	for i := 0; i < len(keys); i++ {
		tmp1 := strings.TrimPrefix(keys[i], prefix)
		tmp2 := strings.TrimSuffix(tmp1, sep)
		rs = append(rs, tmp2)
	}

	return
}

// PutKey 创建
func (sy *ConsulCli) PutKey(key string, value []byte) error {

	data := &consul.KVPair{
		Key:   key,
		Value: value,
	}

	if _, err := sy.Client.Put(data, nil); err != nil {
		global.Logger.Errorf("Put Consul Key %v Failed. Err: %v", key, err.Error())
		return err
	}

	return nil
}

// DelKey 删除
func (sy *ConsulCli) DelKey(key string) error {

	if _, err := sy.Client.Delete(key, nil); err != nil {
		global.Logger.Errorf("Delete Consul Key %v Failed. Err: %v", key, err.Error())
		return err
	}

	return nil
}
