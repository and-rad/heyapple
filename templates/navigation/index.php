<div>
	<div id='calendar2'>
		<div>
			<button class='app-navigation-noclose'>
				<span class='arrow prev'></span>
			</button>
			<div>
				<select class='month'>
					<option value='0'><?php p($l->t('month.1')); ?></option>
					<option value='1'><?php p($l->t('month.2')); ?></option>
					<option value='2'><?php p($l->t('month.3')); ?></option>
					<option value='3'><?php p($l->t('month.4')); ?></option>
					<option value='4'><?php p($l->t('month.5')); ?></option>
					<option value='5'><?php p($l->t('month.6')); ?></option>
					<option value='6'><?php p($l->t('month.7')); ?></option>
					<option value='7'><?php p($l->t('month.8')); ?></option>
					<option value='8'><?php p($l->t('month.9')); ?></option>
					<option value='9'><?php p($l->t('month.10')); ?></option>
					<option value='10'><?php p($l->t('month.11')); ?></option>
					<option value='11'><?php p($l->t('month.12')); ?></option>
				</select>
				<select class='year'>
					<?php $year = date('Y'); ?>
					<option value='<?php echo $year-1; ?>'><?php echo $year-1; ?></option>
					<option value='<?php echo $year; ?>'><?php echo $year; ?></option>
					<option value='<?php echo $year+1; ?>'><?php echo $year+1; ?></option>
				</select>
			</div>
			<button class='app-navigation-noclose'>
				<span class='arrow next'></span>
			</button>
		</div>
		<table>
			<thead>
				<tr>
					<th><?php p($l->t('day.7')); ?></th>
					<th><?php p($l->t('day.1')); ?></th>
					<th><?php p($l->t('day.2')); ?></th>
					<th><?php p($l->t('day.3')); ?></th>
					<th><?php p($l->t('day.4')); ?></th>
					<th><?php p($l->t('day.5')); ?></th>
					<th><?php p($l->t('day.6')); ?></th>
				</tr>
			</thead>
			<tbody>
				<tr>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
				</tr>
				<tr>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
				</tr>
				<tr>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
				</tr>
				<tr>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
				</tr>
				<tr>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
				</tr>
				<tr>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
					<td><button class='app-navigation-noclose'></button></td>
				</tr>
			</tbody>
		</table>
	</div>
	<div class='app-navigation-new' style='display:none;'>
		<button class='icon-add'><?php p($l->t('newlist')); ?></button>
	</div>
	<ul id='list-category'></ul>
</div>

