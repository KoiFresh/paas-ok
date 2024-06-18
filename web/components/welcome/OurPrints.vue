<script setup lang="ts">
const items = [
  "https://picsum.photos/600/800?random=1",
  "https://picsum.photos/600/800?random=2",
  "https://picsum.photos/600/800?random=3",
  "https://picsum.photos/600/800?random=4",
  "https://picsum.photos/600/800?random=5",
  "https://picsum.photos/600/800?random=6",
];
</script>

<template>
  <div>
    <h2 id="about-us"><UIcon name="i-heroicons-cube" /> Our Prints</h2>
    <UCarousel
      v-slot="{ item }"
      :items="items"
      :ui="{ item: 'basis-full md:basis-1/2 lg:basis-1/3' }"
      class="rounded-lg overflow-hidden"
      arrows
    >
      <img :src="item" class="w-full max-h-90 object-cover" draggable="false" />
    </UCarousel>
  </div>
</template>
