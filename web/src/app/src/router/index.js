import { createRouter, createWebHashHistory } from "vue-router";
import FoodView from "../views/FoodView.vue";
import RecipeView from "../views/RecipeView.vue";
import DiaryView from "../views/DiaryView.vue";
import ShoppingView from "../views/ShoppingView.vue";
import ProfileView from "../views/ProfileView.vue";
import SettingsView from "../views/SettingsView.vue";

const router = createRouter({
	history: createWebHashHistory(),
	routes: [
		{
			path: "/",
			name: "diary",
			component: DiaryView,
		},
		{
			path: "/recipes",
			name: "recipes",
			component: RecipeView,
		},
		{
			path: "/food",
			name: "food",
			component: FoodView,
		},
		{
			path: "/shopping",
			name: "shopping",
			component: ShoppingView,
		},
		{
			path: "/profile",
			name: "profile",
			component: ProfileView,
		},
		{
			path: "/settings",
			name: "settings",
			component: SettingsView,
		},
	],
});

export default router;
