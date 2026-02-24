import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'
import './assets/main.css'
import { intersectDirective } from './directives/intersect'

const app = createApp(App)

app.directive('intersect', intersectDirective)

app.use(createPinia())
app.use(router)

app.mount('#app')
