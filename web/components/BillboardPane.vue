<template>
  <div class="content">
    <div class="headline">
      <h1>Print Everything, Everywhere all at once!</h1>
      <p class="description">
        An easy to use 3d printing service with focus on high quality prints for
        everyone
      </p>
      <div class="actions">
        <UButton>
          <NuxtLink to="#getting-started">Getting Started</NuxtLink>
          <UIcon name="i-heroicons-arrow-right" />
        </UButton>
      </div>
    </div>
    <div class="image">
      <NuxtImg src="/images/logo.jpg" class="rounded-lg overflow-hidden" />
    </div>
  </div>
</template>

<style scoped>
.content {
  display: flex;
  flex-direction: row;
  gap: 2em;
  justify-content: center;
  align-items: center;
  min-height: 500px;
}

.actions {
  margin: 1em 0;
}

.headline {
  flex-basis: 30%;
  font-size: x-large;
}

.description {
  font-size: small;
  color: #666;
}

.image {
  flex-basis: 50%;
  max-width: 200px;
}
</style>
