<script setup lang="ts">
const propeties = defineProps<{
  options: {
    label: string;
    value: string;
  }[];
}>();

const value = defineModel<string>();
</script>

<template>
  <div class="flex flex-wrap gap-1 mx-5">
    <UButton
      v-for="item in propeties.options"
      :variant="item.value == value ? 'solid' : 'outline'"
      @click="value = item.value"
    >
      {{ item.label }}
    </UButton>
  </div>
</template>
