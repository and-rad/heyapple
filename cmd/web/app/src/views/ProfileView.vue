<script setup>
import Main from "../components/Main.vue";
import PasswordField from "../../../login/src/components/Password.vue";
import { ref, inject } from "vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n();

const main = ref(null);

function onNavItem(id) {
	let target = document.getElementById(id);
	target.scrollIntoView({ behavior: "smooth", block: "start" });
}

function onSaveEmail(evt) {}

function onChangePassword(evt) {}
</script>

<template>
	<Main ref="main" class="settings no-dt">
		<template #filter>
			<nav>
				<ul>
					<li>
						<a @click="onNavItem('head-account')"> {{ t("nav.account") }} </a>
					</li>
					<li>
						<a @click="onNavItem('head-body')"> {{ t("nav.body") }} </a>
					</li>
				</ul>
			</nav>
		</template>

		<template #main>
			<section>
				<h2 id="head-account">{{ t("nav.account") }}</h2>
				<form>
					<label>{{ t("profile.email") }}</label>
					<input type="email" name="email" />
					<button type="submit" @click="onSaveEmail" class="async">{{ t("btn.save") }}</button>
				</form>
				<form>
					<label>{{ t("profile.passold") }}</label>
					<PasswordField ref="passField" name="pass1" />
					<label>{{ t("profile.passnew") }}</label>
					<PasswordField ref="passField" name="pass2" withBar="true" />
					<button type="submit" @click="onChangePassword" class="async">{{ t("btn.changepass") }}</button>
				</form>
			</section>

			<section>
				<h2 id="head-body">{{ t("nav.body") }}</h2>
			</section>
		</template>
	</Main>
</template>

<style>
#filter nav ul {
	list-style: none;
	white-space: nowrap;
	padding: 0;
	margin: 1em 0;
}

#filter nav a {
	display: block;
	padding: 0.5em 1em;
	color: var(--color-text);
	border-left: 3px solid transparent;
	transition: color var(--transition-style);
	cursor: pointer;
}

#filter nav a:hover,
#filter nav a:active,
#filter nav a:focus {
	color: var(--color-primary);
}

#filter nav a.selected {
	border-color: var(--color-secondary);
}

main.settings #main .controls {
	border-bottom: var(--border-light);
}

main.settings .content {
	padding: 1em;
}

main.settings .content h2 {
	margin: 2em 0 1em;
}

main.settings .content h2:first-child {
	margin-top: 0em;
}

main.settings .content section {
	max-width: 480px;
	margin: auto;
}

main.settings .content form {
	margin: 2em 0;
}

main.settings .content label {
	white-space: nowrap;
	text-overflow: ellipsis;
	overflow-x: hidden;
	color: var(--color-text-light);
	margin-top: 1em;
	display: block;
}

main.settings .content form > button {
	margin: 1em 0 1em;
}

@media only screen and (min-width: 500px) {
	main.settings .content form > button {
		margin-left: auto;
		width: 240px;
		min-width: 240px;
	}
}

@media only screen and (min-width: 800px) {
	main.settings #main .controls {
		display: none;
	}
}

@media only screen and (min-width: 1280px) {
	main.settings #main {
		right: 300px;
	}
}
</style>