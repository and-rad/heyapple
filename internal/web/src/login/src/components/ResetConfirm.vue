<script setup>
import PasswordField from "./Password.vue";
import { RouterLink, useRoute } from "vue-router";
import { ref, inject } from "vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n();
const csrf = inject("csrfToken");

const pass = ref("");
const msg = ref({ msg: "", level: "" });
const passField = ref(null);

function confirm(evt) {
	evt.preventDefault();
	passField.value.hide();

	if (pass.value == "") {
		msg.value = { msg: t("register.errpassempty"), level: "err" };
		return false;
	}

	let form = evt.target.closest("form");
	evt.target.disabled = true;
	fetch("/auth/reset", {
		method: "PUT",
		headers: {
			"Content-Type": "application/x-www-form-urlencoded",
			"X-CSRF-Token": csrf,
		},
		body: new URLSearchParams(new FormData(form)),
	})
		.then((response) => {
			if (response.ok) {
				msg.value = { msg: t("resetconf.success") };
			} else {
				msg.value = { msg: t("resetconf.err" + response.status), level: "err" };
			}
		})
		.catch(() => (msg.value = { msg: t("err.conn"), level: "err" }))
		.finally(() => (evt.target.disabled = false));
}
</script>

<template>
	<header>
		<HeaderImage id="logo" />
		<div id="app-name">
			<span>{{ t("reset.title1") }}</span><span>{{ t("reset.title2") }}</span>
		</div>
	</header>
	<form>
		<p v-html="t('resetconf.hint')"></p>
		<Message v-bind="msg" />
		<label>{{ t("form.pass") }}</label>
		<PasswordField ref="passField" name="pass" withBar="true" v-model="pass" />
		<input type="hidden" name="token" :value="useRoute().params.token" />
		<button type="submit" @click="confirm" class="async">{{ t("resetconf.action") }}</button>
		<p><RouterLink to="/">{{ t("confirm.return") }}</RouterLink></p>
	</form>
</template>

<style></style>
