package cloud_tim

type BaseParam struct {
	Ext       string   `json:"ext,omitempty"`
	PlayTimes int      `json:"playtimes,omitempty"`
	Sig       string   `json:"sig"`
	Tel       TelParam `json:"tel"`
	Time      int64    `json:"time"`
}

func NewVoiceParam(ext string, mobile string, time int64, msg string) VoiceParam {
	return NewVoiceParamComplex(ext, 2, mobile, "86", time, msg)
}

func NewVoiceParamComplex(ext string, playtimes int, mobile string, nationcode string, time int64, msg string) VoiceParam {
	// playtimes add a limit
	return VoiceParam{
		BaseParam: BaseParam{
			Ext:       ext,
			PlayTimes: playtimes,
			Tel: TelParam{
				Mobile:     mobile,
				NationCode: nationcode,
			},
			Time: time,
		},
		Msg: msg,
	}
}

type VoiceParam struct {
	BaseParam
	Msg string `json:"msg"`
}

type TelParam struct {
	Mobile     string `json:"mobile"`
	NationCode string `json:"nationcode"`
}

func NewVoiceTplParam(ext string, mobile string, time int64, tplId int, params ...string) VoiceTplParam {
	return NewVoiceTplParamComplex(ext, 2, mobile, "86", time, tplId, params...)
}

func NewVoiceTplParamComplex(ext string, playtimes int, mobile string, nationcode string, time int64, tplId int, params ...string) VoiceTplParam {
	return VoiceTplParam{
		BaseParam: BaseParam{
			Ext:       ext,
			PlayTimes: playtimes,
			Tel: TelParam{
				Mobile:     mobile,
				NationCode: nationcode,
			},
			Time: time,
		},
		TplId:  tplId,
		Params: params,
	}
}

type VoiceTplParam struct {
	BaseParam
	TplId  int      `json:"tpl_id"`
	Params []string `json:"params"`
}

type VoiceResp struct {
	Result int    `json:"result"`
	RrrMsg string `json:"errmsg"`
	CallId string `json:"callid"`
	Ext    string `json:"ext,omitempty"`
}
