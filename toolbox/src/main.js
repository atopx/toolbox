import { createApp } from "vue";
import App from "./App.vue";
import Layui from '@layui/layui-vue'
import '@layui/layui-vue/lib/index.css'
import './style.css'

const app = createApp(App)
app.use(Layui)
app.mount("#app");
