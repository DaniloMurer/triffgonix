import { type ModelsGame } from '#shared/utils';
import { defineStore } from '#imports';
import { useGameService } from '~/composables/game.service';

export const useGameStore = defineStore('game', () => {
  const gamesState = ref<ModelsGame[]>([]);
  const gameService = useGameService();

  const getGames = () => {
    return gamesState.value;
  };

  const fetchGames = async () => {
    gamesState.value = await gameService.fetchGames();
    return gamesState;
  };

  const setGames = (newGames: ModelsGame[]) => {
    gamesState.value = newGames;
  };

  return {
    gamesState,
    setGames,
    getGames,
    fetchGames,
  };
});
