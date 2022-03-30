import { createRouter, createWebHashHistory } from "vue-router";
import LoginView from "../components/Login.vue";
import RegisterView from "../components/Register.vue";
import ResetView from "../components/Reset.vue";

const router = createRouter({
	history: createWebHashHistory(),
	routes: [
		{
			path: "/",
			name: "login",
			component: LoginView,
		},
		{
			path: "/signup",
			name: "signup",
			component: RegisterView,
		},
		{
			path: "/reset",
			name: "reset",
			component: ResetView,
		},
		{
			path: "/:pathmatch(.*)",
			redirect: "/",
		},
	],
});

export default router;
