import * as THREE from 'three';
import { STLLoader } from "three/addons/loaders/STLLoader.js";
import { OBJLoader } from "three/addons/loaders/OBJLoader.js";
import { OrbitControls } from "three/addons/controls/OrbitControls.js";

function use3DScene() {
	const stage = ref<{
		renderer: THREE.WebGLRenderer,
		camera: THREE.PerspectiveCamera,
		controls: OrbitControls,
		scene: THREE.Scene,
	} | null>();

	const loaders = {
		stl: new STLLoader(),
		obj: new OBJLoader(),
	};

	function setCanvas(canvas: HTMLCanvasElement) {
		const renderer = new THREE.WebGLRenderer({ canvas: canvas, antialias: true, alpha: true });
		const camera = new THREE.PerspectiveCamera(75, canvas.width / canvas.height, 0.1, 1000);
		const controls = new OrbitControls(camera, renderer.domElement);
		const scene = new THREE.Scene();

		stage.value = {
			renderer: renderer,
			camera: camera,
			controls: controls,
			scene: scene,
		};

		function animate() {
			renderer.render(scene, camera);
			controls.update();
		}

		renderer.setAnimationLoop(animate);
	}

	function setSize(width: number, height: number) {
		if (!stage.value) {
			console.warn('You have to set the canvas first');
			return;
		}

		stage.value.renderer.setSize(
			width, height
		);
		stage.value.camera.aspect = width / height;
		stage.value.camera.updateProjectionMatrix();
	}

	function setObject(geometry: THREE.BufferGeometry) {
		if (!stage.value) {
			throw new Error('You have to set the canvas first');
		}

		stage.value.scene.clear();

		const material = new THREE.MeshPhongMaterial({});
		const obj = new THREE.Mesh(geometry, material);
		obj.rotateX((-Math.PI * 2 * 1) / 4);
		stage.value.scene.add(obj);

		const bounding = new THREE.Box3().setFromObject(obj);
		obj.position.y = -bounding.min.y;

		function createLight(x: number, y: number, z: number): THREE.PointLight {
			const light = new THREE.PointLight(0xffffff);
			light.position.set(x, y, z);
			light.lookAt(0, 0, 0);

			return light;
		}

		const factor = 10.0001;

		stage.value.scene.add(
			createLight(
				bounding.max.x * factor,
				bounding.max.y * factor,
				bounding.max.z * factor
			)
		);
		stage.value.scene.add(createLight(bounding.max.x * factor, bounding.min.y * factor, 0));
		stage.value.scene.add(createLight(bounding.min.x * factor, bounding.max.y * factor, 0));
		stage.value.scene.add(createLight(bounding.min.x * factor, bounding.min.y * factor, 0));

		const light = new THREE.HemisphereLight(0xf6e86d, 0x404040, 0.5);
		stage.value.scene.add(light);

		const grid = new THREE.GridHelper(300, 30);
		stage.value.scene.add(grid);

		const pos = Math.max(bounding.max.x, bounding.max.y, bounding.max.z);

		stage.value.camera.position.z = pos * 1.5;
		stage.value.camera.position.y = pos * 2;
		stage.value.camera.position.x = pos;
	}

	function clear() {
		if (!stage.value) {
			return;
		}

		stage.value.scene.clear();
	}

	return {
		setCanvas: setCanvas,
		setSize: setSize,
		setObject: setObject,
		loaders: loaders,
		clear: clear,
	};
}

export { use3DScene };
