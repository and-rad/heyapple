<?php
namespace OCA\HeyApple\Controller;

use OCP\IConfig;
use OCP\IRequest;
use OCP\AppFramework\Http\JSONResponse;
use OCP\AppFramework\Controller;
use OCP\Files\IRootFolder;

class DataController extends Controller {
	private $userId;
	private $config;
	private $root;

	public function __construct(
		$AppName,
		IRequest $request,
		$UserId,
		IConfig $config,
		IRootFolder $rootFolder
	) {
		parent::__construct($AppName, $request);
		$this->userId = $UserId;
		$this->config = $config;
		$this->root = $rootFolder;
	}

	/**
	 * @NoAdminRequired
	 */
	public function lists() : JSONResponse {
		$dir = $this->config->getUserValue($this->userId, $this->appName, 'directory', 'FDDB');

		$root = $this->root->getUserFolder($this->userId);
		if (!$root->nodeExists($dir)) {
			return new JSONResponse(['success' => false, 'message' => "directory doesn't exist"]);
		}

		$node = $root->get($dir);
		if ($node->getType() != \OCP\Files\FileInfo::TYPE_FOLDER) {
			return new JSONResponse(['success' => false, 'message' => "not a directory"]);
		}

		if (!$node->isUpdateable()) {
			return new JSONResponse(['success' => false, 'message' => "read-only directory"]);
		}

		$data = $this->loadLists($node);
		$ok = count($data) > 0;
		$msg = $ok ? "lists.ok" : "lists.err";

		return new JSONResponse(['success' => $ok, 'message' => $msg, 'data' => $data]);
	}

	/**
	 * @NoAdminRequired
	 */
	public function completed() : JSONResponse {
		return new JSONResponse(['success' => $ok, 'message' => $msg, 'data' => $data]);
	}

	private function loadLists($node) : array {
		$data = [];

		foreach ($node->getDirectoryListing() as $f) {
			if (strcasecmp($f->getExtension(), 'csv') == 0) {
				$csv = array_map(function($a){ return str_getcsv($a, ";"); }, file($this->abs($f)));
				array_walk($csv, function(&$a) use ($csv) {
					$a = array_slice($a, 0, 3);
					if (!mb_check_encoding($a[1],'UTF-8')) {
						$a[1] = utf8_encode($a[1]);
					}
				});
				array_shift($csv);
				$data[basename($f->getName(), '.csv')] = $csv;
			}
		}

		return $data;
	}

	private function abs($node) : string {
		return $this->config->getSystemValue('datadirectory').$node->getPath();
	}
}
