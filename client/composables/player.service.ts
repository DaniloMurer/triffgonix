import { getApiPlayer } from '#shared/utils';

export const usePlayerService = () => {
  const fetchPlayers = async () => {
    const response = await getApiPlayer();
    if (response.data) {
      return response.data;
    }
    return [];
  };

  return {
    fetchPlayers,
  };
};
