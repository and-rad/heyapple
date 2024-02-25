import { createRouter, createWebHashHistory } from "vue-router";
import LoginView from "../components/Login.vue";
import RegisterView from "../components/Register.vue";
import ResetView from "../components/Reset.vue";
import ConfirmView from "../components/Confirm.vue";
import ResetConfirmView from "../components/ResetConfirm.vue";

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
			path: "/confirm/:token",
			name: "confirm",
			component: ConfirmView,
		},
		{
			path: "/reset",
			name: "reset",
			component: ResetView,
		},
		{
			path: "/reset/:token",
			name: "resetconfirm",
			component: ResetConfirmView,
		},
		{
			path: "/:pathmatch(.*)",
			redirect: "/",
		},
	],
});

export default router;
