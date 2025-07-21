import axios from 'axios';

const BASE_URL = 'http://localhost:5000/api/users';

export const getPosts = async (page = 1, limit = 5, field = 'id', order = 'desc') => {
  try {
    const response = await axios.get(`${BASE_URL}?page=${page}&limit=${limit}&field=${field}&order=${order}`);
    return response.data;
  } catch (error) {
    console.error('API Error:', error);
    throw error;
  }
};

export const getUserById = async (id) => {
  const res = await fetch(`/api/user/${id}`);
  return await res.json();
};

export const updateUser = async (id, userData) => {
  const res = await fetch(`/api/user/${id}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(userData),
  });
  return await res.json();
};
