import { createRouter, createWebHashHistory } from "vue-router";

// Generate routes from files under `src/pages`.
// - `pages/index.vue` -> `/`
// - `pages/about.vue` -> `/about`
// - `pages/features.vue` -> `/features`
const pageModules = import.meta.glob("../pages/*.vue");

const routes = Object.keys(pageModules)
  .sort()
  .map((file) => {
    const name = file
      .split("/")
      .pop()
      .replace(/\.vue$/, "");
    const path = name === "Index" ? "/" : `/${name}`;

    return {
      path,
      name,
      component: pageModules[file],
    };
  });

// Fallback to home for unknown routes.
routes.push({ path: "/:pathMatch(.*)*", redirect: "/" });

export const router = createRouter({
  history: createWebHashHistory(),
  routes,
});
