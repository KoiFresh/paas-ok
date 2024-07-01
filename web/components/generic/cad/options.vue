<script setup lang="ts">
const propeties = defineProps<{
  disabled?: boolean;
}>();

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
        :disabled="propeties.disabled"
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
        :disabled="propeties.disabled"
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
        :disabled="propeties.disabled"
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
        <div
          class="label flex justify-between"
          style="width: 109%; transform: translateX(-4%)"
        >
          <span class="text-red-500"> 0%</span>
          <span>25%</span>
          <span>50%</span>
          <span>75%</span>
          <span class="text-red-500">100%</span>
        </div>
        <URange
          :disabled="propeties.disabled"
          :min="0"
          :max="100"
          :step="5"
          v-model="model.infill"
        />
      </div>
    </template>
  </UAccordion>
</template>

<style scoped>
.label {
}

.label span {
  width: 10%;
  text-align: center;
}
</style>
