import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'

import App from './App.vue'
import Home from './content/HomePage.vue'
import Login from './user/LoginPage.vue'
import Register from './user/RegisterPage.vue'
import Question from './content/QuestionPage.vue'
import QuestionList from './content/QuestionListPage.vue'

const routes = [
    { path: '/', component: Home },
    { path: '/login', component: Login },
    { path: '/register', component: Register },
    { path: '/question/:id', component: Question },
    { path: '/questions', component: QuestionList },
    // { path: '/questions/all', component: QuestionList },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

const app = createApp(App)

app.use(router)

app.mount('#app')
