<script setup lang="ts">
  import { useGames } from '~/composables/use-games';
  import { getPlayers } from '~/composables/get-players';
  import { usePlayerStore } from '~/store/user.store';

  const { games } = useGames();
  const { players } = getPlayers();

  const userStore = usePlayerStore();

  userStore.setPlayers(players.value);
</script>
<template>
  <UContainer class="p-20 flex flex-wrap justify-center gap-4">
    <UCard v-for="game in games" :key="game.id" class="mb-4 w-80">
      <template #header>
        <h1>{{ game.name }}</h1>
      </template>
      <h3>Players</h3>
      <UTable :data="game.players" class="flex-1 mt-4" />
    </UCard>
  </UContainer>
</template>
