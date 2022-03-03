<script setup>
import PasswordField from "./Password.vue"
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

	fetch("/auth/local", {
		method: "POST",
		headers: {
			"Content-Type": "application/x-www-form-urlencoded",
			"X-CSRF-Token": csrf,
		},
		body: new URLSearchParams(new FormData(evt.target.closest("form"))),
	}).then((response) => {
		if (response.ok) {
			window.location = "/app";
		} else {
			msg.value = { msg: t("login.err" + response.status), level: "err" };
		}
	});
}
</script>

<template>
	<header>
		<HeaderImage id="logo" />
		<div id="app-name">
			<span>{{ $t("login.title1") }}</span><span>{{ $t("login.title2") }}</span>
		</div>
	</header>
	<form>
		<Message v-bind="msg" />
		<label>{{ $t("form.email") }}</label>
		<input type="email" name="email" v-model="email" />
		<label>{{ $t("form.pass") }} <RouterLink to="/reset">{{ $t("form.reset") }}</RouterLink></label>
		<PasswordField name="pass" />
		<input type="submit" :value="$t('login.action')" @click="confirm" />
		<p>{{ $t("login.signup") }} <RouterLink to="/signup">{{ $t("register.action") }}</RouterLink>.</p>
	</form>
</template>

<style>
</style>
