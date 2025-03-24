import { getApiUser } from '#shared/utils';

export const usePlayerService = () => {
  const fetchPlayers = async () => {
    const response = await getApiUser();
    if (response.data) {
      return response.data;
    }
    return [];
  };

  return {
    fetchPlayers,
  };
};
