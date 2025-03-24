import type { DtoPlayer } from '#shared/utils';
import { usePlayerService } from '~/composables/player.service';

export const usePlayerStore = defineStore('player', () => {
  const players = ref<DtoPlayer[]>([]);
  const playerService = usePlayerService();

  const getPlayers = () => {
    return players.value;
  };

  const fetchPlayers = async () => {
    players.value = await playerService.fetchPlayers();
    return players;
  };

  const setPlayers = (newUsers: DtoPlayer[]) => {
    players.value = newUsers;
  };

  return {
    players,
    setPlayers,
    getPlayers,
    fetchPlayers,
  };
});
