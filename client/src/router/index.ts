import Vue from "vue";
import VueRouter, { RouteConfig } from "vue-router";

Vue.use(VueRouter);

const routes: Array<RouteConfig> = [
  {
    name: "Problem",
    path: "/problem/:id",
    component: () =>
      import(/* webpackChunkName: "problem" */ "../components/Problem.vue"),
    children: [{
      name: "EditProblem",
      // EditProblem will be rendered inside Problem's <router-view>
      // when /problem/:id/edit is matched
      path: 'edit',
      component: () =>
        import(/* webpackChunkName: "edit-problem" */ "../components/EditProblem.vue"),
    },
    {
      name: "SolveProblem",
      // SolveProblem will be rendered inside Problem's <router-view>
      // when /problem/:id/solve is matched
      path: 'solve',
      component: () =>
        import(/* webpackChunkName: "solve-problem" */ "../components/SolveProblem.vue")
    }]
  },
  {
    name: "AllProblems",
    path: "/",
    component: () =>
      import(
        /* webpackChunkName: "all-problems" */ "../components/AllProblems.vue"
      )
  },
  { path: "*", redirect: "/" }
];

// 3. Create the router instance and pass the `routes` option
// You can pass in additional options here, but let's
// keep it simple for now.
export const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes // short for `routes: routes`
});
