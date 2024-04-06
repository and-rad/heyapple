<script setup>
import Main from "../components/Main.vue";
import PasswordField from "../../../login/src/components/Password.vue";
import { ref, inject } from "vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n();
const csrf = inject("csrfToken");
const log = inject("log");
const prefs = inject("prefs");

const isSaving = ref(false);

const main = ref(null);

function onNavItem(id) {
	let target = document.getElementById(id);
	target.scrollIntoView({ behavior: "smooth", block: "start" });
}

function onSaveEmail(evt) {
	isSaving.value = true;
	let form = evt.target.form;

	fetch("/auth/email", {
		method: "PUT",
		headers: {
			"Content-Type": "application/x-www-form-urlencoded",
			"X-CSRF-Token": csrf,
		},
		body: new URLSearchParams(new FormData(form)),
	})
		.then((response) => {
			if (!response.ok) {
				throw t("savemail.err" + response.status);
			}
			log.msg(t("savemail.ok"));
		})
		.catch((err) => log.err(err))
		.finally(() => (isSaving.value = false));
}

function onRollUsername(evt) {
	isSaving.value = true;

	fetch("/api/v1/name", {
		method: "PUT",
		headers: { "X-CSRF-Token": csrf },
	})
		.then((response) => {
			if (!response.ok) {
				throw t("savename.err" + response.status);
			}
			return response.json();
		})
		.then((data) => {
			prefs.value.account.name = data.name;
			log.msg(t("savename.ok"));
		})
		.catch((err) => log.err(err))
		.finally(() => (isSaving.value = false));
}

function onChangePassword(evt) {
	isSaving.value = true;
	let form = evt.target.form;

	fetch("/auth/pass", {
		method: "PUT",
		headers: {
			"Content-Type": "application/x-www-form-urlencoded",
			"X-CSRF-Token": csrf,
		},
		body: new URLSearchParams(new FormData(form)),
	})
		.then((response) => {
			if (!response.ok) {
				throw t("savepass.err" + response.status);
			}
			log.msg(t("savepass.ok"));
		})
		.catch((err) => log.err(err))
		.finally(() => (isSaving.value = false));
}

function onDeleteUser(evt) {
	isSaving.value = true;

	fetch("/api/v1/user", {
		method: "DELETE",
		headers: { "X-CSRF-Token": csrf },
	})
		.then((response) => {
			if (!response.ok) {
				throw t("deluser.err" + response.status);
			}
			log.msg(t("deluser.ok"));
			setTimeout(() => {
				window.location = "/";
			}, 3000);
		})
		.catch((err) => log.err(err))
		.finally(() => (isSaving.value = false));
}
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
					<li>
						<a @click="onNavItem('head-danger')"> {{ t("nav.danger") }} </a>
					</li>
				</ul>
			</nav>
		</template>

		<template #main>
			<section>
				<h2 id="head-account">{{ t("nav.account") }}</h2>
				<form @submit.prevent>
					<label>{{ t("profile.email") }}</label>
					<input type="email" name="email" :value="prefs.account.email" />
					<p v-html="t('profile.emailhint')"></p>
					<button type="button" :disabled="isSaving" @click="onSaveEmail" class="async">
						{{ t("btn.save") }}
					</button>
				</form>
				<form @submit.prevent>
					<label>
						{{ t("profile.name") }} <a href="#">{{ t("profile.namelink") }}</a>
					</label>
					<input readonly type="text" name="name" :value="prefs.account.name" />
					<p v-html="t('profile.namehint')"></p>
					<button type="button" :disabled="isSaving" @click="onRollUsername" class="async">
						{{ t("btn.reroll") }}
					</button>
				</form>
				<form @submit.prevent>
					<label>{{ t("profile.passold") }}</label>
					<PasswordField ref="passField" name="passold" />
					<label>{{ t("profile.passnew") }}</label>
					<PasswordField ref="passField" name="passnew" withBar="true" />
					<p v-html="t('profile.passhint')"></p>
					<button type="button" :disabled="isSaving" @click="onChangePassword" class="async">
						{{ t("btn.changepass") }}
					</button>
				</form>
			</section>

			<section>
				<h2 id="head-body">{{ t("nav.body") }}</h2>
			</section>

			<section class="danger">
				<h2 id="head-danger">{{ t("nav.danger") }}</h2>
				<form @submit.prevent>
					<p v-html="t('profile.deletehint')"></p>
					<button type="button" :disabled="isSaving" @click="onDeleteUser" class="async">
						{{ t("btn.deleteuser") }}
					</button>
				</form>
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

main.settings .content section.danger {
	background-color: var(--color-warn-light);
	border: 1px solid var(--color-warn);
	border-radius: 8px;
	padding: 1em;
}

main.settings .content section.danger form:last-child {
	margin-bottom: 0;
}

main.settings .content section.danger h2 {
	color: var(--color-secondary);
}

main.settings .content section.danger button {
	background-color: var(--color-secondary);
}

main.settings .content form {
	margin: 2em 0 4em;
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

main.settings label > a {
	float: right;
}

main.settings input + p,
main.settings span.password + p {
	margin-top: 1em;
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
		right: 0;
	}

	main.settings .content {
		padding-right: calc(300px + 1em);
	}
}
</style>
