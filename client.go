package cloud_tim

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

const (
	basePath = "https://cloud.tim.qq.com/"
	POST     = "POST"
)

var httpServer http.Client

func init() {
	httpServer = http.Client{}
}

func NewCloudTimClient(appId string, appKey string) *CloudTimClient {
	return &CloudTimClient{
		appId:  appId,
		appKey: appKey,
	}
}

type CloudTimClient struct {
	appId  string
	appKey string
}

func (ctc *CloudTimClient) SendVoice(ctx context.Context, voice VoiceParam) *RequestCaller {
	randInt := random()
	voice.Sig = buildSig(ctc.appKey, voice.BaseParam.Tel.Mobile, strconv.FormatInt(randInt, 10), strconv.FormatInt(voice.Time, 10))
	rc := newRequestCaller(ctx, POST, "v5/tlsvoicesvr/sendcvoice", ctc.appId, randInt)
	rc.marshalBody(voice)
	return rc
}

func (ctc *CloudTimClient) SendVoiceTpl(ctx context.Context, voiceTpl VoiceTplParam) *RequestCaller {
	randInt := random()
	voiceTpl.Time = randInt
	voiceTpl.Sig = buildSig(ctc.appKey, voiceTpl.BaseParam.Tel.Mobile, strconv.FormatInt(randInt, 10), strconv.FormatInt(voiceTpl.Time, 10))
	rc := newRequestCaller(ctx, POST, strconv.FormatInt(randInt, 10), ctc.appId, randInt)
	rc.marshalBody(voiceTpl)
	return rc
}

func newRequestCaller(ctx context.Context, method string, relativePath string, appId string, randInt int64) *RequestCaller {

	uuu, _ := url.Parse(basePath)
	finalPath := path.Join(uuu.Path, relativePath)
	uuu.Path = finalPath
	query := url.Values{}
	query.Set("sdkappid", appId)
	query.Set("random", strconv.FormatInt(randInt, 10))
	uuu.RawQuery = query.Encode()
	return &RequestCaller{
		server: httpServer,
		ctx:    ctx,
		path:   finalPath,
		method: method,
		url:    uuu,
	}
}

type RequestCaller struct {
	server http.Client
	ctx    context.Context
	body   []byte
	path   string

	method string
	url    *url.URL
}

func (rc *RequestCaller) marshalBody(data interface{}) error {
	result, err := json.Marshal(data)
	if err != nil {
		return err
	}
	rc.body = result
	return nil
}

func (rc *RequestCaller) Do() (*VoiceResp, error) {
	request := &http.Request{}
	request = request.WithContext(rc.ctx)
	request.Method = rc.method
	request.URL = rc.url
	request.Body = ioutil.NopCloser(bytes.NewReader(rc.body))

	fmt.Println("method ", request.Method)
	fmt.Println("url ", request.URL)
	fmt.Println("params ", request.URL.RawQuery)
	body, err := ioutil.ReadAll(request.Body)
	fmt.Println("body ", string(body))

	resp, err := rc.server.Do(request)
	fmt.Println(resp.Body)
	if err != nil {
		return nil, err
	}

	result := &VoiceResp{}
	json.NewDecoder(resp.Body).Decode(result)
	return result, nil
}
