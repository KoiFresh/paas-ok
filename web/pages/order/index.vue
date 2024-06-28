<script setup lang="ts">
const files = ref<File[]>([]);

const result = ref<RemoteSlicerResult | null>(null);
const warning = ref<string | null>(null);
const { slice, isSlicing } = useRemoteSlicer();

async function onSlice() {
  result.value = null;
  warning.value = null;
  try {
    result.value = await slice(files.value[0]);
  } catch (error) {
    warning.value = String(error);
  }
}

/**
 * Called when the uploaded files changed
 * @param event FileList
 */
function onFilesChanged(event: FileList) {
  result.value = null;
  warning.value = null;
  const uploads: File[] = [];

  for (let i = 0; i < event.length; i++) {
    const file = event.item(i);
    if (!file) {
      continue;
    }
    uploads.push(file);
  }

  files.value = uploads;
}
</script>

<template>
  <div>
    <div>
      <h1>
        <OkLogo class="logo" />
        Print Everything, Everywhere all at once!
      </h1>
    </div>
    <div class="form">
      <OrderPreview :files="files" />
      <UProgress v-if="isSlicing" animation="carousel" />
      <UInput
        variant="outline"
        color="primary"
        type="file"
        accept=".stl"
        icon="i-heroicons-folder"
        @change="onFilesChanged"
      />
      <div v-if="result" class="gap-1 flex">
        <UBadge :label="`Preis: ${result.price} €`" />
        <UBadge :label="`Material: ${result.material}`" />
      </div>
      <UInput
        v-if="warning"
        color="red"
        :ui="{ color: 'red' }"
        :value="warning"
        readonly
        icon="i-heroicons-exclamation-triangle"
      />
      <div class="submit">
        <UButton
          @click="onSlice"
          :disabled="!files.length || isSlicing || !!result"
        >
          Slice <Icon name="i-heroicons-square-3-stack-3d" />
        </UButton>
      </div>
    </div>
  </div>
</template>

<style scoped>
h1 {
  font-size: x-large;
  margin: 1em 0;
}

.logo {
  display: inline-block;
  height: 2em;
  margin-top: -3px;
  margin-left: -11px;
}

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
