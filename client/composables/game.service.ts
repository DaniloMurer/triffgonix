import { type DtoGame, getApiGame, type ModelsGame, postApiGame } from '#shared/utils';

export const useGameService = () => {
  const games = ref<ModelsGame[]>([]);
  const loading = ref(false);

  const fetchGames = async () => {
    loading.value = true;
    try {
      const response = await getApiGame();
      if (response.data) {
        return response.data;
      }
      return [];
    } catch (error) {
      console.error('Error fetching games:', error);
      return [];
    } finally {
      loading.value = false;
    }
  };

  return {
    games,
    loading,
    fetchGames,
  };
};

export const createGame = async (game: DtoGame) => {
  const response = await postApiGame({ body: game });
  console.log(response);
  if (response.response.ok) {
    return response.data;
  } else {
    throw new Error('error while creating game');
  }
};
