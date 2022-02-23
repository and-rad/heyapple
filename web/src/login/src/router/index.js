import { createRouter, createWebHashHistory, createWebHistory } from "vue-router";
import LoginView from "../components/Login.vue";

const router = createRouter({
	history: createWebHashHistory(),
	routes: [
		{
			path: "/",
			name: "home",
			component: LoginView,
		},
	],
});

export default router;
