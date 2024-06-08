import { createRouter, createWebHistory } from 'vue-router'
const HomeView = () => import('@/view/HomeView.vue')
const LoginView = () => import('@/view/LoginView.vue')
const RegisterView = () => import('@/view/RegisterView.vue')
const AskView = () => import('@/view/AskView.vue')
const InboxView = () => import('@/view/InboxView.vue')

/**@type {import('vue-router').RouteRecordRaw} */
const routes = [
    { path: '/', name: 'home', component: HomeView },
    { path: '/login', name: 'login', component: LoginView },
    { path: '/register', name: 'register', component: RegisterView },
    { path: '/:email', name: 'ask', component: AskView, props: true },
    { path: '/inbox', name: 'inbox', component: InboxView },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

router.beforeEach((to, from) => {
    const { title } = to.meta
    document.title = title || 'AskMe'
})

export default router