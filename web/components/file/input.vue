<script setup lang="ts">
const model = defineModel<File[]>({
  default: [],
});

/**
 * Called when the uploaded files changed
 * @param event FileList
 */
function onFilesChanged(event: FileList) {
  const uploads: File[] = [];

  for (let i = 0; i < event.length; i++) {
    const file = event.item(i);
    if (!file) {
      continue;
    }
    uploads.push(file);
  }

  model.value = uploads;
}
</script>

<template>
  <UInput
    variant="outline"
    color="primary"
    type="file"
    accept=".stl"
    icon="i-heroicons-folder"
    @change="onFilesChanged"
  />
</template>
