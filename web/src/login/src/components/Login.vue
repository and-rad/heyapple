<script setup>
import { RouterLink } from "vue-router";
import { ref } from "vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n();
const reEmail = /^[^@]+@[^@]+$/;

const email = ref("");
const msg = ref({ msg: "", level: "" });

function confirm(evt) {
	evt.preventDefault();
	msg.value = {};

	if (email.value.search(reEmail) == -1) {
		msg.value = { msg: t("login.errmail"), level: "err" };
		return false;
	}

	fetch("/auth/local", {
		method: "POST",
		headers: { "Content-Type": "application/x-www-form-urlencoded" },
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
		<input type="password" name="pass" />
		<footer>
			<span>{{ $t("login.signup") }} <RouterLink to="/signup">{{ $t("register.action") }}</RouterLink>.</span>
		</footer>
		<input type="submit" :value="$t('login.action')" @click="confirm" />
	</form>
</template>

<style>
</style>
