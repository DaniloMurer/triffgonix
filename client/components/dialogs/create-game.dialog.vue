<script setup lang="ts">
  import type { DtoGame } from '#shared/utils';
  import { getPlayers, onMounted, ref } from '#imports';
  import type { FormError } from '@nuxt/ui';
  import { createGame } from '~/composables/use-games';

  const formState = ref<DtoGame>({});
  const selectedPlayers = ref<number[]>([]);
  // const playerStore = usePlayerStore();
  // const players = ref<ModelsPlayer[]>();
  const { players } = getPlayers();

  onMounted(() => {});

  const validate = (state: DtoGame): FormError[] => {
    const errors: FormError[] = [];
    if (!state.name) errors.push({ name: 'name', message: 'Name is required' });
    if (!state.gameMode) errors.push({ name: 'gameMode', message: 'Game Mode is required' });
    return errors;
  };

  const onSubmit = async () => {
    console.log(formState.value);
    console.log(selectedPlayers.value);
    const result = await createGame({
      name: formState.value.name,
      gameMode: formState.value.gameMode,
      players: players.value.filter(player => selectedPlayers.value.includes(player.id)),
    });
    console.log(result);
  };
</script>
<template>
  <UModal title="Create Game">
    <UButton label="Create Game" icon="i-lucide-plus" />
    <template #body>
      <UForm
        :validate="validate"
        :state="formState"
        class="space-y-4 flex flex-col items-center"
        @submit="onSubmit"
      >
        <UFormField label="Name" name="name">
          <UInput v-model="formState.name" />
        </UFormField>
        <UFormField label="Game Mode" name="gameMode">
          <UInput v-model="formState.gameMode" />
        </UFormField>
        <UFormField label="Players" name="players">
          <USelect
            :items="players"
            v-model="selectedPlayers"
            label-key="username"
            value-key="id"
            multiple
            class="w-52"
          />
        </UFormField>
        <div class="flex justify-end">
          <UButton icon="i-lucide-save" type="submit">Save</UButton>
        </div>
      </UForm>
    </template>
  </UModal>
</template>
