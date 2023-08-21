import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap/dist/js/bootstrap'

import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css';
import * as Icons from '@element-plus/icons-vue'


const app = createApp(App);
for (let i in Icons) {
    app.component(i, Icons[i])
}
app.use(store).use(router).use(ElementPlus).mount('#app')
