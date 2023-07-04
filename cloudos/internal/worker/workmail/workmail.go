package workmail

import (
	"cloudos/common/system"
	"cloudos/common/utils"
	"cloudos/internal/model"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"gopkg.in/gomail.v2"
)

// Scheduler 定时发工作邮件, 解决两个问题：
// 1. 用markdown写邮件问题
// 2. 定时发送问题(可以提前写、汇总，省的到写的时候话费大量时间思考)
func Scheduler(flag string) error {
	flags := strings.SplitN(flag, "/", 2)
	if len(flags) != 2 {
		return errors.New("invalid params, run --help cat examples")
	}
	noteId, err := strconv.ParseInt(flags[0], 10, 64)
	if err != nil {
		return errors.New("invalid noteId params, run --help cat examples")
	}
	t, err := time.ParseInLocation(time.DateTime, flags[1], time.Local)
	if err != nil {
		return errors.New("invalid worktime params, run --help cat examples")
	}
	sleeps := t.Unix() - time.Now().Unix()
	if sleeps < 0 {
		return errors.New("worktime must greater than current time")
	}
	note := new(model.NoteDao).First("id = ?", noteId)
	if note == nil {
		return fmt.Errorf("not found this note by %d", noteId)
	}
	// IDEA 有时间可以补充一个收邮件并且校验任务是否成功的逻辑
	log.Printf("sleep %ds... work on [%s] send mail <%s>\n", sleeps, t.Format(time.DateTime), note.Title)
	time.Sleep(time.Duration(sleeps) * time.Second)
	handler := system.GetHandler()
	message := gomail.NewMessage()
	if handler.Config.Workmail.To != "" {
		message.SetHeader("To", handler.Config.Workmail.To)
		log.Printf("wkmail to: %s\n", handler.Config.Workmail.To)
	}

	if handler.Config.Workmail.Cc != "" {
		message.SetHeader("Cc", handler.Config.Workmail.Cc)
		log.Printf("wkmail cc: %s\n", handler.Config.Workmail.Cc)
	}

	message.SetAddressHeader("From", handler.Config.Workmail.Email, handler.Config.Workmail.Nickname)
	message.SetHeader("Subject", note.Title)
	message.SetBody("text/html", render(note.Content))
	return gomail.NewDialer(
		handler.Config.Workmail.Server,
		handler.Config.Workmail.Port,
		handler.Config.Workmail.Email,
		handler.Config.Workmail.Password,
	).DialAndSend(message)
}

func render(mark string) string {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock | parser.NoEmptyLineBeforeBlock
	doc := parser.NewWithExtensions(extensions).Parse(utils.Bytes(mark))
	opts := html.RendererOptions{Flags: html.CommonFlags | html.HrefTargetBlank}
	content := utils.String(markdown.Render(doc, html.NewRenderer(opts)))
	return strings.Replace(template, "{{content}}", content, 1)
}

