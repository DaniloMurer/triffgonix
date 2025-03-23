import { getApiUser, type ModelsPlayer } from '#shared/utils';

export const usePlayerService = () => {
  const players = ref<ModelsPlayer[]>([]);

  const fetchPlayers = async () => {
    const response = await getApiUser();
    if (response.data) {
      return response.data;
    }
    return [];
  };

  return {
    players,
    fetchPlayers,
  };
};
