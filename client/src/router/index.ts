import Vue from "vue";
import VueRouter, { RouteConfig } from "vue-router";

Vue.use(VueRouter);

const routes: Array<RouteConfig> = [
  {
    name: "Problem",
    path: "/problem/:id",
    component: () =>
      import(/* webpackChunkName: "problem" */ "../components/Problem.vue"),
    children: [
      {
        name: "EditProblem",
        // EditProblem will be rendered inside Problem's <router-view>
        // when /problem/:id/edit is matched
        path: "edit",
        component: () =>
          import(
            /* webpackChunkName: "edit-problem" */ "../components/EditProblem.vue"
          )
      },
      {
        name: "SolveProblem",
        // SolveProblem will be rendered inside Problem's <router-view>
        // when /problem/:id/solve is matched
        path: "solve",
        component: () =>
          import(
            /* webpackChunkName: "solve-problem" */ "../components/SolveProblem.vue"
          )
      }
    ]
  },
  {
    name: "AllProblems",
    path: "/",
    component: () =>
      import(
        /* webpackChunkName: "all-problems" */ "../components/AllProblems.vue"
      )
  },
  {
    name: "LandingPage",
    path: "/landing",
    component: () =>
      import(
        /* webpackChunkName: "landing-page" */ "../components/LandingPage.vue"
      )
  },
  {
    name: "LastProblem",
    path: "/last-problem",
    redirect: to => {
      return { name: 'SolveProblem', params: { id: "29fe43a0-5874-44e2-9ad8-1b32ea4621ab"} }
    }

  },
  {
    name: "SavedProblems",
    path: "/saved-problems",
    component: () =>
      import(
        /* webpackChunkName: "saved-problems" */ "../components/SavedProblems.vue"
      )
  },
  {
    name: "MyProblems",
    path: "/my-problems",
    component: () =>
      import(
        /* webpackChunkName: "my-problems" */ "../components/MyProblems.vue"
      )
  },
  {
    name: "History",
    path: "/history",
    component: () =>
      import(
        /* webpackChunkName: "history" */ "../components/History.vue"
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
