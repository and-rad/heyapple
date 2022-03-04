<template>
	<span class="icon sort-arrow"></span>
</template>

<style>
.sort-arrow {
	display: inline-block;
	width: 10px;
	height: 10px;
	background: center no-repeat var(--icon-triangle-up);
	background-size: 20px;
}

.desc .sort-arrow {
	background-image: var(--icon-triangle-down);
}
</style>