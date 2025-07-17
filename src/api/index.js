import axios from 'axios';

const BASE_URL = 'http://localhost:5000/api/users';

export const getPosts = async (page = 1, limit = 5) => {
  try {
    const response = await axios.get(`${BASE_URL}?page=${page}&limit=${limit}`);
    return response.data;
  } catch (error) {
    console.error('API Error:', error);
    throw error;
  }
};
