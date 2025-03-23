import type { ModelsPlayer } from '#shared/utils';
import { usePlayerService } from '~/composables/player.service';

export const usePlayerStore = defineStore('player', () => {
  const players = ref<ModelsPlayer[]>([]);

  const getPlayers = () => {
    return players.value;
  };

  const fetchPlayers = async () => {
    const playerService = usePlayerService();

    players.value = await playerService.fetchPlayers();
    return players;
  };

  const setPlayers = (newUsers: ModelsPlayer[]) => {
    console.log('setting players');
    players.value = newUsers;
  };

  return {
    players,
    setPlayers,
    getPlayers,
    fetchPlayers,
  };
});
