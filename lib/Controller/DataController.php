<?php
namespace OCA\HeyApple\Controller;

use OCP\IConfig;
use OCP\IL10N;
use OCP\IRequest;
use OCP\AppFramework\Http\JSONResponse;
use OCP\AppFramework\Controller;
use OCP\Files\IRootFolder;

class DataController extends Controller {
	private $userId;
	private $config;
	private $root;
	private $l;

	public function __construct(
		$AppName,
		IRequest $request,
		$UserId,
		IConfig $config,
		IRootFolder $rootFolder,
		IL10N $l
	) {
		parent::__construct($AppName, $request);
		$this->userId = $UserId;
		$this->config = $config;
		$this->root = $rootFolder;
		$this->l = $l;
	}

	/**
	 * @NoAdminRequired
	 */
	public function lists() : JSONResponse {
		$dir = $this->config->getUserValue($this->userId, $this->appName, 'directory');
		$data = $this->loadLists($dir);
		$aisles = $this->loadAisles($dir);

		foreach ($data as &$list) {
			foreach ($list as &$item) {
				foreach ($aisles as $aisle) {
					if (in_array($item[2],$aisle->ids)) {
						$item[] = $this->l->t($aisle->key);
						break;
					}
				}
			}
		}

		$ok = count($data) > 0;
		$msg = $ok ? "msg.ok" : "err.empty";

		return new JSONResponse(['success' => $ok, 'message' => $msg, 'data' => $data]);
	}

	/**
	 * @NoAdminRequired
	 */
	public function completed() : JSONResponse {
		$dir = $this->config->getUserValue($this->userId, $this->appName, 'directory');
		list($ok, $msg) = $this->checkDirectory($dir);
		if (!$ok) {
			return new JSONResponse(['success' => $ok, 'message' => $msg]);
		}

		$data = $this->loadCompleted($this->root->getUserFolder($this->userId)->get($dir));
		$ok = $data != NULL;
		$msg = $ok ? "msg.ok" : "err.empty";

		return new JSONResponse(['success' => $ok, 'message' => $msg, 'data' => $data]);
	}

	/**
	 * @NoAdminRequired
	 */
	public function config() : JSONResponse {
		$dir = $this->config->getUserValue($this->userId, $this->appName, 'directory');

		return new JSONResponse([
			'directory' => $dir,
		]);
	}

	/**
	 * @NoAdminRequired
	 */
	public function scan(string $dir) : JSONResponse {
		list($ok, $msg) = $this->checkDirectory($dir);

		if ($ok) {
			$this->config->setUserValue($this->userId, $this->appName, 'directory', $dir);
			return $this->lists();
		}

		return new JSONResponse(['success' => $ok, 'message' => $msg]);
	}

	/**
	 * @NoAdminRequired
	 */
	public function complete() : JSONResponse {
		list($ok, $msg) = $this->saveCompleted($body = $this->request->getParams());
		return new JSONResponse(['success' => $ok, 'message' => $msg]);
	}

	private function checkDirectory($dir) : array {
		$root = $this->root->getUserFolder($this->userId);
		if (!$root->nodeExists($dir)) {
			return [false, "err.nodir"];
		}

		$node = $root->get($dir);
		if ($node->getType() != \OCP\Files\FileInfo::TYPE_FOLDER) {
			return [false, "err.notdir"];
		}

		if (!$node->isUpdateable()) {
			return [false, "err.rodir"];
		}

		return [true, 'dir.ok'];
	}

	private function loadLists(string $dir) : array {
		$data = [];

		$node = $this->root->getUserFolder($this->userId)->get($dir);
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

	private function loadAisles(string $dir) : array {
		$data = [];

		$node = $this->root->getUserFolder($this->userId)->get($dir);
		if ($node->nodeExists('aisles.json')) {
			$data = json_decode($node->get('aisles.json')->getContent());
		}

		return $data;
	}

	private function loadCompleted($node) : ?object {
		$data = NULL;

		if ($node->nodeExists('completed.json')) {
			$data = json_decode($node->get('completed.json')->getContent());
		}

		return $data;
	}

	private function saveCompleted($data) : array {
		unset($data["_route"]);
		$data = json_encode($data);

		$dir = $this->config->getUserValue($this->userId, $this->appName, 'directory');
		$node = $this->root->getUserFolder($this->userId)->get($dir);
		if (!$node->nodeExists('completed.json')) {
			if (!$node->newFile('completed.json')) {
				return [false, 'err.fnew'];
			}
		}

		$node->get('completed.json')->putContent($data);
		return [true, 'msg.ok'];
	}

	private function abs($node) : string {
		return $this->config->getSystemValue('datadirectory').$node->getPath();
	}
}
