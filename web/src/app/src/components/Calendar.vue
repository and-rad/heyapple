<script setup>
import ArrowImage from "./images/ImageRightArrow.vue";
import { ref } from "vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n();
const years = ref(["2020", "2021", "2022"]);
</script>

<template>
	<div class="calendar">
		<div>
			<button class="prev icon"><ArrowImage /></button>
			<div>
				<select class="month">
					<option value="0">{{ t("month.1") }}</option>
					<option value="1">{{ t("month.2") }}</option>
					<option value="2">{{ t("month.3") }}</option>
					<option value="3">{{ t("month.4") }}</option>
					<option value="4">{{ t("month.5") }}</option>
					<option value="5">{{ t("month.6") }}</option>
					<option value="6">{{ t("month.7") }}</option>
					<option value="7">{{ t("month.8") }}</option>
					<option value="8">{{ t("month.9") }}</option>
					<option value="9">{{ t("month.10") }}</option>
					<option value="10">{{ t("month.11") }}</option>
					<option value="11">{{ t("month.12") }}</option>
				</select>
				<select class="year">
					<option v-for="year in years" :value="year">{{ year }}</option>
				</select>
			</div>
			<button class="next icon"><ArrowImage /></button>
		</div>
		<table>
			<thead>
				<tr>
					<th>{{ t("day.cal7") }}</th>
					<th>{{ t("day.cal1") }}</th>
					<th>{{ t("day.cal2") }}</th>
					<th>{{ t("day.cal3") }}</th>
					<th>{{ t("day.cal4") }}</th>
					<th>{{ t("day.cal5") }}</th>
					<th>{{ t("day.cal6") }}</th>
				</tr>
			</thead>
			<tbody>
				<tr>
					<td><button>1</button></td>
					<td><button>2</button></td>
					<td><button>3</button></td>
					<td><button>4</button></td>
					<td><button>5</button></td>
					<td><button>6</button></td>
					<td><button>7</button></td>
				</tr>
				<tr>
					<td><button></button></td>
					<td><button></button></td>
					<td><button></button></td>
					<td><button></button></td>
					<td><button></button></td>
					<td><button></button></td>
					<td><button></button></td>
				</tr>
				<tr>
					<td><button></button></td>
					<td><button></button></td>
					<td><button></button></td>
					<td><button></button></td>
					<td><button></button></td>
					<td><button></button></td>
					<td><button></button></td>
				</tr>
				<tr>
					<td><button></button></td>
					<td><button></button></td>
					<td><button></button></td>
					<td><button></button></td>
					<td><button></button></td>
					<td><button></button></td>
					<td><button></button></td>
				</tr>
				<tr>
					<td><button></button></td>
					<td><button></button></td>
					<td><button></button></td>
					<td><button></button></td>
					<td><button></button></td>
					<td><button></button></td>
					<td><button></button></td>
				</tr>
				<tr>
					<td><button></button></td>
					<td><button></button></td>
					<td><button></button></td>
					<td><button></button></td>
					<td><button></button></td>
					<td><button></button></td>
					<td><button></button></td>
				</tr>
			</tbody>
		</table>
	</div>
</template>

<style>
.calendar {
	white-space: nowrap;
}

.calendar > div {
	display: flex;
	justify-content: space-between;
}

.calendar > div > div select {
	width: auto;
	display: inline-block;
}

.calendar > div > div .month {
	margin-right: 0.25em;
}

.calendar .prev svg {
	transform: rotate(180deg);
}

.calendar table {
	border-collapse: collapse;
	width: 100%;
	margin-top: 0.5em;
}

.calendar th,
.calendar td {
	width: calc(100% / 7);
	padding: 0;
}

.calendar th {
	font-weight: 700;
	color: var(--color-primary);
}

.calendar td button {
	width: 100%;
	height: 40px;
	background: none;
	border-radius: 0;
	color: var(--color-text);
	line-height: 1;
}
</style>
