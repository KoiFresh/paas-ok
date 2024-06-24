<script setup lang="ts">
const config = useRuntimeConfig();

const message = ref("");
const email = ref("");
const files = ref<File[]>([]);

/**
 * Called when the form is submitted
 */
async function onSubmit() {
  const form = new FormData();
  form.append("email", email.value);
  form.append("message", message.value);
  for (const file of files.value) {
    form.append("files", file, file.name);
  }

  try {
    const result = await fetch(`${config.public.apiBaseUrl}/order`, {
      method: "POST",
      body: form,
    });

    if (!result.ok) {
      alert("Oops, something went wrong!");
    }
  } catch (error) {
    alert(`Oops, something went wrong! ${error}`);
  }
}

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
      <UTextarea
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
      />
      <UInput
        variant="outline"
        color="primary"
        placeholder="Name"
        type="file"
        multiple
        icon="i-heroicons-folder"
        @change="onFilesChanged"
      />
      <div class="submit">
        <UButton
          @click="onSubmit"
          :disabled="!message || !email || !files.length"
          >Drucken <Icon name="i-heroicons-printer"
        /></UButton>
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
