<script setup lang="ts">
  import { useSocketService } from '~/composables/socket.service';
  import type { GameStateContent, IncomingSocketMessage } from '#shared/types/socket';
  import { ref } from '#imports';

  const games = ref<GameStateContent[]>();
  const socketService = useSocketService();

  // eslint-disable-next-line no-undef
  const onMessage = (message: MessageEvent) => {
    console.log('message received');
    const content = JSON.parse(message.data) as IncomingSocketMessage;
    if (content.type === 'games') {
      games.value = content.content as GameStateContent[];
    }
  };
  socketService.connectToSocket('/ws/dart');
  socketService.listenOnMessage(onMessage);
</script>
<template>
  <UContainer class="p-20 flex flex-wrap justify-center gap-4">
    <UCard v-for="game in games" :key="game.id" class="mb-4 w-80">
      <template #header>
        <h1>{{ game.name }}</h1>
      </template>
      <h3>Players</h3>
      <UTable :data="game.players.allPlayers" class="flex-1 mt-4" />
    </UCard>
  </UContainer>
</template>
