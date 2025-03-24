<script setup lang="ts">
  import type { DtoGame } from '#shared/utils';
  import { ref, useI18n } from '#imports';
  import type { FormError } from '@nuxt/ui';
  import { createGame } from '~/composables/game.service';
  import { useGameStore } from '~/store/game.store';
  import { usePlayerStore } from '~/store/player.store';

  const i18n = useI18n();
  const gameStore = useGameStore();
  const playerStore = usePlayerStore();

  const players = await playerStore.fetchPlayers();

  const formState = ref<DtoGame>({});
  const selectedPlayers = ref<number[]>([]);
  const isOpen = ref(false);

  const validate = (state: DtoGame): FormError[] => {
    const errors: FormError[] = [];
    if (!state.name) errors.push({ name: 'name', message: i18n.t('name-required-error') });
    if (!state.gameMode)
      errors.push({ name: 'gameMode', message: i18n.t('gamemode-required-error') });
    if (!(selectedPlayers.value.length >= 2))
      errors.push({ name: 'players', message: i18n.t('players-required-error') });
    return errors;
  };

  const onSubmit = async () => {
    await createGame({
      name: formState.value.name,
      gameMode: formState.value.gameMode,
      players: players.value.filter(player => selectedPlayers.value.includes(player.id!)),
    }).catch(error => {
      console.error(error);
    });
    gameStore.fetchGames();
    onClose();
  };

  const onClose = () => {
    isOpen.value = false;
    formState.value = {};
    selectedPlayers.value = [];
  };
</script>
<template>
  <UModal v-model:open="isOpen" :close="false" :dismissible="false" :title="$t('create-game')">
    <UButton :label="$t('create-game')" icon="i-lucide-plus" />
    <template #body>
      <UForm
        :validate="validate"
        :state="formState"
        class="space-y-4 flex flex-col items-center"
        @submit="onSubmit"
      >
        <UFormField :label="$t('game-name')" name="name">
          <UInput v-model="formState.name" />
        </UFormField>
        <UFormField :label="$t('game-mode')" name="gameMode">
          <UInput v-model="formState.gameMode" />
        </UFormField>
        <UFormField :label="$t('players')" name="players">
          <USelect
            :items="players"
            v-model="selectedPlayers"
            label-key="name"
            value-key="id"
            multiple
            class="w-52"
          />
        </UFormField>
        <div class="flex justify-end gap-11">
          <UButton icon="i-lucide-save" type="submit">{{ $t('save') }}</UButton>
          <UButton icon="i-lucide-x" @click="onClose">{{ $t('cancel') }}</UButton>
        </div>
      </UForm>
    </template>
  </UModal>
</template>
