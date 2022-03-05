<script setup>
const prop = defineProps(["data", "placeholder"]);
const emit = defineEmits(["result"]);

var timeoutHandle = undefined;

function confirm(evt) {
	evt.preventDefault();
	clearTimeout(timeoutHandle);
	timeoutHandle = setTimeout(function () {
		fetchData(evt.target.closest("form"));
	}, 500);
}

function fetchData(form) {
	let formData = new FormData(form);
	let filtered = prop.data.filter((d) => {
		for (let k of formData.keys()) {
			// special cases first
			if (k == "name") {
				let name = formData.get(k).toLowerCase();
				if (!d[k].toLowerCase().includes(name)) {
					return false;
				}
				continue;
			}

			// numeric values next
			let [first, last] = formData.getAll(k).map((v) => parseFloat(v));
			if (!isNaN(first) && !isNaN(last)) {
				if (d[k] < first || last < d[k]) {
					return false;
				}
			}
		}
		return true;
	});
	emit("result", filtered);
}
</script>

<template>
	<form>
		<input type="text" name="name" :placeholder="placeholder" @input="confirm" />
		<slot :confirm="confirm">Additional filters</slot>
	</form>
</template>

<style></style>
