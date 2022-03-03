<script setup></script>

<template>
	<footer>
		<nav>
			<span>{{ $t("foot.copy") }}</span>
			<a href="#">{{ $t("foot.legal") }}</a>
			<a href="#">{{ $t("foot.privacy") }}</a>
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
