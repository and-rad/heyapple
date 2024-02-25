<script setup>
import { RouterLink } from "vue-router";
import { ref, inject } from "vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n();
const reEmail = /^[^@]+@[^@]+$/;
const csrf = inject("csrfToken");

const email = ref("");
const msg = ref({ msg: "", level: "" });

function confirm(evt) {
	evt.preventDefault();

	if (email.value.search(reEmail) == -1) {
		msg.value = { msg: t("login.errmail"), level: "err" };
		return false;
	}

	let addr = email.value;
	evt.target.disabled = true;
	fetch("/auth/reset", {
		method: "POST",
		headers: {
			"Content-Type": "application/x-www-form-urlencoded",
			"X-CSRF-Token": csrf,
		},
		body: new URLSearchParams(new FormData(evt.target.closest("form"))),
	}).then((response) => {
		if (response.ok) {
			msg.value = { msg: t("reset.success", { addr: addr }) };
			clearForm(evt);
		} else {
			msg.value = { msg: t("reset.err" + response.status), level: "err" };
			evt.target.disabled = false;
		}
	});
}

function clearForm(evt) {
	evt.target.disabled = false;
	email.value = "";
}
</script>

<template>
	<RouterLink to="/" class="back">&#5176; {{ t("form.back") }}</RouterLink>
	<header>
		<HeaderImage id="logo" />
		<div id="app-name">
			<span>{{ t("reset.title1") }}</span><span>{{ t("reset.title2") }}</span>
		</div>
	</header>
	<form>
		<p>{{ t("reset.hint") }}</p>
		<Message v-bind="msg" />
		<label>{{ t("form.email") }}</label>
		<input type="email" name="email" v-model="email" />
		<button type="submit" @click="confirm" class="async">{{ t("reset.action") }}</button>
	</form>
</template>

<style>
a.back {
	display: block;
	margin-bottom: 2em;
	transition: color var(--transition-style);
}

a.back:hover {
	color: var(--color-primary-dark);
}
</style>
