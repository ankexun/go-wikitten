package main

import (
	"html/template"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/russross/blackfriday.v2"
)

//根据文件后缀名选择解析方式
//返回值: content, lang, error
func Parse(file string) (template.HTML, string, error) {
	switch Suffix(file) {
	case ".md":
		if input, err := ioutil.ReadFile(file); err == nil {

			unsafe := blackfriday.Run(input)
			html := template.HTML(string(bluemonday.UGCPolicy().SanitizeBytes(unsafe)))
			return html, "", nil
		} else {
			return "", "", err
		}
	case ".jpg", ".png", ".gif":
		html := template.HTML(`<img style="-webkit-user-select: none;cursor: zoom-out;" src="` + file + `">`)
		return html, "", nil
	case ".html", ".htm":
		if input, err := ioutil.ReadFile(file); err == nil {

			// unsafe := blackfriday.Run(input)
			html := template.HTML(string(bluemonday.UGCPolicy().SanitizeBytes(input)))
			return html, "", nil
		} else {
			return "", "", err
		}
	//使用prettify高亮代码
	case ".sh", ".css", ".js", ".py", ".rb", ".sql":
		if input, err := ioutil.ReadFile(file); err == nil {

			str := string(input[:])
			html := template.HTML(`<pre class="prettyprint linenums">` + str + `</pre>`)
			return html, "", nil
		} else {
			return "", "", err
		}
	//使用codemirror高亮代码
	case ".scm":
		if input, err := ioutil.ReadFile(file); err == nil {

			str := string(input[:])
			html := template.HTML(`<textarea id="myTextarea" class="form-control" rows="42">` + str + `</textarea>`)

			return html, "text/x-scheme", nil
		} else {
			return "", "", err
		}
	case ".php":
		if input, err := ioutil.ReadFile(file); err == nil {

			str := string(input[:])
			html := template.HTML(`<textarea id="myTextarea" class="form-control" rows="42">` + str + `</textarea>`)

			return html, "application/x-httpd-php", nil
		} else {
			return "", "", err
		}
	case ".xml":
		if input, err := ioutil.ReadFile(file); err == nil {

			str := string(input[:])
			html := template.HTML(`<textarea id="myTextarea" class="form-control" rows="42">` + str + `</textarea>`)

			return html, "application/xml", nil
		} else {
			return "", "", err
		}
	default:
		return "", "", nil
	}
}

//解析文件后缀名
func Suffix(file string) string {
	file = strings.ToLower(filepath.Base(file))
	if i := strings.LastIndex(file, "."); i > -1 {
		if file[i:] == "bz2" || file[i:] == ".gz" || file[i:] == ".xz" {
			if j := strings.LastIndex(file[:i], "."); j > -1 && strings.HasPrefix(file[j:], ".tar") {
				return file[j:]
			}
		}
		return file[i:]
	}
	return file
}
