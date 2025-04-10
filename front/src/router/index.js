import { createRouter, createWebHistory } from 'vue-router'
import LoginView from '../views/LoginView.vue'
import RegisterView from '../views/RegisterView.vue'
import DashboardView from '../views/DashboardView.vue'
import ScanView from '../views/Scan.vue'
import { useAuthStore } from '../stores/auth'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: '/dashboard'
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView,
      meta: { requiresAuth: false }
    },
    {
      path: '/register',
      name: 'register',
      component: RegisterView,
      meta: { requiresAuth: false }
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: DashboardView,
      meta: { requiresAuth: true }
    },
    {
      path: '/scan',
      name: 'scan',
      component: ScanView,
      meta: { requiresAuth: true }
    }
  ]
})

// router.beforeEach(async (to, from, next) => {
//   const authStore = useAuthStore()
//   const token = localStorage.getItem('token')
  
//   // 如果有token但store未初始化，验证token有效性
//   if (token && !authStore.isAuthenticated) {
//     try {
//       const response = await axios.get('http://127.0.0.1:9001/api/validate', {
//         headers: { Authorization: token }
//       })
//       authStore.login(response.data.user)
//     } catch {
//       localStorage.removeItem('token')
//     }
//   }

//   if (to.meta.requiresAuth && !authStore.isAuthenticated) {
//     next('/login')
//   } else if (!to.meta.requiresAuth && authStore.isAuthenticated) {
//     next('/dashboard')
//   } else {
//     next()
//   }
// })

export default router