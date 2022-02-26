<script setup></script>

<template>
	<footer>
		<nav>
			<a href="#">{{ $t("nav.legal") }}</a>
			<a href="#">{{ $t("nav.privacy") }}</a>
		</nav>
	</footer>
</template>

<style>
footer {
	padding: 1em;
	border-top: var(--border-light);
	text-align: center;
}

footer nav {
	font-size: 14px;
}

footer nav a {
	color: #aaaaaa;
	display: inline-block;
	margin: 0 0.5em;
}

footer nav a:hover {
	color: var(--color-primary);
}
</style>
