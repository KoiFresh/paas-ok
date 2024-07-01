<script setup lang="ts">
const model = defineModel<SliceOptions>({
  default: {
    color: "red",
    material: "pla",
    infill: 30,
    quality: "low",
  },
});

const items = [
  {
    slot: "color",
    label: "Farbe",
    icon: "i-heroicons-swatch",
    description: "aa",
  },
  {
    slot: "quality",
    label: "Qualität",
    icon: "i-heroicons-wrench-screwdriver",
  },
  { slot: "material", label: "Material", icon: "i-heroicons-beaker" },
  { slot: "infill", label: "Infill", icon: "i-heroicons-chart-pie" },
];
</script>

<template>
  <UAccordion variant="outline" class="text-blue-500" :items="items">
    <template #color="{ item }">
      <GenericPicker
        :options="[
          { label: 'Weiß', value: 'white' },
          { label: 'Rot', value: 'red' },
          { label: 'Schwarz', value: 'black' },
        ]"
        v-model="model.color"
      />
    </template>

    <template #quality>
      <GenericPicker
        :options="[
          { label: 'Low', value: 'low' },
          { label: 'Medium', value: 'medium' },
          { label: 'High', value: 'high' },
        ]"
        v-model="model.quality"
      />
    </template>

    <template #material>
      <GenericPicker
        :options="[
          { label: 'PLA', value: 'pla' },
          { label: 'PETG', value: 'petg' },
          { label: 'ABS', value: 'abs' },
          { label: 'TPU', value: 'tpu' },
        ]"
        v-model="model.material"
      />
    </template>

    <template #infill>
      <div class="mx-5">
        <div class="flex justify-between">
          <span>10%</span>
          <span>30%</span>
          <span>50%</span>
          <span>70%</span>
          <span>90%</span>
        </div>
        <URange :min="10" :max="90" :step="10" v-model="model.infill" />
      </div>
    </template>
  </UAccordion>
</template>
