<script setup lang="ts">
const canvas = ref<HTMLCanvasElement | null>(null);
const container = ref<HTMLDivElement | null>(null);
const scene = use3DScene();

const properties = defineProps<{
  files: File[];
}>();

function resize() {
  scene.setSize(
    container.value!.getBoundingClientRect().width,
    container.value!.getBoundingClientRect().height
  );
}

watch(
  () => properties.files,
  (files) => {
    scene.clear();

    for (const file of files) {
      const uri = URL.createObjectURL(file);
      scene.loaders.stl.load(
        uri,
        (geometry) => {
          scene.setObject(geometry);
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
  scene.setCanvas(canvas.value);
  resize();
});
</script>

<template>
  <div
    ref="container"
    class="container ring-primary-500 ring-inset ring-1 shadow-sm rounded-md"
  >
    <div v-if="!files.length" class="text-center text-gray-500 m-2">
      Wähle eine Datei aus
    </div>
    <canvas class="canvas" ref="canvas" />
  </div>
</template>

<style scoped>
.container {
  min-width: 100%;
  aspect-ratio: 1;
  max-height: 400px;
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
