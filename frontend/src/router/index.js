import { createRouter, createWebHistory } from "vue-router";
import CreateNote from "../views/CreateNote.vue";
import ViewNote from "../views/ViewNote.vue";

const routes = [
  {
    path: "/",
    name: "CreateNote",
    component: CreateNote,
  },
  {
    path: "/:url",
    name: "ViewNote",
    component: ViewNote,
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
