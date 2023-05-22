// core
import { createApp } from "vue"
import App from "@/App.vue"
import store from "@/store"
import router from "@/router"
import "@/router/permission"
// load
import { loadSvg } from "@/icons"
import { loadPlugins } from "@/plugins"
// css
import "uno.css"
import "normalize.css"
import "element-plus/dist/index.css"
import "element-plus/theme-chalk/dark/css-vars.css"
import "vxe-table/lib/style.css"
import "vxe-table-plugin-element/dist/style.css"
import "@/styles/index.scss"

import VMdEditor from "@kangc/v-md-editor/lib/codemirror-editor"
import VMdPreview from "@kangc/v-md-editor/lib/preview"
import "@kangc/v-md-editor/lib/style/codemirror-editor.css"
import "@kangc/v-md-editor/lib/style/preview.css"
import vuepressTheme from "@kangc/v-md-editor/lib/theme/vuepress.js"
import "@kangc/v-md-editor/lib/theme/style/vuepress.css"
import Prism from "prismjs"
// codemirror 编辑器的相关资源
import Codemirror from "codemirror"
// mode
import "codemirror/mode/markdown/markdown"
import "codemirror/mode/javascript/javascript"
import "codemirror/mode/css/css"
import "codemirror/mode/htmlmixed/htmlmixed"
import "codemirror/mode/vue/vue"
// edit
import "codemirror/addon/edit/closebrackets"
import "codemirror/addon/edit/closetag"
import "codemirror/addon/edit/matchbrackets"
// placeholder
import "codemirror/addon/display/placeholder"
// active-line
import "codemirror/addon/selection/active-line"
// scrollbar
import "codemirror/addon/scroll/simplescrollbars"
import "codemirror/addon/scroll/simplescrollbars.css"
// style
import "codemirror/lib/codemirror.css"

VMdEditor.Codemirror = Codemirror
VMdEditor.use(vuepressTheme, { Prism })
VMdPreview.use(vuepressTheme, { Prism })
const app = createApp(App)

/** 加载插件 */
loadPlugins(app)
/** 加载全局 SVG */
loadSvg(app)

app.use(VMdEditor)
app.use(VMdPreview)
app.use(store)
app.use(router)
app.mount("#app")
