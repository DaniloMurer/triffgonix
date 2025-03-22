import type { ModelsPlayer } from '#shared/utils';

export const usePlayerStore = defineStore('player', () => {
  const players = ref<ModelsPlayer[]>([]);

  const getPlayers = () => {
    console.log('getting players');
    return players.value;
  };

  const setPlayers = (newUsers: ModelsPlayer[]) => {
    console.log('setting players');
    players.value = newUsers;
  };

  return {
    users: players,
    setPlayers,
    getPlayers,
  };
});
