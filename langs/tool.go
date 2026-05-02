package langs

import (
	"bootpkg/pkg/core/modules/enmus"
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"path"
	"strings"
)

//go:embed locales
var locales embed.FS

var l *Localizer

func init() {
	dirEntries, _ := locales.ReadDir("locales")
	localizations := make(map[string]string, 0)
	for _, de := range dirEntries {
		if !de.IsDir() && strings.Contains(de.Name(), ".yaml") {
			vi := viper.New()
			vi.SetConfigType("yaml")
			fs, err := locales.Open("locales/" + de.Name())
			if err != nil {
				log.Fatalf("read go:embed locales failed:%v  %v", de.Name(), err)
			}
			if err := vi.ReadConfig(fs); err != nil {
				log.Fatalf("read config failed:%v  %v", de.Name(), err)
			}
			filenameall := path.Base(de.Name())
			filesuffix := path.Ext(de.Name())
			fileprefix := filenameall[0 : len(filenameall)-len(filesuffix)]
			for k, v := range vi.AllSettings() {
				localizations[fileprefix+"."+k] = fmt.Sprintf("%v", v)
			}
		}
	}
	l = &Localizer{Localizations: localizations}
}

// GetWithLocale
//
//	@Description: 获取多语言配置
//	@param locale 语言简码
//	@param key 语言标识
//	@param replacements 语言替换
//	@return string 国际化语言
func GetWithLocale(locale, key string, replacements ...*Replacements) string {
	return l.GetWithLocale(locale, key, replacements...)
}

// GetWithLocaleGin
//
//	@Description: 获取多语言配置
//	@param c 语言简码
//	@param key 语言标识
//	@param replacements 语言替换
//	@return string 国际化语言
func GetWithLocaleGin(c *gin.Context, key string, replacements ...*Replacements) string {
	locale := c.GetHeader(enmus.LANGUAGE_HEADER)
	return l.GetWithLocale(locale, key, replacements...)
}
