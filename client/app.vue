<script setup lang="ts">
const items = ref([
    {
      label: 'Triffgonix',
      icon: 'i-lucide-home',
      to: '/'
    }
])
</script>

<template>
  <UApp>
    <div class="m-3">
      <UNavigationMenu :items="items" orientation="horizontal"/>
      <NuxtRouteAnnouncer/>
      <NuxtPage/>
    </div>
  </UApp>
</template>
