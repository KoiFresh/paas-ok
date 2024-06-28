<script setup lang="ts">
const config = useRuntimeConfig();

const message = ref("");
const email = ref("");
const files = ref<File[]>([]);

const result = ref<RemoteSlicerResult | null>(null);
const { slice, isSlicing } = useRemoteSlicer();

async function onSlice() {
  result.value = null;
  result.value = await slice(files.value[0]);
}

/**
 * Called when the uploaded files changed
 * @param event FileList
 */
function onFilesChanged(event: FileList) {
  result.value = null;
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
      <!--<UTextarea
        v-model="message"
        variant="outline"
        color="primary"
        placeholder="Deine Nachricht"
        :rows="7"
      />
      <UInput
        v-model="email"
        variant="outline"
        color="primary"
        placeholder="Email"
        icon="i-heroicons-at-symbol"
      />-->
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
