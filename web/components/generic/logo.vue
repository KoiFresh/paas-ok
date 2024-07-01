<script setup lang="ts">
const mode = useColorMode();
const color = computed(() => (mode.value === "dark" ? "white" : "black"));
</script>

<template>
  <NuxtImg
    format="webp"
    :src="`/images/logo_454x454_${color}_free.png`"
    class="rounded-lg overflow-hidden"
    alt="Logo"
    @click="() => $router.push('/')"
    style="cursor: pointer"
  />
</template>
