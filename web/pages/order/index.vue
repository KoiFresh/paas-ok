<script setup lang="ts">
const files = ref<File[]>([]);
const result = ref<RemoteSlicerResult | null>(null);
const toast = useToast();
const options = ref<SliceOptions>({
  color: "red",
  material: "pla",
  infill: 30,
  quality: "low",
});
const { slice, progress } = useRemoteSlicer();

/**
 * Called when the uploaded files changed. Resets the slicing result.
 */
watch(files, () => (result.value = null));
watch(
  () =>
    options.value.color +
    options.value.material +
    options.value.infill +
    options.value.quality,
  () => (result.value = null)
);
const ready = computed(
  () => files.value.length && !result.value && !progress.value
);

/**
 * Called when the slice button is clicked. Slices the uploaded file.
 */
async function onSlice() {
  result.value = null;
  try {
    result.value = await slice(files.value[0], options.value);
  } catch (error) {
    toast.add({
      title: "Warning",
      description: String(error),
      icon: "i-heroicons-exclamation-triangle",
      color: "yellow",
      timeout: 10000,
    });
  }
}
</script>

<template>
  <div>
    <GenericHeadline />
    <div class="form">
      <GenericCadPreview :files="files" :color="options.color" />
      <UProgress v-if="progress" animation="carousel" />
      <FileInput v-model="files" />
      <GenericCadOptions v-model="options" :disabled="progress" />

      <div v-if="result" class="gap-1 flex">
        <UBadge :label="`Preis: ${result.price} €`" />
        <UBadge :label="`Material: ${result.material}`" />
      </div>
      <div class="submit">
        <UButton @click="onSlice" :disabled="!ready">
          Slice <Icon name="i-heroicons-square-3-stack-3d" />
        </UButton>
      </div>
    </div>
  </div>
</template>

<style scoped>
.form {
  display: flex;
  flex-direction: column;
  gap: 0.2em;
}

.submit {
  display: flex;
  justify-content: flex-end;
  margin-top: 1em;
}
</style>
