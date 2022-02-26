import { createRouter, createWebHashHistory } from "vue-router";
import FoodView from "../views/FoodView.vue";

const router = createRouter({
	history: createWebHashHistory(),
	routes: [
		{
			path: "/",
			name: "food",
			component: FoodView,
		},
	],
});

export default router;
