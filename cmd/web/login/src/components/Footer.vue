<script setup>
import { useI18n } from "vue-i18n";
const { t } = useI18n();
</script>

<template>
	<footer>
		<nav>
			<span>{{ t("foot.copy") }}</span>
			<a href="/terms" target="_blank">{{ t("foot.terms") }}</a>
			<a href="/legal" target="_blank">{{ t("foot.legal") }}</a>
			<a href="/privacy" target="_blank">{{ t("foot.privacy") }}</a>
		</nav>
	</footer>
</template>

<style>
footer {
	font-size: 14px;
	padding: 1em;
	border-top: var(--border-light);
	text-align: center;
}

footer nav span,
footer nav a {
	color: #aaa;
	display: inline-block;
	margin: 0 0.5em;
}

footer nav a:hover {
	color: var(--color-primary);
}
</style>
