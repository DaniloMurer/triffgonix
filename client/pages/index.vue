<script setup lang="ts">
import {useGames} from "~/composables/use-games";
import {usePlayers} from "~/composables/use-players";

const {games, loading} = useGames();
const {players} = usePlayers();
</script>
<template>
  <UContainer class="p-20 flex flex-wrap justify-center gap-4">
    <UCard v-for="user in players" :key="user.id" class="mb-4 w-80">
      <template #header>
        <h1>{{ user.username }}</h1>
      </template>
      <UBadge>{{ user.id }}</UBadge>
    </UCard>
  </UContainer>
</template>
