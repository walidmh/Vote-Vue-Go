import VueRouter from 'vue-router'
import HomePage from '../views/HomePage'
import LoginPage from '../views/LoginPage'
import SignupPage from '../views/SignupPage'
import UserPage from '../views/UserPage'
import CreatePollPage from '../views/CreatePollPage'
import ErrorPage from '../views/ErrorPage'
import NotFound from '../views/NotFound'
import store from '../helpers/store'

const router = new VueRouter({
  mode: "history",
  routes: [
    {
      path: "/",
      name: "HomePage",
      component: HomePage,
      meta: { 
        requiresAuth: true
      }
    },
    {
      path: "/login",
      name: "Login",
      component: LoginPage
    },
    {
      path: "/signup",
      name: "Signup",
      component: SignupPage
    },
    {
      path: "/profil",
      name: "User",
      component: UserPage,
      meta: { 
        requiresAuth: true
      }
    },
    {
      path: "/createPoll",
      name: "CreatePoll",
      component: CreatePollPage,
      meta: { 
        requiresAuth: true
      }
    },
    {
      path: "/error",
      name: "Error",
      component: ErrorPage
    },
    {
      path: "*",
      component: NotFound
    }
  ]
})

router.beforeEach((to, from, next) => {
  if(to.matched.some(record => record.meta.requiresAuth)) {
    if (store.getters.isLoggedIn) {
      console.log('loggé')
      next()
    }
    else{
      console.log('not logged')
      next({
          path: '/error',
          query: { redirect: to.fullPath }
        })
    }
  } else {
    next() 
  }
})

export default router;

