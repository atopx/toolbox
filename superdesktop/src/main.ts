import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import './styles/app.css';
import App from './app.vue'
import 'highlight.js/styles/atom-one-light.css'
import 'highlight.js/lib/common'
import hljsVuePlugin from '@highlightjs/vue-plugin'
import router from './router'


const app = createApp(App);
app.use(ElementPlus);
app.use(hljsVuePlugin);
app.use(router);
app.mount('#app');
