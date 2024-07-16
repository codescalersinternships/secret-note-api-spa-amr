import { createRouter, createWebHistory } from "vue-router";
import CreateNote from "../views/CreateNote.vue";
import ViewNote from "../views/ViewNote.vue";
import SignIn from "../views/SignIn.vue";
import SignUp from "../views/SignUp.vue";

const routes = [
  {
    path: "/create",
    name: "CreateNote",
    component: CreateNote,
    meta: { requiresAuth: true },
  },
  {
    path: "/:url",
    name: "ViewNote",
    component: ViewNote,
  },
  {
    path: "/signin",
    name: "SignIn",
    component: SignIn,
  },
  {
    path: "/signup",
    name: "SignUp",
    component: SignUp,
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

router.beforeEach((to, from, next) => {
  const requiresAuth = to.matched.some(record => record.meta.requiresAuth);
  const isAuthenticated = !!localStorage.getItem('user');

  if (requiresAuth && !isAuthenticated) {
    next('/signin');
  } else {
    next();
  }
});

export default router;
