<script setup lang="ts">
  // import { useGameStore } from '~/store/game.store';
  import { useSocketService } from '~/composables/socket.service';
  import type {
    GameStateContent,
    IncomingSocketMessage,
    NewGameContent,
  } from '#shared/types/socket';
  import { ref } from '#imports';

  // const gameStore = useGameStore();
  const games = ref<GameStateContent[] | NewGameContent>();
  const socketService = useSocketService();
  let isNewGame = false;

  // eslint-disable-next-line no-undef
  const onMessage = (message: MessageEvent) => {
    console.log('new message: ', JSON.parse(message.data));
    const content = JSON.parse(message.data) as IncomingSocketMessage;
    if (content.content) {
      if (content.content instanceof NewGameContent) {
        isNewGame = true;
      }
      games.value = content.content;
    }
  };

  socketService.listenOnMessage(onMessage);
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
