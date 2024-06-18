<script setup lang="ts"></script>

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
        variant="outline"
        color="primary"
        placeholder="Deine Nachricht"
        :rows="7"
      />
      <UInput
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
      />
      <div class="submit">
        <UButton>Drucken <Icon name="i-heroicons-printer" /></UButton>
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
