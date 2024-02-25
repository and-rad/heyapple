<script setup>
import { RouterLink, useRoute } from "vue-router";
import { ref, inject, onMounted } from "vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n();
const csrf = inject("csrfToken");
const msg = ref({ msg: "", level: "" });

function confirm(token) {
	fetch("/auth/confirm", {
		method: "PUT",
		headers: {
			"Content-Type": "application/x-www-form-urlencoded",
			"X-CSRF-Token": csrf,
		},
		body: new URLSearchParams({ token: token }),
	}).then((response) => {
		if (response.ok) {
			msg.value = { msg: t("confirm.success") };
		} else {
			msg.value = { msg: t("confirm.err" + response.status), level: "err" };
		}
	});
}

onMounted(() => {
	confirm(useRoute().params.token);
});
</script>

<template>
	<header>
		<HeaderImage id="logo" />
		<div id="app-name">
			<span>{{ t("confirm.title1") }}</span><span>{{ t("confirm.title2") }}</span>
		</div>
	</header>
	<Message v-bind="msg" />
	<RouterLink to="/">{{ t("confirm.return") }}</RouterLink>
</template>

<style></style>
