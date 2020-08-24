<?php

return [
	'routes' => [
		['name' => 'page#index', 'url' => '/', 'verb' => 'GET'],
		['name' => 'data#lists', 'url' => '/api/0.1/lists', 'verb' => 'GET'],
		['name' => 'data#completed', 'url' => '/api/0.1/completed', 'verb' => 'GET'],
		['name' => 'data#config', 'url' => '/api/0.1/config', 'verb' => 'GET'],
	]
];
