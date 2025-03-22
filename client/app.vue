<script setup lang="ts">
  import CreateGameDialog from './components/dialogs/create-game.dialog.vue';
  import { ref } from '#imports';

  const items = ref([
    {
      label: 'Triffgonix',
      icon: 'i-lucide-home',
      to: '/',
    },
    {
      slot: 'create-game',
    },
  ]);
</script>

<template>
  <UApp>
    <div class="m-3">
      <UNavigationMenu :items="items" orientation="horizontal" class="justify-center">
        <template #create-game>
          <create-game-dialog />
        </template>
      </UNavigationMenu>
      <NuxtRouteAnnouncer />
      <NuxtPage />
    </div>
  </UApp>
</template>
