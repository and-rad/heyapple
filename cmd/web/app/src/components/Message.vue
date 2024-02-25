<script setup>
import SaveImage from "./images/ImageSave.vue";
import WarnImage from "./images/ImageWarn.vue";
import { onMounted } from "vue";

const prop = defineProps(["msg"]);
const emit = defineEmits(["timeout"]);

onMounted(() => {
	setTimeout(function () {
		emit("timeout", prop.msg.id);
	}, prop.msg.time);
});
</script>

<template>
	<div class="message" :class="[msg.type, msg.id]">
		<SaveImage v-if="msg.type == 'message'" />
		<WarnImage v-if="msg.type != 'message'" />
		<p>{{ msg.msg }}</p>
	</div>
</template>

<style>
.message {
	padding: 12px;
	background-color: #fff;
	border-radius: 4px;
	margin: 8px 16px;
	font-size: 14px;
	box-shadow: 0 0 8px #bbb;
	position: relative;
	display: flex;
	align-items: center;
}

.message svg {
	width: 16px;
	height: 16px;
	margin-right: 12px;
	fill: var(--color-text);
}

.message p {
	position: relative;
	z-index: 10;
}

.message:before {
	content: "";
	position: absolute;
	left: 0;
	top: 0;
	bottom: 0;
	border-radius: 4px 0 0 4px;
	border-right-color: transparent !important;
	border: 1px solid var(--color-good);
	background-color: var(--color-good-light);
	-webkit-animation: anim-msg-timeout 3s linear;
	animation: anim-msg-timeout 3s linear;
}

.message.error:before {
	border-color: var(--color-bad);
	background-color: var(--color-bad-light);
	-webkit-animation-duration: 5s;
	animation-duration: 5s;
}

.message.warning:before {
	background-color: var(--color-warn-light);
	border-color: var(--color-warn);
	-webkit-animation-duration: 4s;
	animation-duration: 4s;
}

@-webkit-keyframes anim-msg-timeout {
	0% {
		opacity: 1;
		width: 100%;
	}
	50% {
		opacity: 1;
	}
	100% {
		opacity: 0;
		width: 0;
	}
}

@keyframes anim-msg-timeout {
	0% {
		opacity: 1;
		width: 100%;
	}
	50% {
		opacity: 1;
	}
	100% {
		opacity: 0;
		width: 0;
	}
}
</style>