const template = `<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <title>Document</title>
    </head>
    <body>
        <div class="vuepress-markdown-body">{{content}}</div>
    </body>

    <style>
        .vuepress-markdown-body code[class*="v-md-prism-"],
        .vuepress-markdown-body pre[class*="v-md-prism-"] {
            color: #ccc;
            font-size: 1em;
            font-family: Consolas, Monaco, Andale Mono, Ubuntu Mono, monospace;
            line-height: 1.5;
            white-space: pre;
            text-align: left;
            word-wrap: normal;
            word-break: normal;
            word-spacing: normal;
            -webkit-hyphens: none;
            -ms-hyphens: none;
            hyphens: none;
            background: none;
        }
        .vuepress-markdown-body > :first-child,
        .vuepress-markdown-body
            > div[data-v-md-line]:first-child
            > :first-child {
            margin-top: 0 !important;
        }
        .vuepress-markdown-body > :last-child,
        .vuepress-markdown-body > div[data-v-md-line]:last-child > :last-child {
            margin-bottom: 0 !important;
        }
        .vuepress-markdown-body pre[class*="v-md-prism-"] {
            margin: 0.5em 0;
            padding: 1em;
            overflow: auto;
        }
        .vuepress-markdown-body :not(pre) > code[class*="v-md-prism-"],
        .vuepress-markdown-body pre[class*="v-md-prism-"] {
            background: #2d2d2d;
        }
        .vuepress-markdown-body :not(pre) > code[class*="v-md-prism-"] {
            padding: 0.1em;
            white-space: normal;
            border-radius: 0.3em;
        }
        .vuepress-markdown-body .token.block-comment,
        .vuepress-markdown-body .token.cdata,
        .vuepress-markdown-body .token.comment,
        .vuepress-markdown-body .token.doctype,
        .vuepress-markdown-body .token.prolog {
            color: #999;
        }
        .vuepress-markdown-body .token.punctuation {
            color: #ccc;
        }
        .vuepress-markdown-body .token.attr-name,
        .vuepress-markdown-body .token.deleted,
        .vuepress-markdown-body .token.namespace,
        .vuepress-markdown-body .token.tag {
            color: #e2777a;
        }
        .vuepress-markdown-body .token.function-name {
            color: #6196cc;
        }
        .vuepress-markdown-body .token.boolean,
        .vuepress-markdown-body .token.function,
        .vuepress-markdown-body .token.number {
            color: #f08d49;
        }
        .vuepress-markdown-body .token.class-name,
        .vuepress-markdown-body .token.constant,
        .vuepress-markdown-body .token.property,
        .vuepress-markdown-body .token.symbol {
            color: #f8c555;
        }
        .vuepress-markdown-body .token.atrule,
        .vuepress-markdown-body .token.builtin,
        .vuepress-markdown-body .token.important,
        .vuepress-markdown-body .token.keyword,
        .vuepress-markdown-body .token.selector {
            color: #cc99cd;
        }
        .vuepress-markdown-body .token.attr-value,
        .vuepress-markdown-body .token.char,
        .vuepress-markdown-body .token.regex,
        .vuepress-markdown-body .token.string,
        .vuepress-markdown-body .token.variable {
            color: #7ec699;
        }
        .vuepress-markdown-body .token.entity,
        .vuepress-markdown-body .token.operator,
        .vuepress-markdown-body .token.url {
            color: #67cdcc;
        }
        .vuepress-markdown-body .token.bold,
        .vuepress-markdown-body .token.important {
            font-weight: 700;
        }
        .vuepress-markdown-body .token.italic {
            font-style: italic;
        }
        .vuepress-markdown-body .token.entity {
            cursor: help;
        }
        .vuepress-markdown-body .token.inserted {
            color: green;
        }
        .vuepress-markdown-body code {
            margin: 0;
            padding: 0.25rem 0.5rem;
            color: #476582;
            font-size: 0.85em;
            background-color: rgba(27, 31, 35, 0.05);
            border-radius: 3px;
        }
        .vuepress-markdown-body code .token.deleted {
            color: #ec5975;
        }
        .vuepress-markdown-body code .token.inserted {
            color: #3eaf7c;
        }
        .vuepress-markdown-body pre,
        .vuepress-markdown-body pre[class*="v-md-prism-"] {
            margin: 0.85rem 0;
            padding: 1.25rem 1.5rem;
            overflow: auto;
            line-height: 1.4;
            background-color: #282c34;
            border-radius: 6px;
        }
        .vuepress-markdown-body pre[class*="v-md-prism-"] code,
        .vuepress-markdown-body pre code {
            padding: 0;
            color: #fff;
            background-color: transparent;
            border-radius: 0;
        }
        .vuepress-markdown-body div[class*="v-md-pre-wrapper-"] {
            position: relative;
            background-color: #282c34;
            border-radius: 6px;
        }
        .vuepress-markdown-body div[class*="v-md-pre-wrapper-"] pre,
        .vuepress-markdown-body
            div[class*="v-md-pre-wrapper-"]
            pre[class*="v-md-prism-"] {
            position: relative;
            z-index: 1;
            background: transparent;
        }
        .vuepress-markdown-body div[class*="v-md-pre-wrapper-"]::before {
            position: absolute;
            top: 0.8em;
            right: 1em;
            z-index: 3;
            color: hsla(0, 0%, 100%, 0.4);
            font-size: 0.75rem;
        }
        .vuepress-markdown-body
            div[class*="v-md-pre-wrapper-"]:not(.line-numbers-mode)
            .line-numbers-wrapper {
            display: none;
        }
        .vuepress-markdown-body
            div[class*="v-md-pre-wrapper-"].line-numbers-mode
            pre {
            padding-left: 4.5rem;
            vertical-align: middle;
        }
        .vuepress-markdown-body
            div[class*="v-md-pre-wrapper-"].line-numbers-mode
            .line-numbers-wrapper {
            position: absolute;
            top: 0;
            width: 3.5rem;
            padding: 1.25rem 0;
            color: hsla(0, 0%, 100%, 0.3);
            line-height: 1.4;
            text-align: center;
        }
        .vuepress-markdown-body
            div[class*="v-md-pre-wrapper-"].line-numbers-mode
            .line-numbers-wrapper
            br {
            -webkit-user-select: none;
            user-select: none;
        }
        .vuepress-markdown-body
            div[class*="v-md-pre-wrapper-"].line-numbers-mode
            .line-numbers-wrapper
            .line-number {
            position: relative;
            z-index: 4;
            font-size: 0.85em;
            -webkit-user-select: none;
            user-select: none;
        }
        .vuepress-markdown-body
            div[class*="v-md-pre-wrapper-"].line-numbers-mode::after {
            position: absolute;
            top: 0;
            left: 0;
            z-index: 2;
            width: 3.5rem;
            height: 100%;
            background-color: #282c34;
            border-right: 1px solid rgba(0, 0, 0, 0.66);
            border-radius: 6px 0 0 6px;
            content: "";
        }
        .vuepress-markdown-body div[class~="v-md-pre-wrapper-js"]::before {
            content: "js";
        }
        .vuepress-markdown-body div[class~="v-md-pre-wrapper-ts"]::before {
            content: "ts";
        }
        .vuepress-markdown-body div[class~="v-md-pre-wrapper-html"]::before {
            content: "html";
        }
        .vuepress-markdown-body div[class~="v-md-pre-wrapper-md"]::before {
            content: "md";
        }
        .vuepress-markdown-body div[class~="v-md-pre-wrapper-vue"]::before {
            content: "vue";
        }
        .vuepress-markdown-body div[class~="v-md-pre-wrapper-css"]::before {
            content: "css";
        }
        .vuepress-markdown-body div[class~="v-md-pre-wrapper-sass"]::before {
            content: "sass";
        }
        .vuepress-markdown-body div[class~="v-md-pre-wrapper-scss"]::before {
            content: "scss";
        }
        .vuepress-markdown-body div[class~="v-md-pre-wrapper-less"]::before {
            content: "less";
        }
        .vuepress-markdown-body div[class~="v-md-pre-wrapper-stylus"]::before {
            content: "stylus";
        }
        .vuepress-markdown-body div[class~="v-md-pre-wrapper-go"]::before {
            content: "go";
        }
        .vuepress-markdown-body div[class~="v-md-pre-wrapper-java"]::before {
            content: "java";
        }
        .vuepress-markdown-body div[class~="v-md-pre-wrapper-c"]::before {
            content: "c";
        }
        .vuepress-markdown-body div[class~="v-md-pre-wrapper-sh"]::before {
            content: "sh";
        }
        .vuepress-markdown-body div[class~="v-md-pre-wrapper-yaml"]::before {
            content: "yaml";
        }
        .vuepress-markdown-body div[class~="v-md-pre-wrapper-py"]::before {
            content: "py";
        }
        .vuepress-markdown-body div[class~="v-md-pre-wrapper-docker"]::before {
            content: "docker";
        }
        .vuepress-markdown-body
            div[class~="v-md-pre-wrapper-dockerfile"]::before {
            content: "dockerfile";
        }
        .vuepress-markdown-body
            div[class~="v-md-pre-wrapper-makefile"]::before {
            content: "makefile";
        }
        .vuepress-markdown-body
            div[class~="v-md-pre-wrapper-javascript"]::before {
            content: "js";
        }
        .vuepress-markdown-body
            div[class~="v-md-pre-wrapper-typescript"]::before {
            content: "ts";
        }
        .vuepress-markdown-body div[class~="v-md-pre-wrapper-markup"]::before {
            content: "html";
        }
        .vuepress-markdown-body
            div[class~="v-md-pre-wrapper-markdown"]::before {
            content: "md";
        }
        .vuepress-markdown-body div[class~="v-md-pre-wrapper-json"]::before {
            content: "json";
        }
        .vuepress-markdown-body div[class~="v-md-pre-wrapper-ruby"]::before {
            content: "rb";
        }
        .vuepress-markdown-body div[class~="v-md-pre-wrapper-python"]::before {
            content: "py";
        }
        .vuepress-markdown-body div[class~="v-md-pre-wrapper-bash"]::before {
            content: "sh";
        }
        .vuepress-markdown-body div[class~="v-md-pre-wrapper-php"]::before {
            content: "php";
        }
        .vuepress-markdown-body .arrow {
            display: inline-block;
            width: 0;
            height: 0;
        }
        .vuepress-markdown-body .arrow.up {
            border-bottom: 6px solid #ccc;
        }
        .vuepress-markdown-body .arrow.down,
        .vuepress-markdown-body .arrow.up {
            border-right: 4px solid transparent;
            border-left: 4px solid transparent;
        }
        .vuepress-markdown-body .arrow.down {
            border-top: 6px solid #ccc;
        }
        .vuepress-markdown-body .arrow.right {
            border-left: 6px solid #ccc;
        }
        .vuepress-markdown-body .arrow.left,
        .vuepress-markdown-body .arrow.right {
            border-top: 4px solid transparent;
            border-bottom: 4px solid transparent;
        }
        .vuepress-markdown-body .arrow.left {
            border-right: 6px solid #ccc;
        }
        .vuepress-markdown-body:not(.custom) {
            padding: 2rem 2.5rem;
        }
        @media (max-width: 959px) {
            .vuepress-markdown-body:not(.custom) {
                padding: 2rem;
            }
        }
        @media (max-width: 419px) {
            .vuepress-markdown-body:not(.custom) {
                padding: 1.5rem;
            }
        }
        .vuepress-markdown-body .table-of-contents .badge {
            vertical-align: middle;
        }
        .vuepress-markdown-body {
            color: #2c3e50;
            font-size: 16px;
            font-family: -apple-system, BlinkMacSystemFont, Segoe UI, Roboto,
                Oxygen, Ubuntu, Cantarell, Fira Sans, Droid Sans, Helvetica Neue,
                sans-serif;
            -webkit-font-smoothing: antialiased;
            -moz-osx-font-smoothing: grayscale;
            background-color: #fff;
        }
        .vuepress-markdown-body:not(.custom) a:hover {
            text-decoration: underline;
        }
        .vuepress-markdown-body:not(.custom) p.demo {
            padding: 1rem 1.5rem;
            border: 1px solid #ddd;
            border-radius: 4px;
        }
        .vuepress-markdown-body:not(.custom) img {
            max-width: 100%;
        }
        .vuepress-markdown-body.custom {
            margin: 0;
            padding: 0;
        }
        .vuepress-markdown-body.custom img {
            max-width: 100%;
        }
        .vuepress-markdown-body a {
            font-weight: 500;
            text-decoration: none;
        }
        .vuepress-markdown-body a,
        .vuepress-markdown-body p a code {
            color: #3eaf7c;
        }
        .vuepress-markdown-body p a code {
            font-weight: 400;
        }
        .vuepress-markdown-body kbd {
            padding: 0 0.15em;
            background: #eee;
            border: 0.15rem solid #ddd;
            border-bottom: 0.25rem solid #ddd;
            border-radius: 0.15rem;
        }
        .vuepress-markdown-body blockquote {
            margin: 1rem 0;
            padding: 0.25rem 0 0.25rem 1rem;
            color: #1f39af;
            font-size: 1rem;
            border-left: 0.2rem solid #dfe2e5;
        }
        .vuepress-markdown-body blockquote > p {
            margin: 0;
        }
        .vuepress-markdown-body ol,
        .vuepress-markdown-body ul {
            margin: 1em 0;
            padding-left: 1.2em;
        }
        .vuepress-markdown-body strong {
            font-weight: 600;
        }
        .vuepress-markdown-body h1,
        .vuepress-markdown-body h2,
        .vuepress-markdown-body h3,
        .vuepress-markdown-body h4,
        .vuepress-markdown-body h5,
        .vuepress-markdown-body h6 {
            font-weight: 600;
            line-height: 1.25;
        }
        .vuepress-markdown-body h1 {
            margin: 0.67em 0;
            font-size: 2.2rem;
        }
        .vuepress-markdown-body h2 {
            margin: 0.83em 0;
            padding-bottom: 0.3rem;
            font-size: 1.65rem;
            border-bottom: 1px solid #eaecef;
        }
        .vuepress-markdown-body h3 {
            margin: 1em 0;
            font-size: 1.35rem;
        }
        .vuepress-markdown-body h4 {
            margin: 1.33em 0;
        }
        .vuepress-markdown-body h5 {
            margin: 1.67em 0;
        }
        .vuepress-markdown-body h6 {
            margin: 2.33em 0;
        }
        .vuepress-markdown-body i,
        .vuepress-markdown-body em {
            font-style: italic;
        }
        .vuepress-markdown-body ul {
            list-style-type: disc;
        }
        .vuepress-markdown-body ul ul,
        .vuepress-markdown-body ol ul {
            list-style-type: circle;
        }
        .vuepress-markdown-body ol ol ul,
        .vuepress-markdown-body ol ul ul,
        .vuepress-markdown-body ul ol ul,
        .vuepress-markdown-body ul ul ul {
            list-style-type: square;
        }
        .vuepress-markdown-body ol {
            list-style-type: decimal;
        }
        .vuepress-markdown-body .line-number,
        .vuepress-markdown-body code,
        .vuepress-markdown-body kbd {
            font-family: source-code-pro, Menlo, Monaco, Consolas, Courier New,
                monospace;
        }
        .vuepress-markdown-body ol,
        .vuepress-markdown-body p,
        .vuepress-markdown-body ul {
            line-height: 1.7;
        }
        .vuepress-markdown-body hr {
            border: 0;
            border-top: 1px solid #eaecef;
        }
        .vuepress-markdown-body table {
            display: block;
            margin: 1rem 0;
            overflow-x: auto;
            border-collapse: collapse;
        }
        .vuepress-markdown-body tr {
            border-top: 1px solid #dfe2e5;
        }
        .vuepress-markdown-body tr:nth-child(2n) {
            background-color: #f6f8fa;
        }
        .vuepress-markdown-body td,
        .vuepress-markdown-body th {
            padding: 0.6em 1em;
            border: 1px solid #dfe2e5;
        }
        .vuepress-markdown-body .v-md-svg-outbound {
            position: relative;
            top: -1px;
            display: inline-block;
            color: #aaa;
            vertical-align: middle;
        }
        @media (max-width: 419px) {
            .vuepress-markdown-body h1 {
                font-size: 1.9rem;
            }
            .vuepress-markdown-body div[class*="v-md-pre-wrapper-"] {
                margin: 0.85rem -1.5rem;
                border-radius: 0;
            }
        }
    </style>
</html>
`
