<script setup lang="ts">
import * as THREE from "three";
import { STLLoader } from "three/addons/loaders/STLLoader.js";
import { OBJLoader } from "three/addons/loaders/OBJLoader.js";
import { OrbitControls } from "three/addons/controls/OrbitControls.js";

const canvas = ref<HTMLCanvasElement | null>(null);
const container = ref<HTMLDivElement | null>(null);
const renderer = ref<THREE.WebGLRenderer | null>(null);
const controls = ref<OrbitControls | null>(null);
const camera = ref<THREE.Camera | null>(null);
const scene = new THREE.Scene();
const loaders = {
  stl: new STLLoader(),
  obj: new OBJLoader(),
};

const properties = defineProps<{
  files: File[];
}>();

function resize() {
  if (!renderer.value || !container.value) {
    return;
  }

  renderer.value.setSize(
    container.value.getBoundingClientRect().width,
    container.value.getBoundingClientRect().height
  );
}

function createScene(geometry: THREE.BufferGeometry, camera: THREE.Camera) {
  scene.clear();

  const material = new THREE.MeshPhongMaterial({});
  const obj = new THREE.Mesh(geometry, material);
  obj.rotateX((-Math.PI * 2 * 1) / 4);
  scene.add(obj);

  const bounding = new THREE.Box3().setFromObject(obj);
  obj.position.y = -bounding.min.y;

  function createLight(x: number, y: number, z: number): THREE.PointLight {
    const light = new THREE.PointLight(0xffffff);
    light.position.set(x, y, z);
    light.lookAt(0, 0, 0);

    return light;
  }

  const factor = 10.0001;

  scene.add(
    createLight(
      bounding.max.x * factor,
      bounding.max.y * factor,
      bounding.max.z * factor
    )
  );
  scene.add(createLight(bounding.max.x * factor, bounding.min.y * factor, 0));
  scene.add(createLight(bounding.min.x * factor, bounding.max.y * factor, 0));
  scene.add(createLight(bounding.min.x * factor, bounding.min.y * factor, 0));

  const light = new THREE.HemisphereLight(0xf6e86d, 0x404040, 0.5);
  scene.add(light);

  const grid = new THREE.GridHelper(300, 30);
  scene.add(grid);

  camera.position.z = bounding.max.z * 10;
  camera.position.y = bounding.max.y * 5;
}

watch(
  () => properties.files,
  (files) => {
    scene.clear();

    for (const file of files) {
      const uri = URL.createObjectURL(file);
      loaders.stl.load(
        uri,
        (geometry) => {
          createScene(geometry, new THREE.PerspectiveCamera());
          URL.revokeObjectURL(uri);
        },
        () => {},
        (error) => {
          console.warn("Error loading STL file", error);
          URL.revokeObjectURL(uri);
        }
      );
    }
  }
);

onMounted(() => {
  if (!canvas.value) {
    throw new Error("Canvas element not found");
  }

  window.addEventListener("resize", resize);

  renderer.value = new THREE.WebGLRenderer({
    antialias: true,
    canvas: canvas.value,
    alpha: true,
  });

  camera.value = new THREE.PerspectiveCamera(
    75,
    container.value!.getBoundingClientRect().width /
      container.value!.getBoundingClientRect().height,
    0.1,
    1000
  );
  renderer.value.setSize(
    container.value!.getBoundingClientRect().width,
    container.value!.getBoundingClientRect().height
  );

  controls.value = new OrbitControls(camera.value, canvas.value);

  const geometry = new THREE.BoxGeometry(1, 1, 1);
  createScene(geometry, camera.value);

  function animate() {
    if (!camera.value) {
      throw new Error("Camera not found");
    }

    renderer.value?.render(scene, camera.value);
    controls.value?.update();
  }

  renderer.value.setAnimationLoop(animate);
});
</script>

<template>
  <div
    ref="container"
    class="container ring-primary-500 ring-inset ring-1 shadow-sm rounded-md"
  >
    <canvas ref="canvas" />
  </div>
</template>

<style scoped>
.container {
  width: 100%;
  max-width: 400px;
  aspect-ratio: 1;
  position: relative;
}

.canvas {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
}
</style>
