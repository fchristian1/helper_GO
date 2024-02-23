package helper_GO

import (
	"os"
)

type LangFile struct {
	LangName string            `json:"langName"`
	Messages map[string]string `json:"messages"`
}

func InitLang() LangFile {
	//read ./lang.json file
	b := FILE_LoadFileAsByte("./lang.json")

	//unmarshal to langJson
	l := JSON_UnmarshalToInterface[[]LangFile](b)
	//get langENV from enviroment directly its a import cycle problem
	langENV := os.Getenv("API_LANGUAGE")
	if langENV == "" {
		langENV = "en"
	}
	for _, v := range l {
		if v.LangName == langENV {
			return v
		}
	}
	panic("lang not found")
}

func (l *LangFile) GetTextFromLang(key string) string {
	return l.Messages[key]
}
