package util

import (
	"github.com/gin-gonic/gin"
	"github.com/n1try/kithub2/app/config"
	"github.com/n1try/kithub2/app/util"
	"html/template"
	"strings"
)

func GetFuncMap() template.FuncMap {
	return template.FuncMap{
		"strIndex":    strIndex,
		"strRemove":   strRemove,
		"htmlSafe":    htmlSafe,
		"randomColor": util.RandomColor,
	}
}

type TplCtx struct {
	Path      string
	Constants struct {
		FacultyIndex int
		VvzBaseUrl   string
	}
	Alerts []string
	Errors []string
}

func GetTplCtx(c *gin.Context) TplCtx {
	var alerts []string
	if _, ok := c.Request.URL.Query()["postsignup"]; ok {
		alerts = []string{config.StrAlertSignupSuccessful}
	}

	return TplCtx{
		Path: c.FullPath(),
		Constants: struct {
			FacultyIndex int
			VvzBaseUrl   string
		}{
			FacultyIndex: config.FacultyIdx,
			VvzBaseUrl:   config.KitVvzBaseUrl,
		},
		Alerts: alerts,
	}
}

// Template funcs

func strIndex(x int, v string) string {
	return string([]rune(v)[:1])
}

func strRemove(html, needle string) string {
	return strings.ReplaceAll(html, needle, "")
}

func htmlSafe(html string) template.HTML {
	return template.HTML(html)
}
