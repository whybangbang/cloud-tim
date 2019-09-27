package cloud_tim

import (
	"context"
	"fmt"
	"testing"
)

var appId string

func init() {
	appId = "xxxx"
}

func TestNewCloudTimClient(t *testing.T) {
	timClient := NewCloudTimClient(appId, "")
	timClient.SendVoice(context.Background(), NewVoiceParam("ext", "18888888888", 1569565981, "测试消息")).Do()
}

func TestCloudTimClient_SendVoiceTpl(t *testing.T) {
	timClient := NewCloudTimClient(appId, "")
	resp, err := timClient.SendVoiceTpl(context.Background(), NewVoiceTplParam("ext", "18888888888", 1569565981, 12, "a")).Do()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}
